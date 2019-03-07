package main

import (
	"bufio"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":9092")
	conn, _ := ln.Accept()

	bs := []byte("hello world")
	blen := len(bs)
	wl := 0
	for ;; {
		l, e := conn.Write(bs)
		if e != nil {
			println(e.Error())
			break
		}
		wl += l
		if blen > wl {
			bs = bs[l:]
		} else {
			break
		}
	}

	conn.(*net.TCPConn).CloseWrite()

	reader := bufio.NewReader(conn)
	for ;; {
		var bs = make([]byte, 10)
		_, err := reader.Read(bs)
		if err != nil {
			break
		}
	}
	conn.(*net.TCPConn).CloseRead()

	ln.Close()
}
