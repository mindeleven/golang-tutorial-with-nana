// ///////////////////////////////
// Coding along with Golang Tutorial for Beginners | Full Go Course by Nana
// https://www.youtube.com/watch?v=yyUHQIec83I
// Source code and comments take from the videos of this course
// Go documentation @ https://go.dev/doc/
// https://youtu.be/yyUHQIec83I?feature=shared&t=2758
// ///////////////////////////////
package main

import (
	"fmt"
)

func main() {
	const conferenceTickets = 15
	// assigning a type explicitly so that value can't be negative
	var remainingTickets uint = 50
	// alternative shorthand version (only for variables)
	// syntactic sugar of go language
	conferenceName := "Go Conference"

	// printing out the values
	// for printing see https://pkg.go.dev/fmt@go1.22.0
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// printing out the types
	fmt.Printf("conferenceName is type %T, conferenceTickets is type %T, remainingTickets is type %T\n", conferenceName, conferenceTickets, remainingTickets)

	var userName string
	var userTickets int
	_, _ = userName, userTickets
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
