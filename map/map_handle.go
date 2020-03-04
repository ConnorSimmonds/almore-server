//The map package handles all of the dungeon map file handling. Through this, we're able to update the map files,
//and so on.
package _map

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Updates the map at the position of arguments x, y with the value of argument value
func UpdateMap(x uint8, y uint8, value uint8, currentMap *os.File) {
	currentMap.WriteAt([]byte{value}, 0) //I need to properly set this up - there's going to be issues. It's easy enough to say "go to position x,y"
}

//Opens a map, and returns it for later use. If the file doesn't exist, it will create it.
func OpenMap(userID uint32, dungeonID uint16, mapNum uint16) *os.File {
	var stringBuilder strings.Builder
	fileFormat := ".dng"
	stringBuilder.WriteString("Maps/")
	//Now we add our ints into our builder
	stringBuilder.WriteString(strconv.FormatUint(uint64(userID), 10) + "/" + strconv.FormatUint(uint64(dungeonID), 10) + "_" + strconv.FormatUint(uint64(mapNum), 10) + fileFormat)
	fmt.Print(stringBuilder.String())
	file, fileError := os.OpenFile("Maps/1/map1_1.dng", os.O_RDWR, os.ModePerm)
	if fileError != nil {
		//There's been some kind of error, record it in the appropriate debug log (with a timestamp) and then create the map (failsafe)
		fmt.Println(fileError.Error())
		return createMap(stringBuilder.String())
	} else {
		return file
	}
}

//Creates the map file
func createMap(filename string) *os.File {
	var returnFile *os.File
	returnFile, _ = os.Create(filename)
	return returnFile //return the file
}
