package main

import "fmt"
import "sync" // For thread-safe observer list, important in real-world scenarios

// --- 2. Observer (Interface) ---
// Defines the update interface for objects that want to be notified.
type WeatherDisplay interface {
	Update(temperature, humidity, pressure float64)
	GetName() string // To identify displays
}

// --- 1. Subject (Interface) ---
// Defines the interface for objects that can be observed.
type WeatherStationSubject interface {
	RegisterObserver(display WeatherDisplay)
	DeregisterObserver(display WeatherDisplay)
	NotifyObservers()
}

// --- 4. Concrete Observer ---
// A concrete display unit that implements the WeatherDisplay interface.
type CurrentConditionsDisplay struct {
	name        string
	temperature float64
	humidity    float64
	// pressure float64 // Could store this too
}

func NewCurrentConditionsDisplay(name string) *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{name: name}
}

func (c *CurrentConditionsDisplay) GetName() string {
	return c.name
}

func (c *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
	c.temperature = temperature
	c.humidity = humidity
	fmt.Printf("[%s] Current conditions: %.1fF degrees and %.1f%% humidity\n", c.name, temperature, humidity)
}

// Another concrete observer
type ForecastDisplay struct {
	name        string
	temperature float64
	pressure    float64
}

func NewForecastDisplay(name string) *ForecastDisplay {
	return &ForecastDisplay{name: name}
}

func (f *ForecastDisplay) GetName() string {
	return f.name
}

func (f *ForecastDisplay) Update(temperature, humidity, pressure float64) {
	f.temperature = temperature
	f.pressure = pressure
	fmt.Printf("[%s] Forecast: Temperature around %.1fF, Pressure: %.1f hPa\n", f.name, temperature, pressure)
	// Some complex forecast logic based on values...
}

// --- 3. Concrete Subject ---
// The actual weather station that observes conditions.
type WeatherStation struct {
	observers   map[string]WeatherDisplay // Using a map for easy deregistration by name/ID
	temperature float64
	humidity    float64
	pressure    float64
	mu          sync.Mutex // For thread safety when modifying observer list
}

func NewWeatherStation() *WeatherStation {
	return &WeatherStation{
		observers: make(map[string]WeatherDisplay),
	}
}

func (ws *WeatherStation) RegisterObserver(display WeatherDisplay) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.observers[display.GetName()] = display
	fmt.Printf("WeatherStation: Registered %s\n", display.GetName())
}

func (ws *WeatherStation) DeregisterObserver(display WeatherDisplay) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	delete(ws.observers, display.GetName())
	fmt.Printf("WeatherStation: Deregistered %s\n", display.GetName())
}

func (ws *WeatherStation) NotifyObservers() {
	ws.mu.Lock() // Lock to prevent changes to observers map during iteration
	defer ws.mu.Unlock()

	fmt.Println("WeatherStation: Notifying observers...")
	for _, obs := range ws.observers {
		obs.Update(ws.temperature, ws.humidity, ws.pressure)
	}
}

// Method to simulate state change in the Subject
func (ws *WeatherStation) SetMeasurements(temperature, humidity, pressure float64) {
	ws.temperature = temperature
	ws.humidity = humidity
	ws.pressure = pressure
	fmt.Println("\nWeatherStation: New measurements received.")
	ws.NotifyObservers() // Notify all registered observers
}

// --- Client Code ---
func main() {
	// Create the Subject
	weatherStation := NewWeatherStation()

	// Create Concrete Observers
	currentDisplay1 := NewCurrentConditionsDisplay("Living Room Display")
	forecastDisplay1 := NewForecastDisplay("Kitchen Forecast")
	currentDisplay2 := NewCurrentConditionsDisplay("Bedroom Display")

	// Register observers with the subject
	weatherStation.RegisterObserver(currentDisplay1)
	weatherStation.RegisterObserver(forecastDisplay1)
	weatherStation.RegisterObserver(currentDisplay2)

	// Simulate weather changes - observers get notified automatically
	weatherStation.SetMeasurements(80, 65, 30.4)
	weatherStation.SetMeasurements(82, 70, 29.2)

	// Deregister an observer
	weatherStation.DeregisterObserver(currentDisplay2)

	// Simulate another weather change - only remaining observers get notified
	weatherStation.SetMeasurements(78, 90, 29.8)

	// Try to deregister an observer that's already gone
	weatherStation.DeregisterObserver(currentDisplay2) // Will show no effect as it's already deleted
}
