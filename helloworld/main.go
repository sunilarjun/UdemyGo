package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

/*
How do we run the code in our project?
Using go CLI:
go run filename example: go run main.go || go build filename (builds file but does not execute, creates executbale however)

What does 'package main' mean?
Types of packages: executable or reusable (like dependencies ie helper code).
Specifically the word main is used to create executable packages.
Main is sacred word, any other word won't create an executable. Any executable must also have a function called main.

What does 'import "fmt"' mean?
Importing the fmt standard library package into package main.
This library is included in go by default.
One can import many packages either in the standard library or helper code from other engineers.

What's that 'func' thing?
Main function that takes in arguments and has function body below between curly braces.

How is the main.go file organized?
package main					-- 		package declaration

import "fmt"					--		import other packages that we need

func main() {
	fmt.Println("Hi there!")	--		declare functions, tell go to do things
}
*/
