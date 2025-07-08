package main

import "fmt"

// --- 1. Strategy (Interface) ---
// Defines the common interface for all payment methods.
type PaymentStrategy interface {
	Pay(amount float64) error
}

// --- 2. Concrete Strategy(s) ---

// CreditCardPayment is a concrete strategy for credit card payments.
type CreditCardPayment struct {
	cardNumber string
	cvv        string
}

func NewCreditCardPayment(cardNumber, cvv string) *CreditCardPayment {
	return &CreditCardPayment{cardNumber: cardNumber, cvv: cvv}
}

func (c *CreditCardPayment) Pay(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f using card %s...\n", amount, c.cardNumber)
	// Simulate actual credit card processing logic
	if amount > 1000.0 { // Simulate a large transaction failure
		return fmt.Errorf("credit card payment declined for amount %.2f", amount)
	}
	fmt.Println("Credit card payment successful!")
	return nil
}

// PayPalPayment is a concrete strategy for PayPal payments.
type PayPalPayment struct {
	email string
}

func NewPayPalPayment(email string) *PayPalPayment {
	return &PayPalPayment{email: email}
}

func (p *PayPalPayment) Pay(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f for account %s...\n", amount, p.email)
	// Simulate actual PayPal API interaction
	if amount > 500.0 { // Simulate a PayPal limit
		return fmt.Errorf("PayPal payment limit exceeded for amount %.2f", amount)
	}
	fmt.Println("PayPal payment successful!")
	return nil
}

// CryptocurrencyPayment is a concrete strategy for cryptocurrency payments.
type CryptocurrencyPayment struct {
	walletAddress string
	cryptoType    string
}

func NewCryptocurrencyPayment(walletAddress, cryptoType string) *CryptocurrencyPayment {
	return &CryptocurrencyPayment{walletAddress: walletAddress, cryptoType: cryptoType}
}

func (c *CryptocurrencyPayment) Pay(amount float64) error {
	fmt.Printf("Processing %.2f in %s to wallet %s...\n", amount, c.cryptoType, c.walletAddress)
	// Simulate blockchain transaction
	if amount < 0.01 { // Simulate minimum transaction amount
		return fmt.Errorf("%s payment minimum not met for amount %.2f", c.cryptoType, amount)
	}
	fmt.Println("Cryptocurrency payment successful!")
	return nil
}

// --- 3. Context ---
// The ShoppingCart uses a PaymentStrategy to process payments.
type ShoppingCart struct {
	amount float64
	// Context holds a reference to the strategy interface.
	paymentStrategy PaymentStrategy
}

func NewShoppingCart(amount float64) *ShoppingCart {
	return &ShoppingCart{amount: amount}
}

// SetPaymentStrategy allows the client to choose the strategy at runtime.
func (sc *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	sc.paymentStrategy = strategy
	fmt.Printf("ShoppingCart: Payment strategy set.\n")
}

// Checkout delegates the payment task to the current strategy.
func (sc *ShoppingCart) Checkout() error {
	if sc.paymentStrategy == nil {
		return fmt.Errorf("no payment strategy set")
	}
	fmt.Printf("ShoppingCart: Initiating checkout for total $%.2f...\n", sc.amount)
	// The Context delegates to the strategy
	return sc.paymentStrategy.Pay(sc.amount)
}

// --- Client Code ---
func main() {
	// Create a shopping cart with a total amount
	cart1 := NewShoppingCart(120.50)

	// --- Scenario 1: Pay with Credit Card ---
	fmt.Println("\n--- Shopping Cart 1: Paying with Credit Card ---")
	creditCard := NewCreditCardPayment("1234-5678-9012-3456", "123")
	cart1.SetPaymentStrategy(creditCard)
	err := cart1.Checkout()
	if err != nil {
		fmt.Printf("Checkout failed: %v\n", err)
	}

	// --- Scenario 2: Pay with PayPal ---
	cart2 := NewShoppingCart(350.00)
	fmt.Println("\n--- Shopping Cart 2: Paying with PayPal ---")
	payPal := NewPayPalPayment("user@example.com")
	cart2.SetPaymentStrategy(payPal)
	err = cart2.Checkout()
	if err != nil {
		fmt.Printf("Checkout failed: %v\n", err)
	}

	// --- Scenario 3: Pay with Crypto (high amount, might fail strategy specific check) ---
	cart3 := NewShoppingCart(0.005) // Amount below minimum for crypto
	fmt.Println("\n--- Shopping Cart 3: Paying with Crypto (low amount) ---")
	crypto := NewCryptocurrencyPayment("0xAbc123...", "ETH")
	cart3.SetPaymentStrategy(crypto)
	err = cart3.Checkout()
	if err != nil {
		fmt.Printf("Checkout failed: %v\n", err)
	}

	// --- Scenario 4: Dynamic Strategy Change (e.g., credit card fails, try PayPal) ---
	cart4 := NewShoppingCart(1500.00)
	fmt.Println("\n--- Shopping Cart 4: Attempting Credit Card, then PayPal ---")
	creditCardFail := NewCreditCardPayment("1111-2222-3333-4444", "456")
	cart4.SetPaymentStrategy(creditCardFail)
	err = cart4.Checkout()
	if err != nil {
		fmt.Printf("Initial checkout failed: %v. Trying PayPal...\n", err)
		payPalFallback := NewPayPalPayment("backup@example.com")
		cart4.SetPaymentStrategy(payPalFallback)
		err = cart4.Checkout()
		if err != nil {
			fmt.Printf("Fallback checkout also failed: %v\n", err)
		}
	}
}
