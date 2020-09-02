package main

import (
	"fmt"
)

/*
	--- Creature Type ---
*/
type Creature struct {
	Name string
	Real bool
}

// NewCreature --> Creature pointer
func NewCreature(name string, real bool) *Creature {
	newCreature := new(Creature)
	newCreature.Name = name
	newCreature.Real = real

	return newCreature
}
// makeCreature --> Creature object
func makeCreature(name string, real bool) Creature {
	var newCreature Creature
	newCreature.Name = name
	newCreature.Real = real

	return newCreature
}

// ToString --> String representation of Creature
func (c Creature) ToString() string {
	return fmt.Sprintf("Name: '%s', Real: %t\n", c.Name, c.Real)
}