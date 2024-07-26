package day02

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatingAssigningAndDereferencing() {
	// Creating a variable x of type int and assign it a value
	x := 3

	// Print the address of x using the & operator
	fmt.Println("The address of x is:", &x)

	// Declare a pointer variable p that points to an int and assign it the address of x
	p := &x

	// Print the value of x by dereferencing p
	fmt.Println("The value of x is:", *p)
}

func ModifyingValuesThroughPointers() {
	// Declare a variable age of type int
	var age int

	// Write a function increment that takes a pointer to an int and increments the value
	// it points to by 1.
	increment := func(value *int) {
		*value += 1
	}

	// Print the value of age before the call to increment
	fmt.Println("Age before increment is:", age)

	// Call increment passing the address of age
	increment(&age)

	// Print the value of age after the call to increment
	fmt.Println("Age after increment is:", age)
}

func RPGArena1() {
	// define utils
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	didHit := func(generator *rand.Rand) bool {
		return (generator.Intn(20) + 1) > 11
	}
	doDamage := func(i int, knight1 *int, knight2 *int, generator *rand.Rand) {
		if i == 1 {
			*knight2 -= (generator.Intn(6) + 1)
			fmt.Println("Knight1 Hit! Knight2's health is now:", *knight2)
		} else if i == 2 {
			*knight1 -= (generator.Intn(6) + 1)
			fmt.Println("Knight2 Hit! Knight1's health is now:", *knight1)
		} else {
			panic("Damage done to undefined player!")
		}
	}

	fmt.Println("Welcome to RPG Arena simulator, where we simulate combat betweeen two noble knights!")
	fmt.Println()

	knight1, knight2 := 20, 20

main:
	for {
		fmt.Println("Health of Knight 1:", knight1)
		fmt.Println("Health of Knight 2:", knight2)
		fmt.Println()

		for i := 1; i <= 2; i++ {
			if didHit(generator) {
				doDamage(i, &knight1, &knight2, generator)
				if knight1 <= 0 {
					fmt.Println()
					fmt.Println("Knight2 wins!")
					break main
				} else if knight2 <= 0 {
					fmt.Println()
					fmt.Println("Knight1 wins!")
					break main
				}
			} else {
				fmt.Printf("Knight%d Missed!\n", i)
			}
		}

		fmt.Println()
	}
}
