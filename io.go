package physics

import (
	"os"
	"os/exec"
)

/*
 * disable input buffering
 * do not display entered characters on the screen
 */
func Setup() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

/*
 * usage: call this function, get a callback, invoke the callback, and a byte or error will be returned
 * return function {
	* get character
	* get error
	* return character,
 }
*/
func GetCharReader() func() (byte, error) {
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

/*
 * accept event stream object
 * loop {
	* TODO: pull data from input sources
	* TODO: push data into stream
	* delay so we don't gobble up the CPU
 }
*/
func GrabInput(channel EventStream) {
	// create character reader
	getchar := GetCharReader()

	// loop
	for {
		// get a character (represented as a byte) from stdin
		char, err := getchar()

		if err != nil {
			panic(err)
		}

		// push character data onto stream
		channel.Input <- char
	}
}
