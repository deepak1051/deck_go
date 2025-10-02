package main

import "fmt"

// cretae a new type of deck whick is a
// slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
