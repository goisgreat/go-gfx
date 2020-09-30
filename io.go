package physics

import (
	"os"
)

// CreateStdinReader() returns a stdin reader that reports the character that was just pressed and any errors.
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
