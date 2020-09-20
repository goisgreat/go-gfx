package main

import "log"

/*
 * "handle" a given error
 * note that passing `nil` to this function will NOT crash the program
 */
func handle(err error) {
	// if error is not nil...
	if err != nil {
		// log error and crash program
		log.Fatal(err.Error())
	}
}
