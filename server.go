package main

import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"fmt"
	errlib "github.com/ConnorSimmonds/server/errors"
	maplib "github.com/ConnorSimmonds/server/map"
	user "github.com/ConnorSimmonds/server/user"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"os"
)

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
	var filename string
	var db *sql.DB
	var dbCache bool
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

		//Note that this just handles receiving data - the actual proper logic is done elsewhere for most packets
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
			var er *errlib.FileNotFoundError
			dungeonID = binary.LittleEndian.Uint16(clientRead.Next(2))
			mapNum = binary.LittleEndian.Uint16(clientRead.Next(2))
			fmt.Println("Opening map" + string(dungeonID) + "_" + string(mapNum))

			currentMap, er = maplib.OpenMap(userID, dungeonID, mapNum)
			if er != nil {
				filename = er.File
				sendPacket(clientWrite, conn, []byte{12})
			}
			break
		case 14:
			//Create the map, getting the map x/y values from the client
			//Put values into local variables
			x := clientRead.Next(2)[0]
			y := clientRead.Next(2)[0]
			currentMap = maplib.CreateMap(filename, x, y)
			break

		case 20:
			break
		case 21:
			break
		case 30:
			db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/labyrinth")
			if err != nil {
				log.Fatal(err)
			}
			err = db.Ping() //ping the database, to make sure there's no errors
			if err != nil {
				//The database was not set up properly, so figure out what's wrong and act from there.
			}
			break
		case 31: //Tells the server to start caching any and all changes to the party members.
			//This is to reduce the amount of DB calls we make.
			dbCache = true
			break

		case 38:
			if dbCache { //changes have been stored, so apply them

			} else {
				//let the client know that changes haven't even started initializing
			}
		case 39:
			err = db.Close()
			if err != nil {
				//something went wrong while closing the db - look up how to do this
			}
		}

	}
}

func sendPacket(buffer bytes.Buffer, conn net.Conn, packetDetails []byte) {
	buffer.Write(packetDetails)
	_, _ = conn.Write(buffer.Bytes())
}
