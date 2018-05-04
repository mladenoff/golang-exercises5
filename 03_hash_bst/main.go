package main

import "fmt"

type hashable interface {
	Hash() int
}

// This will store values in a BST, but it will sort things by their
// *hash value*. This is like the BST version of a hash map.
type treeNode struct {
	left  *treeNode
	value hashable
	right *treeNode
}

func insert(root *treeNode, value hashable) *treeNode {
	if root == nil {
		return &treeNode{value: value}
	}

	hash := value.Hash()

	if hash < root.value.Hash() {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}

	return root
}

// To add methods to the builtin string class, I must make an "alias"
// for it. A `hashableString` is really just a string at the end of the
// day.
type hashableString string

// I implement the method to make `hashableString` an implementor of the
// hashable interface.
func (s hashableString) Hash() int {
	var value int
	for idx := 0; idx < len(s); idx++ {
		byteValue := s[idx]
		// `value` is a 32bit integer. But byteValue is a byte and thus in
		// the range 0..255. If I combine value and byteValue without
		// shifting, I'll only ever use the least significant 8 bits of the
		// value.
		//
		// Shifting lets me create the 32 bit values:
		// * 00 00 00 bb
		// * 00 00 bb 00
		// * 00 bb 00 00
		// * bb 00 00 00
		shiftAmount := uint((idx % 4) * 8)
		var shiftedByteVal = int(byteValue) << shiftAmount
		// xor the value with the shifted bytes.
		value = value ^ shiftedByteVal
	}

	// Note: this hash function is really too simplistic to be very good.
	// But it works okay for my purposes.

	return value
}

func traverse(root *treeNode, f func(value hashable)) {
	if root == nil {
		return
	}

	traverse(root.left, f)
	f(root.value)
	traverse(root.right, f)
}

func main() {
	myStrings := [...]string{"abc", "xyz", "123", "456"}

	var root *treeNode
	for _, s := range myStrings {
		// You need to "convert" your `string`s into `hashableString`s.
		root = insert(root, hashableString(s))
	}

	traverse(root, func(h hashable) {
		// You can't say h.(string) because Golang will say that this is
		// "impossible", because string does not satisfy hashable.
		//
		// But since a hashableString is synonymous with string, we can use
		// it in the same ways we can use a regular string.
		hs, ok := h.(hashableString)
		if ok {
			fmt.Printf("val %v | hash %v\n", string(hs), h.Hash())
		} else {
			fmt.Println("WTF? I thought we only stored strings in here!")
		}
	})
}
