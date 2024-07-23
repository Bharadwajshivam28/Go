package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalln("can't listen to network")
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("err while accepting.....", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	fmt.Println("connected to: ", conn.RemoteAddr())
	for {
		var buffer [1024]byte
		_, err := conn.Read(buffer[:])
		if err != nil {
			log.Printf("err while reading from connection: %v, exiting...", err)
			return
		}
		fmt.Println("message read: ", string(buffer[:]))
	}
}