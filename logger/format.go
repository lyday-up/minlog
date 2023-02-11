package logger

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"time"
)

type Formatter interface {
	Format(entry *Entry) error
}

type TextFormatter struct {
	IgnoreBasicFields bool
}

func (f *TextFormatter) Format(e *Entry) error {
	if !f.IgnoreBasicFields {
		e.Buffer.WriteString(fmt.Sprintf("%s %s", e.Time.Format(time.RFC3339), LevelMapping[e.Level])) // allocs
		if e.File != "" {
			short := e.File
			for i := len(e.File) - 1; i > 0; i-- {
				if e.File[i] == '/' {
					short = e.File[i+1:]
					break
				}
			}
			e.Buffer.WriteString(fmt.Sprintf(" %s:%d", short, e.Line))
		}
		e.Buffer.WriteString(" ")
	}

	switch e.Format {
	case FmtEmptySeparate:
		e.Buffer.WriteString(fmt.Sprint(e.Args...))
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}
	e.Buffer.WriteString("\n")

	return nil
}

type JsonFormatter struct {
	IgnoreBasicFields bool
}

func (f *JsonFormatter) Format(e *Entry) error {
	if !f.IgnoreBasicFields {
		e.Map["level"] = LevelMapping[e.Level]
		e.Map["time"] = e.Time.Format(time.RFC3339)
		if e.File != "" {
			e.Map["file"] = e.File + ":" + strconv.Itoa(e.Line)
			e.Map["func"] = e.Func
		}

		switch e.Format {
		case FmtEmptySeparate:
			e.Map["message"] = fmt.Sprint(e.Args...)
		default:
			e.Map["message"] = fmt.Sprintf(e.Format, e.Args...)
		}

		return jsoniter.NewEncoder(e.Buffer).Encode(e.Map)
	}

	switch e.Format {
	case FmtEmptySeparate:
		for _, arg := range e.Args {
			if err := jsoniter.NewEncoder(e.Buffer).Encode(arg); err != nil {
				return err
			}
		}
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}

	return nil
}
