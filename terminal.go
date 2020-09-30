package physics

import "os/exec"

// ClearTerminal() uses ascii escape sequences to clear the screen.
func ClearTerminal() {
	print("\033[H\033[2J")
}

// DisableStdinBuffer() disables input buffering from stdin via stty.
// Once invoked, you will no longer have to press ENTER before stdin is submitted.
func DisableStdinBuffer() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
}

// DisableStdinEcho() hides entered characters (from stdin) from being displayed.
func DisableStdinEcho() {
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}
