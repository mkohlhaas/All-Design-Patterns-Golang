package main

import "fmt"

//////////////// Interface ///////////////////////

type iGun interface {
	getName() string
	getPower() int
}

////////////// Our Product ///////////////////////

type gun struct {
	name  string
	power int
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power
}

///////////////// Factories //////////////////////

/////////////////// Ak47 /////////////////////////

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{gun{name: "AK47 gun", power: 4}}
}

///////////////// Maverick ///////////////////////

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{gun{name: "Maverick gun", power: 5}}
}

////////////////// GunType ////////////////////////

type GunType interface {
	isGun()
}

type gunType struct{ sz string }

func (g gunType) isGun() {}

func (g gunType) String() string {
	return g.sz
}

var (
	Ak47     GunType = gunType{"ak47"}
	Maverick GunType = gunType{"maverick"}
)

////////////////// Public API ////////////////////

func NewGunFactory(gunType GunType) (gun iGun) {
	switch gunType {
	case Ak47:
		gun = newAk47()
	case Maverick:
		gun = newMaverick()
	}
	return
}

/////////////////// Main /////////////////////////

func main() {
	    // Ak47
	ak47 := NewGunFactory(Ak47)
	printDetails(ak47)

	// Maverick
	maverick := NewGunFactory(Maverick)
	printDetails(maverick)
}

///////////////// Helper /////////////////////////

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
