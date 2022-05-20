//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting(name string) string{
	return "Hello " + name 
}

func returnMessage() string {
	return "You're silly"
}

func add(a, b, c int) int {
	return a + b + c
}

func return3Numbers()(int , int, int){
	return 5,4,3
} 

func main() {
	sayHello := greeting("Jessica")
	fmt.Println(sayHello)

	getReturnMessage := returnMessage()
	fmt.Println(getReturnMessage)

	sums := add(4,2,4)
	fmt.Println(sums)

	num1 , num2, num3 := return3Numbers()
	fmt.Println(num1, num2, num3)
}
