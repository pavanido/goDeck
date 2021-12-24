package apiserver

import "fmt"

type deck []string

func (d deck) PrintCards() {
	for _, card := range d {
		fmt.Println("card data+", card)
	}
}
