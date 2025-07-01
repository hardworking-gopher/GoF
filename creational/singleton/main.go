package main

import (
	"fmt"
	"sync"
)

// The databaseConnection struct is our singleton object.
// Its fields are not exported to prevent external modification.
type databaseConnection struct {
	connectionString string
}

var (
	instance *databaseConnection
	once     sync.Once
)

// GetDBInstance is the global access point for the singleton instance.
// It uses sync.Once to ensure the instance is created only once.
func GetDBInstance() *databaseConnection {
	once.Do(func() {
		// This function will only be executed the very first time GetDBInstance is called.
		fmt.Println("Creating database connection instance now.")
		instance = &databaseConnection{
			connectionString: "server=my-db;user=root;password=secret;",
		}
	})
	return instance
}

// A method on our singleton to demonstrate its usage.
func (db *databaseConnection) GetConnectionString() string {
	return db.connectionString
}

func main() {
	// We'll use a WaitGroup to simulate concurrent access.
	var wg sync.WaitGroup

	// Start 100 goroutines that all try to get the instance.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conn := GetDBInstance()
			fmt.Printf("Goroutine %d got connection string: %s\n", i, conn.GetConnectionString())
		}(i)
	}

	wg.Wait()

	// All goroutines will receive the same instance, and the creation message
	// will only be printed once.
	fmt.Println("\nFinished. All goroutines used the same singleton instance.")
}
