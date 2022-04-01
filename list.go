package main

type printer interface {
	print()
}

//declare List
type list []printer

//declare method for List Print
func (l list) print() {
	for _, item := range l {
		item.print()
	}
}
