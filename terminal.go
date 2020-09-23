package physics

import "os/exec"

/*
 * clear the terminal using ascii escape sequences
 */
func ClearTerminal() {
	print("\033[H\033[2J")
}

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
