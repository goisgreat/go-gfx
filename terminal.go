package gogfx

/*
 * clear the terminal using ascii escape sequences
 */
func ClearTerminal() {
	print("\033[H\033[2J")
}
