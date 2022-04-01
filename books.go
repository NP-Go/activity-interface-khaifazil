package main

import "fmt"

//Declare books Struct
type book struct {
	title string
	price int
}

//Create print method
func (b *book) print() {
	fmt.Println(b.title, "$", b.price)
}
