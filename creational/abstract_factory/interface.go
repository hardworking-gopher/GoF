package main

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type ISportsFactory interface {
	makeShirt() IShirt
	makeShoe() IShoe
}
