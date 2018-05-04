package main

// This file shows a complex (and mostly irrelevant) point about
// interfaces. I lied that when an interface is stored, we store a
// pointer to the original value. We actually store a pointer to a
// *copy* of the original value.
//
// This matters because mutator methods won't change the original value,
// unless we store a *pointer* in the interface value.
//
// This is what I failed at in lecture.

import "fmt"

type cat struct {
	name string
	age  int
}

func (c cat) speak() {
	fmt.Printf("%v (age %v) goes meow!\n", c.name, c.age)
}

// A mutating method needs to have a pointer reciever. Otherwise a local
// copy of cat will be made and the copy is mutated rather than the
// original.
func (c *cat) ageUp(increment int) {
	c.age += increment
}

type animal interface {
	speak()
}

func main() {
	originalCat := cat{name: "Gizmo", age: 23}

	// Interface variable can store anything implementing the interface.
	var myAnimal animal

	// The interface stores:
	// (1) a pointer to the copy of the cat struct,
	// (2) remembers that we are pointing to a cat.
	myAnimal = originalCat
	var catCopy cat
	catCopy = myAnimal.(cat)
	catCopy.ageUp(123)
	// No effect on original cat!
	fmt.Printf("%#v\n", originalCat)

	// The interface stores:
	// (1) a pointer to a copy of catPointer (itself a pointer)
	// (2) remembers that we are pointing to a cat *pointer*.
	var originalCatPointer *cat
	originalCatPointer = &originalCat
	myAnimal = originalCatPointer

	var catPointerCopy *cat
	// NOTE: cast to regular cat would fail: myAnimal.(cat) would return
	// `ok = false`.
	catPointerCopy = myAnimal.(*cat)
	catPointerCopy.ageUp(123)
	// Changes original cat!
	fmt.Printf("%#v\n", originalCat)
}

// https://research.swtch.com/interfaces
