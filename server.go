package main

import (
	"bytes"
	"fmt"
	errorLibrary "github.com/ConnorSimmonds/server/errors"
	maplib "github.com/ConnorSimmonds/server/map"
	user "github.com/ConnorSimmonds/server/user"
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

	//Some basic values for gameplay purposes
	var userID uint32
	var dungeonID uint16
	var mapNum uint16
	var currentMap *os.File
	byteArray := make([]byte, 1024)

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

		//Note that this just handles receiving data - the actual proper logic is done elsewhere.
		switch clientCode[0] {
		case 0:
			userID = user.InitUser(clientRead.Next(4))
			fmt.Printf("Found user %d", userID)
			sendPacket(clientWrite, conn, []byte{0})
			break
		case 2:
			sendPacket(clientWrite, conn, []byte{2})
			fmt.Println("Closed connection with client.")
			break Loop
		case 1:
			fmt.Println("Received 'ping' from " + conn.RemoteAddr().String())
			sendPacket(clientWrite, conn, []byte{1})
			break
		case 10:
			//Put values into local variables
			x := clientRead.Next(1)[0]
			y := clientRead.Next(1)[0]
			value := clientRead.Next(1)[0]
			maplib.UpdateMap(x, y, value, currentMap)
			break
		case 13:
			currentMap, err = maplib.OpenMap(userID, dungeonID, mapNum)
			if err != nil { //Let's do some error handling
				if err == errorLibrary.ReturnMapFileError() { //The map doesn't exist! Let's get the map from the client
					sendPacket(clientWrite, conn, []byte{12})
				}
			}
			break
		case 14:
			currentMap = maplib.CreateMap("", 0, 0)
			break
		}
	}
}

func sendPacket(buffer bytes.Buffer, conn net.Conn, packetDetails []byte) {
	buffer.Write(packetDetails)
	_, _ = conn.Write(buffer.Bytes())
}
