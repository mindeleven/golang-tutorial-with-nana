// ///////////////////////////////
// Coding along with Golang Tutorial for Beginners | Full Go Course by Nana
// https://www.youtube.com/watch?v=yyUHQIec83I
// Source code and comments take from the videos of this course
// Go documentation @ https://go.dev/doc/
// https://youtu.be/yyUHQIec83I?feature=shared&t=3081
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
	// syntactic sugar of go language (not for declaring types od constants)
	conferenceName := "Go Conference"

	// printing out the values
	// for printing see https://pkg.go.dev/fmt@go1.22.0
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// printing out the types
	fmt.Printf("conferenceName is type %T, conferenceTickets is type %T, remainingTickets is type %T\n", conferenceName, conferenceTickets, remainingTickets)

	// printing the memory location of the remaining tickets
	fmt.Println("memory location of the remaining tickets:", &remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// _, _ = firstName, userTickets

	// asking users for their personal information
	// reading user input from the command line with fmt.Scan()
	// scanning user input and assigning it
	// adding a pointer to the firstName variable
	// a pointer is a variable that points to the memory address of another variable
	// pointers in golang are called special variables
	fmt.Println("Dear user, please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Dear user, please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Dear user, please enter your email address:")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want:")
	fmt.Scan(&userTickets)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v.\n",
		firstName, lastName, userTickets, email)

}
