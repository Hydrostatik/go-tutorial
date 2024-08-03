package main

import (
	"fmt"
	bl "rpg_arena/business_layer"
)

func main() {
	character1 := bl.Character{
		Name:       "Edgar the Fencer",
		Profession: "Fencer",
		Health:     30,
	}

	character2 := bl.Character{
		Name:       "John the Knight",
		Profession: "Knight",
		Health:     45,
	}

	for range 10 {
		character1.Attack(&character2)
		character2.Attack(&character1)
	}

	fmt.Printf("Health of %s: %d\n", character1.Name, character1.Health)
	fmt.Printf("Health of %s: %d\n", character2.Name, character2.Health)
}
