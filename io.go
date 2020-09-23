package physics

import (
	"os"
)

/*
 * usage: call this function, get a callback, invoke the callback, and a byte or error will be returned
 * return function {
	* get character
	* get error
	* return character,
 }
*/
func CreateStdinReader() func() (byte, error) {
	// store a list of bytes
	var bytes []byte = make([]byte, 1)

	// return function to get a character from stdin
	return func() (byte, error) {
		// read character from stdin
		_, err := os.Stdin.Read(bytes)
		// if error...
		if err != nil {
			// return no bytes but error
			return byte(0), err
		}
		// return first byte and no error
		return bytes[0], nil
	}
}
