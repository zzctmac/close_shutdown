package main

import "net"

func main() {
	ln, _ := net.Listen("tcp", ":9091")
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

	conn.Close()

	ln.Close()
}
