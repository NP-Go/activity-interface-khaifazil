package main

import "fmt"

//Declare Computer Accessoies Struct
type computerAccessories struct {
	title string
	price int
}

//Create print method
func (c *computerAccessories) print() {
	fmt.Println(c.title, "$", c.price)
}
