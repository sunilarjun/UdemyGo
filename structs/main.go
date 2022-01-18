package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//option 1 struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//option 2 struct, most common
	//bigmoney := person{firstName: "Big", lastName: "Money"} initial method of declaring a struct
	//fmt.Println(bigmoney)

	//option 3 struct, here struct is created with default zero value for each string: ""
	//var bigmoney person

	//bigmoney.firstName = "Big"
	//bigmoney.lastName = "Money"

	//fmt.Println(bigmoney)
	//fmt.Printf("%+v", bigmoney)

	//new lesson from here

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//pointer stuff, pass by value used by go for mem
	//& - give me mem address where value of jim is
	//jimPointer := &jim
	jim.updateName("jimmy")
	jim.print()
}

//*person is a pointer to a type description
//go shortcut, turn var type person into ptr person for receiver type ptr
func (pointerToPerson *person) updateName(newFirstName string) {
	//* - give me direct access to the value at the mem address, this pointer is a regular operator
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*
slice vs array

a slice is a slice ds and array ds

slice is a ptr to head of array created, capacity, and length
*/
