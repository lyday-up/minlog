import socket
sk = socket.socket()
ip_port = ("127.0.0.1",9876)
sk.bind(ip_port)
sk.listen(5)
while True:
    conn, address = sk.accept()
    while True:
        with open('file.py','ab') as f:
            data = conn.recv(1024)
            if data == b'quit':
                break
            f.write(data)
            conn.send("ok".encode())
    conn.send("success".encode())
    # print('OK')
sk.close()
