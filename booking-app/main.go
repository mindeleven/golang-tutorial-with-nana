// ///////////////////////////////
// Coding along with Golang Tutorial for Beginners | Full Go Course by Nana
// https://www.youtube.com/watch?v=yyUHQIec83I
// Source code and comments take from the videos of this course
// Go documentation @ https://go.dev/doc/
// https://youtu.be/yyUHQIec83I?feature=shared&t=4504
// ///////////////////////////////
package main

import (
	"fmt"
	"strings"
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

	// arrays in go
	// arrays have a fixed size
	// set to 50 because we expect no more than 50 bookings
	// type follows after number
	// can have value afterwards (can be set to zero or empty)
	var bookings [50]string
	_ = bookings

	// slices in go (dynamic list)
	// slice is an abstraction of an array with a dynamic size
	// creating a slice is like creating an array with a size definition
	var bookings_slice []string
	_ = bookings_slice
	// alternative syntax:
	var bookings_slice_2 = []string{}
	_ = bookings_slice_2
	bookings_slice_3 := []string{}
	_ = bookings_slice_3

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
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

		// logic for updating the remaining tickets
		if userTickets > remainingTickets {
			fmt.Printf("We've only %v tickest remaining, so you can't book %v tickets.", remainingTickets, userTickets)
			// user wants too many tickest so be break the loop
			break
		}
		remainingTickets = remainingTickets - userTickets
		// storing the bookings
		bookings[0] = firstName + " " + lastName
		// adding the next element to a slice
		// mot necessary to know the index
		bookings_slice = append(bookings_slice, firstName+" "+lastName)

		// printing the bookings
		fmt.Printf("All bookings (whole array): %v\n", bookings)
		fmt.Printf("My very first booking (first element of array): %v\n", bookings[0])
		fmt.Printf("Type of bookings: %T\n", bookings)
		fmt.Printf("Length of bookings: %v\n", len(bookings))
		fmt.Printf("Length of bookings_slice: %v\n", len(bookings_slice))

		fmt.Printf(
			"Thank you %v %v for booking %v tickets. You'll receive a confirmation email at %v.\n",
			firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets are remaining for the %v.\n", remainingTickets, conferenceName)

		// showing only the first name of each user
		// defining a slice for the first names
		firstNames := []string{}
		// to iterate over bookings we need a range expression
		// range iterates over elements for different data structures
		// for slices and arrays it gives back the index and value for each element
		for _, booking := range bookings_slice {
			// splitting the string
			// strings.Fields() splits the string with whitespace as seperator
			// returns a slice with two elements
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all of our bookings (first names): %v\n", firstNames)

		// conditionals (if... else...)
		// expression that evaluates to either true or false
		// var noTicketsRemaining bool = remainingTickets == 0
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out, please come back next year!")
			break
		}

	}
}
