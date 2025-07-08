package main

import "fmt"

// 2. Concrete Products
type EmailNotifier struct {
	Recipient string
}

func (e *EmailNotifier) Send(message string) error {
	fmt.Printf("Sending Email to %s: %s\n", e.Recipient, message)
	// Simulate email sending logic
	return nil
}

type SMSNotifier struct {
	PhoneNumber string
}

func (s *SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
	// Simulate SMS sending logic
	return nil
}

type PushNotifier struct {
	DeviceToken string
}

func (p *PushNotifier) Send(message string) error {
	fmt.Printf("Sending Push Notification to %s: %s\n", p.DeviceToken, message)
	// Simulate push notification logic
	return nil
}

// --- The Factory (Creator in Go often a function) ---

// 1. Product (Interface)
type Notifier interface {
	Send(message string) error
}

// NotificationType is an enum-like string to specify the type of notifier
type NotificationType string

const (
	TypeEmail NotificationType = "email"
	TypeSMS   NotificationType = "sms"
	TypePush  NotificationType = "push"
)

// NewNotifier is the "Factory Method" function.
// It takes a type and configuration, and returns the appropriate Notifier interface.
func NewNotifier(nt NotificationType, config map[string]string) (Notifier, error) {
	switch nt {
	case TypeEmail:
		recipient, ok := config["recipient"]
		if !ok {
			return nil, fmt.Errorf("recipient not provided for email notifier")
		}
		return &EmailNotifier{Recipient: recipient}, nil
	case TypeSMS:
		phoneNumber, ok := config["phone_number"]
		if !ok {
			return nil, fmt.Errorf("phone_number not provided for SMS notifier")
		}
		return &SMSNotifier{PhoneNumber: phoneNumber}, nil
	case TypePush:
		deviceToken, ok := config["device_token"]
		if !ok {
			return nil, fmt.Errorf("device_token not provided for push notifier")
		}
		return &PushNotifier{DeviceToken: deviceToken}, nil
	default:
		return nil, fmt.Errorf("unknown notification type: %s", nt)
	}
}

// --- Client Code ---
func main() {
	// Create an Email Notifier
	emailConfig := map[string]string{"recipient": "user@example.com"}
	emailNotifier, err := NewNotifier(TypeEmail, emailConfig)
	if err != nil {
		fmt.Println("Error creating email notifier:", err)
	} else {
		emailNotifier.Send("Hello via Email!")
	}

	fmt.Println()

	// Create an SMS Notifier
	smsConfig := map[string]string{"phone_number": "123-456-7890"}
	smsNotifier, err := NewNotifier(TypeSMS, smsConfig)
	if err != nil {
		fmt.Println("Error creating SMS notifier:", err)
	} else {
		smsNotifier.Send("Hello via SMS!")
	}

	fmt.Println()

	// Create a Push Notifier
	pushConfig := map[string]string{"device_token": "abcdef123456"}
	pushNotifier, err := NewNotifier(TypePush, pushConfig)
	if err != nil {
		fmt.Println("Error creating push notifier:", err)
	} else {
		pushNotifier.Send("Hello via Push!")
	}

	fmt.Println()

	// Try to create an unknown notifier
	unknownNotifier, err := NewNotifier("unknown_type", nil)
	if err != nil {
		fmt.Println("Error creating unknown notifier:", err)
	} else {
		unknownNotifier.Send("This won't happen.")
	}
}
