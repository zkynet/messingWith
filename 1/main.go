package main

// notes
// How are variables allocated on the heap and stack
// different ways of declaring variables
// How to print variables
// ...

import (
	"log"
)

const (
	a = "a"
	b = "b"
	c = "c"
)

var (
	d = "d"
	e = "e"
)

func main() {
	f := "f"
	var g string = "g"
	// log.Println(&A)
	log.Println("f and G:", &f, &g)
	f = a
	log.Println("f as a:", f, &f, " <- value copied but the pointer remains the same")
	gg := &f
	log.Println("Pointing to F:", gg, &gg, *gg, " <- gg's value is a pointer to f's pointer. f's value can be returned via *gg")
	*gg = "FF"
	log.Println("New F value:", gg, &gg, *gg, " <- Value of f can be changed through *gg")
	gg = &g
	log.Println("GG pointing to g:", gg, &gg, *gg)

	// testing function return
	Name := NamedReturn()
	var Anon = AnonReturn()
	var Anon2 = AnonReturn()
	log.Println("Name inside Main():", Name, &Name, "<- stack // pre-allocated")
	log.Println("AnonReturn():", Anon, &Anon, "<- stack")
	log.Println("AnonReturn():", Anon2, &Anon2, "<- stack")
	log.Println("d inside main", d, &d, "<- heap")

	// testing local allocation
	local := LocalAllocation()
	log.Println("local in Main():", local, &local, " <- stack // pre-allocated ")

	// Testing allocation order
	before1 := "before1"
	before2 := "before2"
	log.Println("Before1 LocalAllocation():", before1, &before1)
	log.Println("Before2 LocalAllocation():", before2, &before2)

	// testing no escape
	LocalNoEscape()
	next := "next"
	log.Println("Next after LocalAllocation():", next, &next)
}

func NamedReturn() (name string) {
	name = "012345678"
	log.Println("name inside NamedReturn():", name, &name, "<- stack")
	return
}

func AnonReturn() string {
	return "012345"
}
func LocalAllocation() string {
	local := "local"
	log.Println("local inside LocalAllocation():", local, &local, "<- stack")
	log.Println("d inside LocalAllocation():", d, &d, "<- heap")
	return local
}

func LocalNoEscape() {
	onlyLocal := "only local"
	log.Println("onlyLocal inside LocalNoEscape():", onlyLocal, &onlyLocal)
}
