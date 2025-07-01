package main

import "fmt"

// --- 1. Component (Interface) ---
// Defines the common interface for coffee and its condiments.
type Coffee interface {
	GetCost() float64
	GetDescription() string
}

// --- 2. Concrete Component ---
// The basic coffee object.
type SimpleCoffee struct{}

func (c *SimpleCoffee) GetCost() float64 {
	return 2.0 // Base cost of coffee
}

func (c *SimpleCoffee) GetDescription() string {
	return "Simple Coffee"
}

// --- 4. Concrete Decorators ---

// Milk decorator
type Milk struct {
	Coffee // Embeds the base decorator for delegation
}

func NewMilk(c Coffee) Coffee {
	return &Milk{Coffee: c}
}

func (m *Milk) GetCost() float64 {
	return m.Coffee.GetCost() + 0.5 // Add milk cost
}

func (m *Milk) GetDescription() string {
	return m.Coffee.GetDescription() + ", Milk" // Add milk description
}

// Sugar decorator
type Sugar struct {
	Coffee
}

func NewSugar(c Coffee) Coffee {
	return &Sugar{Coffee: c}
}

func (s *Sugar) GetCost() float64 {
	return s.Coffee.GetCost() + 0.2 // Add sugar cost
}

func (s *Sugar) GetDescription() string {
	return s.Coffee.GetDescription() + ", Sugar" // Add sugar description
}

// Caramel decorator
type Caramel struct {
	Coffee
}

func NewCaramel(c Coffee) Coffee {
	return &Caramel{Coffee: c}
}

func (ca *Caramel) GetCost() float64 {
	return ca.Coffee.GetCost() + 0.75 // Add caramel cost
}

func (ca *Caramel) GetDescription() string {
	return ca.Coffee.GetDescription() + ", Caramel" // Add caramel description
}

// --- Client Code ---
func main() {
	// Start with a simple coffee
	myCoffee := &SimpleCoffee{}
	fmt.Printf("Base Coffee: %s - $%.2f\n", myCoffee.GetDescription(), myCoffee.GetCost())

	fmt.Println("\n--- Adding Condiments ---")

	// Add Milk
	milkCoffee := NewMilk(myCoffee)
	fmt.Printf("Milk Coffee: %s - $%.2f\n", milkCoffee.GetDescription(), milkCoffee.GetCost())

	// Add Sugar to the milk coffee
	milkSugarCoffee := NewSugar(milkCoffee)
	fmt.Printf("Milk & Sugar Coffee: %s - $%.2f\n", milkSugarCoffee.GetDescription(), milkSugarCoffee.GetCost())

	// Start fresh, add caramel and then milk
	caramelCoffee := NewCaramel(&SimpleCoffee{}) // Start with a new simple coffee
	caramelMilkCoffee := NewMilk(caramelCoffee)
	fmt.Printf("Caramel & Milk Coffee: %s - $%.2f\n", caramelMilkCoffee.GetDescription(), caramelMilkCoffee.GetCost())

	fmt.Println("\n--- Complex Order ---")
	// Order a coffee with milk, sugar, and extra caramel
	complexCoffee := NewCaramel(NewSugar(NewMilk(&SimpleCoffee{})))
	fmt.Printf("Complex Coffee: %s - $%.2f\n", complexCoffee.GetDescription(), complexCoffee.GetCost())

	// Client code always interacts with the `Coffee` interface,
	// regardless of how many decorators are applied.
	describeAndCost(complexCoffee)
}

func describeAndCost(c Coffee) {
	fmt.Printf("Final Order: %s | Total Cost: $%.2f\n", c.GetDescription(), c.GetCost())
}
