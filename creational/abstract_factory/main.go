package main

import "fmt"

func getSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShirt := nikeFactory.makeShirt()
	nikeShoe := nikeFactory.makeShoe()

	adidasShirt := adidasFactory.makeShirt()
	adidasShoe := adidasFactory.makeShoe()

	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)

	printShirtDetails(adidasShirt)
	printShoeDetails(adidasShoe)
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s, Size: %d\n", s.getLogo(), s.getSize())
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s, Size: %d\n", s.getLogo(), s.getSize())
}
