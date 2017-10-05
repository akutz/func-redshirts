/*
Package main illustrates the difference between func (r *redshirt) and
func (r redshirt).
*/
package main

import (
	"fmt"
)

// redshirt1 is a crew member that elects to beam copies of
// himself to receive data on away missions.
type redshirt1 struct {
	data string
}

// receive is defined as a function of redshirt1,
// therefore "r.data" inside the function refers to
// redshirt1, not an address of redshirt1.
//
// This means if redshirt1 is teleported and he receives
// data on an away mission, it's his copy that receives
// the data, not the original crew member back on the ship!
func (r redshirt1) receive(data string) {
	r.data = data
}

// redshirt2 opts for teleportation based on quantum-entanglement,
// and thus whatever happens to her on away missions is reflected
// back on the deck of the ship
type redshirt2 struct {
	data string
}

// receive is defined as a function on a pointer to redshirt2,
// therefore "r.data" inside the function refers to the address
// of redshirt2, not a copy
//
// This means if redshirt2 is teleported and she receives
// data on an away mission, the same data is now available
// back on the ship!
func (r *redshirt2) receive(data string) {
	r.data = data
}

type redshirt interface {
	receive(string)
}

func teleport(r redshirt) {
	r.receive("data")
}

func main() {

	// Assign a redshirt1 object to r1.
	r1 := redshirt1{}

	// Assign the address of a redshirt2 object to r2.
	r2 := &redshirt2{}

	// Print the data held by each of the redshirts.
	fmt.Printf("r1.data=%s\tr2.data=%s\n", r1.data, r2.data)

	// Teleport both of the redshirts down to the planet
	// so they can accept data from the native population.
	teleport(r1)
	teleport(r2)

	// Print the data held by each of the redshirts.
	//
	// The first redshirt will not be carrying any data,
	// because when he was teleported a copy of him was
	// made and beamed down to the planet's surface. It
	// was this copy that received the data, so the
	// original, back on the ship, has no data!
	//
	// The second redshirt opted to use the teleporter with
	// the new quantum-entanglement upgrade! Everything
	// that happened to her on the planet happened to her
	// at the same time on the ship! Therefore when she
	// received the data on the planet, she also received
	// the data back on the ship.
	fmt.Printf("r1.data=%s\tr2.data=%s\n", r1.data, r2.data)
}
