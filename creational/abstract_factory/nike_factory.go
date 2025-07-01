package main


// Nike Factory
type Nike struct{}

func (n *Nike) makeShirt() IShirt {
	return &Shirt{
		logo: "nike",
		size: 14,
	}
}

func (n *Nike) makeShoe() IShoe {
	return &Shoe{
		logo: "nike",
		size: 14,
	}
}
