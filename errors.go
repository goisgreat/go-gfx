package physics

import "os"

// handle() handles a given error
// [state] should have information about what was being attempted
func Handle(err error, state string) {
	// if error is not nil, print it to the command line and exit
	if err != nil {
		println("Error while", state, err.Error())
		os.Exit(1)
	}
}
