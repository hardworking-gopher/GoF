package main

import "fmt"
import "errors"

// --- 1. Subject (Interface) ---
// Defines the common interface for accessing a document.
type DocumentAccess interface {
	ReadContent() (string, error)
	// Other methods like WriteContent(), GetMetadata() etc.
}

// --- 2. Real Subject ---
// The actual document service that holds the sensitive content.
type RealDocument struct {
	content string
}

func NewRealDocument(content string) *RealDocument {
	return &RealDocument{content: content}
}

func (rd *RealDocument) ReadContent() (string, error) {
	fmt.Println("RealDocument: Providing actual content.")
	return rd.content, nil
}

// --- User Context (for demonstrating protection) ---
type User struct {
	Name  string
	Roles []string
}

func (u *User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

// --- 3. Proxy ---
// A Protection Proxy for DocumentAccess.
type DocumentProtectionProxy struct {
	realDocument *RealDocument // Reference to the Real Subject
	user         *User         // The user trying to access
}

// NewDocumentProtectionProxy is the constructor for our proxy.
func NewDocumentProtectionProxy(doc *RealDocument, user *User) *DocumentProtectionProxy {
	return &DocumentProtectionProxy{
		realDocument: doc,
		user:         user,
	}
}

// ReadContent implements the DocumentAccess interface, adding protection logic.
func (dp *DocumentProtectionProxy) ReadContent() (string, error) {
	fmt.Println("DocumentProtectionProxy: Intercepting ReadContent call.")

	// Protection logic: Only users with "admin" or "viewer" role can read.
	if dp.user == nil || (!dp.user.HasRole("admin") && !dp.user.HasRole("viewer")) {
		return "", errors.New("access denied: insufficient permissions to read document")
	}

	fmt.Println("DocumentProtectionProxy: Permission granted. Delegating to RealDocument.")
	// Delegate the call to the Real Subject
	return dp.realDocument.ReadContent()
}

// --- Client Code ---
func main() {
	// Create a real document
	secretDoc := NewRealDocument("This is highly confidential information.")

	// --- Scenario 1: Admin user accessing document ---
	adminUser := &User{Name: "Admin", Roles: []string{"admin"}}
	adminProxy := NewDocumentProtectionProxy(secretDoc, adminUser)
	fmt.Println("--- Admin User Attempt ---")
	content, err := adminProxy.ReadContent()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Admin read: '%s'\n", content)
	}

	fmt.Println("\n--- Viewer User Attempt ---")
	// --- Scenario 2: Viewer user accessing document ---
	viewerUser := &User{Name: "Viewer", Roles: []string{"viewer"}}
	viewerProxy := NewDocumentProtectionProxy(secretDoc, viewerUser)
	content, err = viewerProxy.ReadContent()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Viewer read: '%s'\n", content)
	}

	fmt.Println("\n--- Guest User Attempt (No role) ---")
	// --- Scenario 3: Guest user (no specific role) accessing document ---
	guestUser := &User{Name: "Guest", Roles: []string{"guest"}} // Or just {}
	guestProxy := NewDocumentProtectionProxy(secretDoc, guestUser)
	content, err = guestProxy.ReadContent()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Guest read: '%s'\n", content) // This should not happen
	}

	fmt.Println("\n--- Unauthenticated Attempt ---")
	// --- Scenario 4: Nil user (unauthenticated) accessing document ---
	unauthProxy := NewDocumentProtectionProxy(secretDoc, nil)
	content, err = unauthProxy.ReadContent()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Unauthenticated read: '%s'\n", content) // This should not happen
	}
}
