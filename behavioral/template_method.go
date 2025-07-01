package main

import "fmt"

// --- 1. Abstract Class / Component (Go equivalent: Interface for steps + orchestrator struct) ---

// HouseBuilder defines the common steps that vary between house types.
// These are the "abstract operations" or "primitive operations" that concrete builders will implement.
type HouseBuilder interface {
	BuildFoundation()
	BuildWalls()
	AddRoof()
	InstallFixtures() // A hook method (can be optional or have a default)
}

// ConstructionProcess orchestrates the overall building algorithm.
// This effectively acts as the "Template Method".
type ConstructionProcess struct {
	builder HouseBuilder // This holds the concrete implementation of the steps
}

// NewConstructionProcess creates a new process for a specific builder.
func NewConstructionProcess(builder HouseBuilder) *ConstructionProcess {
	return &ConstructionProcess{builder: builder}
}

// BuildHouse is the Template Method. It defines the fixed sequence of steps.
// Note: In Go, this method isn't "final" but relies on the interface methods
// being implemented by the `builder` field.
func (cp *ConstructionProcess) BuildHouse() {
	fmt.Println("\n--- Starting House Construction Process ---")
	cp.builder.BuildFoundation()
	cp.builder.BuildWalls()
	cp.builder.AddRoof()
	cp.builder.InstallFixtures() // Call the hook method
	fmt.Println("--- House Construction Finished! ---")
}

// --- 2. Concrete Classes ---

// WoodenHouseBuilder implements HouseBuilder for wooden houses.
type WoodenHouseBuilder struct{}

func (wb *WoodenHouseBuilder) BuildFoundation() {
	fmt.Println("  WoodenHouseBuilder: Laying a simple wooden foundation.")
}

func (wb *WoodenHouseBuilder) BuildWalls() {
	fmt.Println("  WoodenHouseBuilder: Erecting wooden walls.")
}

func (wb *WoodenHouseBuilder) AddRoof() {
	fmt.Println("  WoodenHouseBuilder: Installing a lightweight shingle roof.")
}

func (wb *WoodenHouseBuilder) InstallFixtures() {
	fmt.Println("  WoodenHouseBuilder: Installing standard wooden fixtures.") // Specific implementation
}

// BrickHouseBuilder implements HouseBuilder for brick houses.
type BrickHouseBuilder struct{}

func (bb *BrickHouseBuilder) BuildFoundation() {
	fmt.Println("  BrickHouseBuilder: Pouring a strong concrete foundation.")
}

func (bb *BrickHouseBuilder) BuildWalls() {
	fmt.Println("  BrickHouseBuilder: Laying sturdy brick walls.")
}

func (bb *BrickHouseBuilder) AddRoof() {
	fmt.Println("  BrickHouseBuilder: Adding heavy tiled roof.")
}

func (bb *BrickHouseBuilder) InstallFixtures() {
	fmt.Println("  BrickHouseBuilder: Installing premium marble fixtures.") // Different implementation
}

// --- Client Code ---
func main() {
	// Build a Wooden House
	fmt.Println("Client: Requesting a Wooden House.")
	woodenBuilder := &WoodenHouseBuilder{}
	woodenHouseProcess := NewConstructionProcess(woodenBuilder) // Client provides the specific builder
	woodenHouseProcess.BuildHouse()

	// Build a Brick House
	fmt.Println("\nClient: Requesting a Brick House.")
	brickBuilder := &BrickHouseBuilder{}
	brickHouseProcess := NewConstructionProcess(brickBuilder) // Client provides another specific builder
	brickHouseProcess.BuildHouse()
}
