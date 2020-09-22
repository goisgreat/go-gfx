package physics

import (
	"os"
	"os/exec"
)

/*
 * disable input buffering from stdin (so that you don't have to press ENTER)
 */
func DisableStdinBuffer() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
}

/*
 * disable entered characters from being "echo"ed
 */
func DisableStdinEcho() {
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

/*
 * accept event stream object
 * loop {
	* get input from callback
	*
 }
*/
func SendKeyboardInput(getByte func() []byte, channel EventStream) {
	// loop
	for {
		// get a character (represented as a byte) from stdin
		chars := getByte()

		// push character data onto stream
		for idx := range chars {
			channel.Input <- chars[idx]
		}
	}
}
