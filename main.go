package main

func main() {
	//Create objects
	game1 := game{
		title: "Minecraft",
		price: 5,
	}
	game2 := game{
		title: "World of Warcraft",
		price: 19,
	}
	game3 := game{
		title: "Elite Dangerous",
		price: 54,
	}
	book1 := book{
		title: "Candle in the tomb",
		price: 20,
	}
	book2 := book{
		title: "Barney and Friends",
		price: 10,
	}
	computerAccessories1 := computerAccessories{
		title: "Razer BT earpiece",
		price: 159,
	}
	computerAccessories2 := computerAccessories{
		title: "Razer Keyboard",
		price: 110,
	}
	computerAccessories3 := computerAccessories{
		title: "Logitech Mouse",
		price: 80,
	}

	var items []*game
	items = append(items, &game1, &game2, &game3)

	my := list(items)
	my.print()
}