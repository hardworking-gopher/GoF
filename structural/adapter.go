package main

import "fmt"

// --- 1. Target (Interface that the Client expects) ---
type DataProcessor interface {
	ProcessData(data string) (string, error)
}

// --- 3. Adaptee (The incompatible existing class) ---
// This is likely from an external library or legacy code that we cannot change.
type LegacyDataConverter struct{}

// ConvertData is the method with the incompatible signature/name
func (l *LegacyDataConverter) ConvertOldFormat(rawInput string) (string, error) {
	// Simulate complex legacy data conversion logic
	fmt.Printf("LegacyDataConverter: Converting '%s' from old format...\n", rawInput)
	converted := "OLD_FORMAT_CONVERTED_" + rawInput
	return converted, nil
}

// --- 4. Adapter ---
// This struct will implement the Target interface (DataProcessor)
// and wrap an instance of the Adaptee (LegacyDataConverter).
type LegacyDataConverterAdapter struct {
	legacyConverter *LegacyDataConverter // The adaptee instance
}

// NewLegacyDataConverterAdapter is a constructor for the adapter
func NewLegacyDataConverterAdapter(converter *LegacyDataConverter) *LegacyDataConverterAdapter {
	return &LegacyDataConverterAdapter{
		legacyConverter: converter,
	}
}

// ProcessData implements the DataProcessor interface.
// It translates the call to the adaptee's method.
func (a *LegacyDataConverterAdapter) ProcessData(data string) (string, error) {
	fmt.Printf("Adapter: Translating client's 'ProcessData' call to adaptee's 'ConvertOldFormat'...\n")
	// Here's the core of the adapter: calling the adaptee's method
	// and potentially translating parameters or return values.
	convertedData, err := a.legacyConverter.ConvertOldFormat(data)
	if err != nil {
		return "", fmt.Errorf("adapter error: %w", err)
	}
	return convertedData, nil
}

// --- 2. Client (Code that uses the Target interface) ---
// The client only knows about the DataProcessor interface.
func ClientCode(processor DataProcessor, input string) {
	fmt.Printf("\nClient: Sending '%s' for processing...\n", input)
	result, err := processor.ProcessData(input)
	if err != nil {
		fmt.Printf("Client: Error processing data: %v\n", err)
		return
	}
	fmt.Printf("Client: Received processed data: '%s'\n", result)
}

// --- Main function to demonstrate usage ---
func main() {
	// Scenario 1: Directly using the Adaptee (not compatible with ClientCode)
	// legacyConverter := &LegacyDataConverter{}
	// This would not compile: ClientCode(legacyConverter, "test") because LegacyDataConverter does not implement DataProcessor

	// Scenario 2: Using the Adapter to make the Adaptee compatible
	fmt.Println("--- Using the Adapter ---")
	legacyConverter := &LegacyDataConverter{}
	adapter := NewLegacyDataConverterAdapter(legacyConverter)

	// Now, the client code can use the adapter seamlessly
	ClientCode(adapter, "hello_world")

	fmt.Println("\n--- Another Use Case ---")
	ClientCode(adapter, "sample_data")
}
