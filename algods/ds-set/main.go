package main

import "githug.com/fatih/set"

func main() {
	// create a set with zero items
	s := set.New(set.ThreadSafe) // thread safe version
	// add items
	s.Add("istanbul")
	s.Add("istanbul") // nothing happens if you add duplicate item

	// add multiple items
	s.Add("ankara", "san francisco", 3.14)

	// remove item
	s.Remove("frankfurt")
	s.Remove("frankfurt") // nothing happens if you remove a nonexisting item

	// remove multiple items
	s.Remove("barcelona", 3.14, "ankara")

	// removes an arbitary item and return it
	item := s.Pop()

	// create a new copy
	other := s.Copy()

	// remove all items
	s.Clear()

	// number of items in the set
	len := s.Size()

	// return a list of items
	items := s.List()

	// string representation of set
	fmt.Printf("set is %s", s.String())
}
