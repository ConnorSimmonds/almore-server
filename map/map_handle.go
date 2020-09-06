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
	_, err := currentMap.Seek(0, 0) //reset back to the beginning
	checkError(err)
	w := make([]byte, 1)
	h := make([]byte, 1)

	_, err = currentMap.Read(w)
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
	stringBuilder.WriteString(strconv.FormatUint(uint64(userID), 10) + "/map" + strconv.FormatUint(uint64(dungeonID), 10) + "_" + strconv.FormatUint(uint64(mapNum), 10) + fileFormat)
	file, fileError := os.OpenFile(stringBuilder.String(), os.O_RDWR, os.ModePerm)
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
	checkError(e)
	_, e = returnFile.Write([]byte{x, y})
	checkError(e)
	//We now fill it with dummy characters, and put the x/y values in the header
	for i := 0; i < int(x); i++ {
		for i2 := 0; i2 < int(y); i2++ {
			_, e = returnFile.Write([]byte{0})
			checkError(e)
		}
	}
	return returnFile //return the file
}

func SendMap(currentMap *os.File) ([]byte, error) { //Creates a byte array which will then be returned. This contains the map file
	_, err := currentMap.Seek(0, 0) //reset file
	checkError(err)
	w := make([]byte, 1)
	h := make([]byte, 1)

	_, err = currentMap.Read(w)
	checkError(err)
	_, err = currentMap.ReadAt(h, 1)
	checkError(err)
	byteArray := make([]byte, w[0]*h[0])
	_, err = currentMap.Read(byteArray) //append the file to the byte array, which is the same size as our map file
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}
