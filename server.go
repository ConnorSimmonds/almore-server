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
		clientRead.Reset()
		clientRead.Write(byteArray)
		clientString, err := clientRead.ReadString(10)
		if err != nil {
			fmt.Println("Error while parsing string - exiting.")
			break
		} else {
			switch clientString {
			case "quit\n":
				sendMessage(clientWrite, conn, "Quitting...\n")
				break Loop
			case "ping\n":
				fmt.Println("Received 'ping' from " + conn.RemoteAddr().String())
				sendMessage(clientWrite, conn, "pong\n")
				break
			}
		}
	}

}

func sendMessage(buffer bytes.Buffer, conn net.Conn, message string) {
	buffer.Write([]byte(message))
	conn.Write(buffer.Bytes())
}
