//The map package handles all of the dungeon map file handling. Through this, we're able to update the map files,
//and so on.
package _map

import (
	"os"
)

//Updates the map at the position of arguments x, y with the value of argument value
func UpdateMap(x int, y int, value int, userID uint16) {

}

//Creates the map
func createMap() {
	fileLocation := "Maps/"
	//We need to get the user ID etc. in memory

	os.Create(fileLocation)
}
