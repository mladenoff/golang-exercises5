package main

// Methods are just like functions, except there is a special "reciever"
// argument. This is written in parens before the method name.

import "fmt"

type cat struct {
	name string
	age  int
}

func (c cat) purr() {
	fmt.Println("PURRING")
}

func (c cat) speak() {
	fmt.Printf("%v (age %v) goes meow!\n", c.name, c.age)
}

// Won't work.
// func (c cat) ageUp(increment int) {
// 	c.age += increment
// }

type dog struct {
	name  string
	age   int
	skill string
}

// A struct "satisfies" an interface if it has all the methods declared
// in the interface. An interface is *not* a base class; it has little
// to nothing to do with inheritance. Golang does not have inheritance.
type animal interface {
	speak()
}

func (d dog) speak() {
	fmt.Printf("%v BOW WOW WOW\n", d.name)
}

func main() {
	c := cat{name: "Gizmo", age: 23}
	d := dog{name: "Henry", age: 2, skill: "happiness"}

	// Interface variable can store anything implementing the interface.
	var myAnimal animal
	myAnimal = c
	myAnimal.speak()
	myAnimal = d
	myAnimal.speak()

	// Note: `myAnimal` can hold either a cat (2 cells of memory) or a dog
	// (3 cells of memory), because it really stores a *pointer* to the
	// object (always one word of memory).
	//
	// The variable `myAnimal` also stores the type of the thing that is
	// stored. That way, when we call myAnimal.speak(), we first see what
	// kind of thing is stored (a Cat or a Dog), and then call the
	// appropriate method on the data pointed to by the variable.
	//
	// Therefore an interface is of the form `(pointer, type)`.

	// You can have an array (or slice) of interface values.
	animals := [2]animal{}
	animals[0] = c
	animals[1] = d
	// animals := []animal{}
	// animals = append(animals, c)
	// animals = append(animals, d)

	// for idx := 0; idx < len(animals); idx++ {
	// 	animal := animals[idx]
	// 	animal.speak()
	// }

	for _, animal := range animals {
		// Speak every animal.
		// AND purr every cat...
		animal.speak()

		var catCopy cat
		var ok bool
		// This tries to "downcast" the animal into a cat. If successful,
		// we can call cat methods. Else the catCopy variable is left with
		// the original zero value.
		catCopy, ok = animal.(cat)
		if ok {
			catCopy.purr()
		} else {
			fmt.Println("THIS WAS NOT A CAT")
			fmt.Printf("%#v\n", catCopy)
		}
	}

	// The "zero value" of an interface is (nil, nil): pointer to no data,
	// no type. This is different from the zero value of a cat, which was
	// a cat with zero value string ("") and zero value age (0).
	var uninitializedAnimal animal
	// A call to animal.speak() will crash the program. This is a "null
	// pointer exception" AKA "nil pointer dereference" AKA "segmentation
	// fault".
	uninitializedAnimal.speak()

	// BTW, this *doesn't* cause any crash, because it creates a real cat
	// value.
	var uninitializedCat cat
	uninitializedCat.speak()

	// Does crash, because cat pointer zero value is the null pointer
	// (pointer to memory address zero, which is invalid).
	var uninitializedCatPointer *cat
	uninitializedCatPointer.speak()
}

// More gory details:
// https://research.swtch.com/interfaces
