# Go - The Complete Guide

This is my attempt to get a good grasp on `Go` by going through the Udemy course.

## Notes

* The entry point package must be named main and so must the entrypoint function for the application
* You can only have one main function in your application
* `package main` on top of a file determines which package the file belongs to
* For libraries, you're not required to have a main function
* You must have at least one package for your project
* There's a difference between modules and packages
  * modules represent an application in its totality
  * `go mod init` creates a `go.mod` to manage the module
* You can explicity assign a type to a variable using `var someInt int = 6`

### Functions

* If a function will be used outside the package use `PascalCase` otherwise use `camelCase`

### Built-in types

* int -> A number without decimal places
* float64 -> 64 bit number with decimal places
* float32 -> 32 bit number with decimal places
* string -> A text value
* bool -> boolean value can be `true` or `false`

Other types: uint, int32, rune (int32 alias of unicode code point), uint32, int64, int8, uint8

### Null Values

All Go value types come with a default value, also known as a "null value" which is a value
stored in a variable if no other value is explicitly set

Some examples:

* int -> 0
* float64 -> 0.0
* string -> ""
* bool -> false

## Questions

- Do packages need to follow a folder structure?