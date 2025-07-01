package main

import "fmt"

// --- 5. Receiver ---
// The Light knows how to perform the actual actions (turn on/off).
type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Light is ON")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Light is OFF")
}

// --- 1. Command (Interface) ---
// Declares an interface for executing an operation.
type Command interface {
	Execute()
}

// --- 2. Concrete Commands ---

// TurnOnCommand encapsulates the request to turn the light on.
type TurnOnCommand struct {
	light *Light // Reference to the Receiver
}

func NewTurnOnCommand(light *Light) *TurnOnCommand {
	return &TurnOnCommand{light: light}
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn() // Delegates the actual action to the Receiver
}

// TurnOffCommand encapsulates the request to turn the light off.
type TurnOffCommand struct {
	light *Light // Reference to the Receiver
}

func NewTurnOffCommand(light *Light) *TurnOffCommand {
	return &TurnOffCommand{light: light}
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff() // Delegates the actual action to the Receiver
}

// --- 4. Invoker ---
// A simple remote control button that can execute any Command.
type RemoteControl struct {
	command Command // Holds a Command object
}

func (rc *RemoteControl) SetCommand(cmd Command) {
	rc.command = cmd
}

func (rc *RemoteControl) PressButton() {
	if rc.command != nil {
		fmt.Print("RemoteControl: Button pressed. ")
		rc.command.Execute() // Invokes the command without knowing its details
	} else {
		fmt.Println("RemoteControl: No command set for button.")
	}
}

// --- Client Code ---
func main() {
	// The Receiver: The actual light bulb.
	livingRoomLight := &Light{}

	// The Concrete Commands: Created by the client, linking a specific action to a specific receiver.
	turnOnLight := NewTurnOnCommand(livingRoomLight)
	turnOffLight := NewTurnOffCommand(livingRoomLight)

	// The Invoker: The remote control.
	remote := &RemoteControl{}

	// Client configures the remote control (Invoker) with different commands.
	// The remote (Invoker) doesn't know what a "Light" is or how to "TurnOn/Off".
	// It just knows how to "Execute" a Command.

	remote.SetCommand(turnOnLight)
	remote.PressButton() // Light is ON

	remote.SetCommand(turnOffLight)
	remote.PressButton() // Light is OFF

	// You could even set the same command multiple times
	remote.SetCommand(turnOnLight)
	remote.PressButton() // Light is ON
}
