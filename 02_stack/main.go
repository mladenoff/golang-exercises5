package main

import "fmt"

// Not typical or necessary. Can just say interface{} everywhere.
type blankInterface interface {
}

// My stack can store any kind of thing.
type stack struct {
	store []blankInterface
}

func (s *stack) push(i blankInterface) {
	s.store = append(s.store, i)
}

func (s *stack) pop() blankInterface {
	if len(s.store) == 0 {
		return -1
	}

	val := s.store[len(s.store)-1]
	s.store = s.store[0:(len(s.store) - 1)]
	return val
}

func main() {
	var s stack

	// I can store a mix of stuff.
	s.push(123)
	s.push("345")
	s.push(1.2345)

	fmt.Printf("%#v\n", s.pop())
	fmt.Printf("%#v\n", s.pop())
	fmt.Printf("%#v\n", s.pop())

	// If I only store ints, it's kind of annoying because Golang doesn't
	// know for sure that I'm using `s` just to store ints. I still have
	// to explicitly cast when I pop, so that I can use operations like
	// `+` on the items.
	s.push(123)
	s.push(456)
	s.push(789)

	val1 := s.pop()
	num1 := val1.(int)
	val2 := s.pop()
	num2 := val2.(int)
	// Shorter
	num3 := s.pop().(int)
	sum := num1 + num2 + num3

	fmt.Println(sum)
}
