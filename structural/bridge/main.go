package main

import "fmt"

// --- 3. Implementor (Interface) ---
// Defines the interface for the implementation hierarchy.
// These are the low-level operations that different devices can perform.
type Device interface {
	IsOn() bool
	Enable()
	Disable()
	GetVolume() int
	SetVolume(percent int)
	GetChannel() int
	SetChannel(channel int)
	GetName() string // To distinguish devices
}

// --- 4. Concrete Implementors (Implement Device interface) ---

// TV is a concrete implementation of Device.
type TV struct {
	on      bool
	volume  int
	channel int
}

func (t *TV) IsOn() bool        { return t.on }
func (t *TV) Enable()           { t.on = true; fmt.Println("TV is ON") }
func (t *TV) Disable()          { t.on = false; fmt.Println("TV is OFF") }
func (t *TV) GetVolume() int    { return t.volume }
func (t *TV) SetVolume(percent int) {
	if percent < 0 { percent = 0 }
	if percent > 100 { percent = 100 }
	t.volume = percent
	fmt.Printf("TV Volume set to %d%%\n", t.volume)
}
func (t *TV) GetChannel() int   { return t.channel }
func (t *TV) SetChannel(channel int) {
	t.channel = channel
	fmt.Printf("TV Channel set to %d\n", t.channel)
}
func (t *TV) GetName() string { return "TV" }


// Radio is another concrete implementation of Device.
type Radio struct {
	on      bool
	volume  int
	channel int // Represents frequency or station
}

func (r *Radio) IsOn() bool        { return r.on }
func (r *Radio) Enable()           { r.on = true; fmt.Println("Radio is ON") }
func (r *Radio) Disable()          { r.on = false; fmt.Println("Radio is OFF") }
func (r *Radio) GetVolume() int    { return r.volume }
func (r *Radio) SetVolume(percent int) {
	if percent < 0 { percent = 0 }
	if percent > 100 { percent = 100 }
	r.volume = percent
	fmt.Printf("Radio Volume set to %d%%\n", r.volume)
}
func (r *Radio) GetChannel() int   { return r.channel }
func (r *Radio) SetChannel(channel int) {
	r.channel = channel
	fmt.Printf("Radio Frequency set to %d MHz\n", r.channel)
}
func (r *Radio) GetName() string { return "Radio" }


// --- 1. Abstraction (Interface) ---
// Defines the interface for the abstraction hierarchy.
// It holds a reference to the Implementor.
type RemoteControl interface {
	TogglePower()
	VolumeDown()
	VolumeUp()
	ChannelDown()
	ChannelUp()
	GetDeviceName() string
}

// BaseRemote is a common struct that can be embedded by Refined Abstractions.
// It holds the Device (Implementor) reference.
type BaseRemote struct {
	device Device // This is the "bridge"
}

// --- 2. Refined Abstractions (Implement RemoteControl interface) ---

// BasicRemote implements the RemoteControl interface using a Device.
type BasicRemote struct {
	BaseRemote // Embed the common fields/methods
}

func NewBasicRemote(device Device) *BasicRemote {
	return &BasicRemote{BaseRemote{device: device}}
}

func (br *BasicRemote) TogglePower() {
	fmt.Printf("Basic Remote (%s): ", br.device.GetName())
	if br.device.IsOn() {
		br.device.Disable()
	} else {
		br.device.Enable()
	}
}

func (br *BasicRemote) VolumeDown() {
	fmt.Printf("Basic Remote (%s): ", br.device.GetName())
	br.device.SetVolume(br.device.GetVolume() - 10)
}

func (br *BasicRemote) VolumeUp() {
	fmt.Printf("Basic Remote (%s): ", br.device.GetName())
	br.device.SetVolume(br.device.GetVolume() + 10)
}

func (br *BasicRemote) ChannelDown() {
	fmt.Printf("Basic Remote (%s): ", br.device.GetName())
	br.device.SetChannel(br.device.GetChannel() - 1)
}

func (br *BasicRemote) ChannelUp() {
	fmt.Printf("Basic Remote (%s): ", br.device.GetName())
	br.device.SetChannel(br.device.GetChannel() + 1)
}

func (br *BasicRemote) GetDeviceName() string {
	return br.device.GetName()
}


// AdvancedRemote extends BasicRemote with more functionality.
type AdvancedRemote struct {
	BasicRemote // Embed the basic remote functionality
}

func NewAdvancedRemote(device Device) *AdvancedRemote {
	return &AdvancedRemote{BasicRemote: BaseRemote{device: device}}
}

func (ar *AdvancedRemote) Mute() {
	fmt.Printf("Advanced Remote (%s): Muting... ", ar.device.GetName())
	ar.device.SetVolume(0)
}

// Since it embeds BasicRemote, it already has TogglePower, VolumeUp/Down, ChannelUp/Down methods.
// We can also override them if needed, but here we just add Mute.


// --- Client Code ---
func main() {
	// Create devices (Implementors)
	myTV := &TV{}
	myRadio := &Radio{}

	// Create remotes (Abstractions) and bridge them to devices
	fmt.Println("--- Using Basic Remote with TV ---")
	basicTVRemote := NewBasicRemote(myTV)
	basicTVRemote.TogglePower()
	basicTVRemote.VolumeUp()
	basicTVRemote.ChannelUp()
	basicTVRemote.TogglePower()

	fmt.Println("\n--- Using Basic Remote with Radio ---")
	basicRadioRemote := NewBasicRemote(myRadio)
	basicRadioRemote.TogglePower()
	basicRadioRemote.VolumeUp()
	basicRadioRemote.SetChannel(98) // BasicRemote doesn't have a direct SetChannel method, but Device does
	basicRadioRemote.ChannelDown()
	basicRadioRemote.TogglePower()

	fmt.Println("\n--- Using Advanced Remote with TV ---")
	advancedTVRemote := NewAdvancedRemote(myTV)
	advancedTVRemote.TogglePower()
	advancedTVRemote.VolumeUp()
	advancedTVRemote.Mute() // Advanced remote's unique feature
	advancedTVRemote.TogglePower()
}