package main

import "fmt"
import "errors"

// --- Subsystem Classes ---

// InventoryService handles stock checks and updates
type InventoryService struct{}

func (s *InventoryService) CheckStock(productID string, quantity int) error {
	fmt.Printf("InventoryService: Checking stock for %s, quantity %d...\n", productID, quantity)
	if quantity > 10 { // Simulate insufficient stock
		return errors.New("insufficient stock")
	}
	fmt.Println("InventoryService: Stock available.")
	return nil
}

func (s *InventoryService) DecreaseStock(productID string, quantity int) error {
	fmt.Printf("InventoryService: Decreasing stock for %s by %d...\n", productID, quantity)
	// Actual stock update logic
	fmt.Println("InventoryService: Stock decreased.")
	return nil
}

// PaymentService handles payment processing
type PaymentService struct{}

func (s *PaymentService) ProcessPayment(amount float64, cardNumber string) error {
	fmt.Printf("PaymentService: Processing payment of $%.2f for card %s...\n", amount, cardNumber)
	if amount > 1000.0 { // Simulate payment failure for large amounts
		return errors.New("payment declined")
	}
	// Actual payment gateway integration
	fmt.Println("PaymentService: Payment successful.")
	return nil
}

// NotificationService handles sending notifications
type NotificationService struct{}

func (s *NotificationService) NotifyCustomer(customerEmail, message string) error {
	fmt.Printf("NotificationService: Sending email to %s: '%s'\n", customerEmail, message)
	// Actual email sending logic
	fmt.Println("NotificationService: Email sent.")
	return nil
}

// --- The Facade ---

// OrderFacade provides a simplified interface to the order processing subsystem.
type OrderFacade struct {
	inventoryService    *InventoryService
	paymentService      *PaymentService
	notificationService *NotificationService
}

// NewOrderFacade is a constructor for the Facade
func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		inventoryService:    &InventoryService{},
		paymentService:      &PaymentService{},
		notificationService: &NotificationService{},
	}
}

// PlaceOrder is the simplified, high-level method exposed by the Facade.
// It orchestrates calls to the complex subsystem classes.
func (f *OrderFacade) PlaceOrder(
	productID string,
	quantity int,
	amount float64,
	cardNumber string,
	customerEmail string,
) error {
	fmt.Println("\n--- OrderFacade: Starting PlaceOrder ---")

	// Step 1: Check Inventory
	if err := f.inventoryService.CheckStock(productID, quantity); err != nil {
		fmt.Printf("OrderFacade: Failed to check stock: %v\n", err)
		return fmt.Errorf("order failed due to stock check: %w", err)
	}

	// Step 2: Process Payment
	if err := f.paymentService.ProcessPayment(amount, cardNumber); err != nil {
		fmt.Printf("OrderFacade: Failed to process payment: %v\n", err)
		return fmt.Errorf("order failed due to payment: %w", err)
	}

	// Step 3: Decrease Stock (only if payment was successful)
	if err := f.inventoryService.DecreaseStock(productID, quantity); err != nil {
		fmt.Printf("OrderFacade: Failed to decrease stock: %v\n", err)
		return fmt.Errorf("order failed after payment (stock update): %w", err)
	}

	// Step 4: Notify Customer
	notificationMessage := fmt.Sprintf("Your order for %d x %s has been placed successfully!", quantity, productID)
	if err := f.notificationService.NotifyCustomer(customerEmail, notificationMessage); err != nil {
		fmt.Printf("OrderFacade: Failed to notify customer: %v (order still placed)\n", err)
		// Depending on business rules, a notification failure might not rollback the whole order.
	}

	fmt.Println("--- OrderFacade: PlaceOrder Completed Successfully ---")
	return nil
}

// --- Client Code ---
func main() {
	orderFacade := NewOrderFacade()

	// Client places an order without knowing the complex steps involved
	fmt.Println("Attempting to place a small, normal order:")
	err := orderFacade.PlaceOrder(
		"Laptop-X1",
		1,
		899.99,
		"1111-2222-3333-4444",
		"alice@example.com",
	)
	if err != nil {
		fmt.Printf("Order process failed: %v\n", err)
	}

	fmt.Println("\nAttempting to place a large order (should fail payment):")
	err = orderFacade.PlaceOrder(
		"Server-Rack",
		1,
		1500.00, // Amount over 1000.00 should fail payment
		"5555-6666-7777-8888",
		"bob@example.com",
	)
	if err != nil {
		fmt.Printf("Order process failed: %v\n", err)
	}

	fmt.Println("\nAttempting to place an order with insufficient stock (should fail stock check):")
	err = orderFacade.PlaceOrder(
		"Widget-A",
		15, // Quantity over 10 should fail stock check
		50.00,
		"9999-8888-7777-6666",
		"charlie@example.com",
	)
	if err != nil {
		fmt.Printf("Order process failed: %v\n", err)
	}
}
