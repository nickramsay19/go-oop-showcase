package main

import (
	"fmt"
)

func main() {

	var unicorn Creature = *NewCreature("billy", false)

	fmt.Println(unicorn.ToString())
}
