package day03

import "fmt"

// Go Struct Excercises

// Calculate Rectangle Area and Perimeter

// 1. Create a struct named Rectangle with fields Length and Width (both float64).

type Rectangle struct {
	Length float64
	Width  float64
}

// 2. Write methods Area() and Perimiter() that calculate the area and perimeter of the rectangle, respectively.

func (x Rectangle) Area() float64 {
	return x.Length * x.Width
}

func (x Rectangle) Perimeter() float64 {
	return 2 * (x.Length + x.Width)
}

// 3. Create a Rectangle variable, initialize it with values, print its area and perimeter using the methods.

func RectanglePlayground() {
	rectangle := Rectangle{
		Length: 24,
		Width:  46,
	}

	fmt.Printf(
		"The Area and Perimeter of a Rectangle with length: %0.2f, and width: %0.2f is %0.2f, and %0.2f, respectively. \n",
		rectangle.Length,
		rectangle.Width,
		rectangle.Area(),
		rectangle.Perimeter())
}

// Embedded Structs and Composition

// 1. Create a struct Address with fields like Street, City and State.

type Address struct {
	Street string
	City   string
	State  string
	Person
}

// 2. Create a struct Person that has a FirstName and LastName. It should be embedded in the Address struct

type Person struct {
	FirstName string
	LastName  string
}

// 3. Create a NewAddress() that accepts values for the fields in both Address and Person.
// NewAddress() instantiates an Address with the embedded Person and fills in the appropriate
// fields with values passed in.

func NewAddress(street string, city string, state string, firstName string, lastName string) Address {
	return Address{
		Street: street,
		City:   city,
		State:  state,
		Person: Person{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
}

// 4. Create a method called ShowAddress() that receives an Address instance and have it print the
// first name, last name, street, city and state

func ShowAddress(address Address) {
	fmt.Printf(
		"Address: %s, %s, %s for Person: %s %s\n",
		address.Street,
		address.City,
		address.State,
		address.FirstName,
		address.LastName)
}

// 5. Call NewAddress(), receive the instance and pass it to showAddress()

func AddressPlayground() {
	ShowAddress(NewAddress("570 Randall Mill Street", "Falls Church", "VA", "Mickey", "Gopher"))
}
