package main

import (
	"fmt"
	"strings"
)

// 1. The Product
type Pizza struct {
	Size     string
	Crust    string
	Sauce    string
	Toppings []string
	Cheese   bool
	Dough    string // Specific internal state
}

func (p Pizza) String() string {
	cheeseStatus := "No Cheese"
	if p.Cheese {
		cheeseStatus = "With Cheese"
	}
	return fmt.Sprintf("Pizza:\n  Size: %s\n  Crust: %s\n  Sauce: %s\n  Toppings: [%s]\n  %s\n  Dough: %s",
		p.Size, p.Crust, p.Sauce, strings.Join(p.Toppings, ", "), cheeseStatus, p.Dough)
}

// 2. The Builder
// In Go, the Builder is often a struct with methods that modify its internal state (the pizza being built)
// and return a pointer to itself for chaining.

type PizzaBuilder struct {
	pizza Pizza // The product instance being built
}

// Constructor for the builder

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{
		pizza: Pizza{
			Size:   "Medium", // Set some default values
			Crust:  "Regular",
			Sauce:  "Tomato",
			Cheese: true,
			Dough:  "Standard White",
		},
	}
}

// Building methods (fluent interface)

func (pb *PizzaBuilder) WithSize(size string) *PizzaBuilder {
	pb.pizza.Size = size
	return pb
}

func (pb *PizzaBuilder) WithCrust(crust string) *PizzaBuilder {
	pb.pizza.Crust = crust
	return pb
}

func (pb *PizzaBuilder) WithSauce(sauce string) *PizzaBuilder {
	pb.pizza.Sauce = sauce
	return pb
}

func (pb *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
	pb.pizza.Toppings = append(pb.pizza.Toppings, topping)
	return pb
}

func (pb *PizzaBuilder) WithCheese(hasCheese bool) *PizzaBuilder {
	pb.pizza.Cheese = hasCheese
	return pb
}

// The "build" method

func (pb *PizzaBuilder) Build() Pizza {
	// You can add validation here before returning the final product
	if pb.pizza.Size == "" || pb.pizza.Crust == "" || pb.pizza.Sauce == "" {
		fmt.Println("Warning: Building an incomplete pizza. Make sure Size, Crust, and Sauce are set.")
		// In a real application, you might return an error or a default valid pizza
	}
	return pb.pizza
}

// --- Main function to demonstrate usage ---
func main() {
	test := NewPizzaBuilder().Build()
	fmt.Println("--- Test pizza  ---")
	fmt.Println(test)

	// Build a simple pizza
	fmt.Println("\n--- Meat Lovers Pizza ---")
	// Build a meat lovers pizza
	fmt.Println("\n--- Veggie Delight (No Cheese) ---")
	// Build a veggie pizza without cheese
	veggieNoCheese := NewPizzaBuilder().
		WithSize("Medium").
		WithCrust("Whole Wheat").
		WithSauce("Pesto").
		AddTopping("Mushrooms").
		AddTopping("Onions").
		AddTopping("Bell Peppers").
		WithCheese(false). // Explicitly no cheese
		Build()
	fmt.Println(veggieNoCheese)

	fmt.Println("\n--- Small Default Pizza (Minimal Configuration) ---")
	// Build a pizza using mostly defaults, just changing size
	smallDefault := NewPizzaBuilder().
		WithSize("Small").
		Build()
	fmt.Println(smallDefault)
}
