package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"
)

var defaultLog = NewLog()

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

func NewLog(opts ...Option) *logger {
	logger := &logger{opt: initOptions(opts...)}
	logger.entryPool = &sync.Pool{New: func() interface{} { return entry(logger) }}
	return logger
}

func SetOptions(opts ...Option) {
	defaultLog.SetOptions(opts...)
}

func (l *logger) SetOptions(opts ...Option) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}

func Writer() io.Writer {
	return defaultLog
}

func (l *logger) Writer() io.Writer {
	return l
}

func (l *logger) Write(data []byte) (int, error) {
	l.entry().write(l.opt.stdLevel, FmtEmptySeparate, *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func (l *logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *logger) Debug(args ...interface{}) {
	l.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.entry().write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func (l *logger) Fatal(args ...interface{}) {
	l.entry().write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry().write(DebugLevel, format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.entry().write(InfoLevel, format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.entry().write(WarnLevel, format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry().write(ErrorLevel, format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}

// default logger
func Debug(args ...interface{}) {
	defaultLog.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func Info(args ...interface{}) {
	defaultLog.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func Warn(args ...interface{}) {
	defaultLog.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func Error(args ...interface{}) {
	defaultLog.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func Panic(args ...interface{}) {
	defaultLog.entry().write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func Fatal(args ...interface{}) {
	defaultLog.entry().write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func Debugf(format string, args ...interface{}) {
	defaultLog.entry().write(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLog.entry().write(InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLog.entry().write(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLog.entry().write(ErrorLevel, format, args...)
}

func Panicf(format string, args ...interface{}) {
	defaultLog.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	defaultLog.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}
