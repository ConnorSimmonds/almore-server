//This package handles all of the user modification (i.e. initializing on log-in).
package user

import (
	"encoding/binary"
)

//This initializes a user upon log-in. It returns the userID given, for later usage.
//TODO: Basically all of the extra functionality in this
func InitUser(userID []byte) uint16 {
	var user uint16
	user = binary.BigEndian.Uint16(userID) //Thanks to https://stackoverflow.com/questions/11184336/how-to-convert-from-byte-to-int-in-go-programming
	return user
}

//This function is called if the SQL database cannot find the userID provided in InitUser.
func createUser(userID byte) {

}
