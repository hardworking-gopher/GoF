package main

// Adidas Factory
type Adidas struct{}

func (a *Adidas) makeShirt() IShirt {
	return &Shirt{
		logo: "adidas",
		size: 15,
	}
}

func (a *Adidas) makeShoe() IShoe {
	return &Shoe{
		logo: "adidas",
		size: 15,
	}
}
