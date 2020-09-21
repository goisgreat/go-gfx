package physics

import "os"

/*
 * "handle" a given error
 * note that passing `nil` to this function will NOT crash the program
 */
func Handle(err error, state string) {
	// if error is not nil, print it to the command line
	if err != nil {
		println(state, ": ", err.Error())
		os.Exit(1)
	}
}
