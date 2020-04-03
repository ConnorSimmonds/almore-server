//The map package handles all of the dungeon map file handling. Through this, we're able to update the map files,
//and so on.
package _map

import (
	"fmt"
	errlib "github.com/ConnorSimmonds/server/errors"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

//Updates the map at the position of arguments x, y with the value of argument value
func UpdateMap(x uint8, y uint8, value uint8, currentMap *os.File) {
	//firstly, let's get the width/height
	currentMap.Seek(0, 0) //reset back to the beginning
	w := make([]byte, 1)
	h := make([]byte, 1)

	_, err := currentMap.Read(w)
	checkError(err)
	_, err = currentMap.ReadAt(h, 1)
	checkError(err)

	if x > w[0] || y > h[0] {
		//we're out of bounds - somehow
		return
	}
	offset := int64((h[0]*y)+x) + 2
	num, err := currentMap.WriteAt([]byte{value}, offset) //I need to properly set this up - there's going to be issues. It's easy enough to say "go to position x,y"

	if err != nil {
		if num != 1 { //First off, let's see if we were even able to WRITE our one byte
			panic(err)
		}
	}
}

//Opens a map, and returns it for later use. If the file doesn't exist, it will create it.
func OpenMap(userID uint32, dungeonID uint16, mapNum uint16) (*os.File, *errlib.FileNotFoundError) {
	var stringBuilder strings.Builder
	fileFormat := ".dng"
	stringBuilder.WriteString("Maps/")
	//Now we add our ints into our builder
	stringBuilder.WriteString(strconv.FormatUint(uint64(userID), 10) + "/" + strconv.FormatUint(uint64(dungeonID), 10) + "_" + strconv.FormatUint(uint64(mapNum), 10) + fileFormat)
	fmt.Print(stringBuilder.String())
	file, fileError := os.OpenFile("Maps/1/map1_0.dng", os.O_RDWR, os.ModePerm)
	if fileError != nil {
		//There's been some kind of error, record it in the appropriate debug log (with a timestamp) and then create the map (failsafe)
		fmt.Println(fileError.Error())
		//let's check the type of error - if it's a file not found, we handle it. Otherwise, panic.
		return nil, errlib.MapFileNotFoundError(stringBuilder.String(), fileError)
	} else {
		return file, nil
	}
}

//Creates the map file
func CreateMap(filename string, x uint8, y uint8) *os.File {
	returnFile, e := os.Create(filename)
	if e != nil { //We encountered an error we weren't supposed to!
		//we'll attempt to handle it
		panic(e) //Flip out and panic!
	}
	returnFile.Write([]byte{x, y})
	//We now fill it with dummy characters, and put the x/y values in the header
	for i := 0; i < int(x); i++ {
		for i2 := 0; i2 < int(y); i2++ {
			returnFile.Write([]byte{0})
		}
	}
	return returnFile //return the file
}
