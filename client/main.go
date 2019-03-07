package main

import (
	"bufio"
	"flag"
	"io"
	"net"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "9091", "server port")
	flag.Parse()

	conn, err := net.Dial("tcp", ":" + port)
	if err != nil {
		println("connect failed")
		return
	}
	var bs []byte
	reader := bufio.NewReader(conn)
	for ;; {
		var tb []byte
		tb = make([]byte, 10)
		n, err := reader.Read(tb)
		if err == io.EOF {
			println("end!!")
			break
		}
		if err != nil {
			println(err.Error())
			break
		}
		bs = append(bs, tb[0:n]...)
	}

	println(string(bs))

	conn.Close()

}
