package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

import _ "github.com/go-sql-driver/mysql"

func main() {
	//Do note that this is literally copy/pasted from the GO documentation
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error while setting up port.")
		os.Exit(1)
	}
	fmt.Println("Accepting requests on port 8080.")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("Error while accepting request.")
			os.Exit(1)
		}
		fmt.Println("Found connection: " + conn.RemoteAddr().String() + " of type " + conn.LocalAddr().Network())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	//Simple test to receive and send to a client
	var clientRead bytes.Buffer
	var clientWrite bytes.Buffer
	byteArray := make([]byte, 1024)

	sendMessage(clientWrite, conn, "Hello World\n")
	fmt.Println("Waiting for response from " + conn.RemoteAddr().String())
	defer conn.Close()
Loop:
	for true {
		_, err := conn.Read(byteArray)
		if err != nil {
			fmt.Println("Error while reading from client.")
			break
		}
		//Getting a response from the client
		clientRead.Reset()
		clientRead.Write(byteArray)
		clientCode := clientRead.Next(1)
		switch clientCode[0] {
		case 2:
			sendPacket(clientWrite, conn, []byte{2})
			break Loop
		case 1:
			fmt.Println("Received 'ping' from " + conn.RemoteAddr().String())
			sendPacket(clientWrite, conn, []byte{1})
			break
		}
	}
}

func sendMessage(buffer bytes.Buffer, conn net.Conn, message string) {
	buffer.Write([]byte(message))
	conn.Write(buffer.Bytes())
}

func sendPacket(buffer bytes.Buffer, conn net.Conn, packetDetails []byte) {
	buffer.Write(packetDetails)
	conn.Write(buffer.Bytes())
}
