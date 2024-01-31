package main

import (
 "encoding/json"
 "fmt"
 "math/rand"
 "net/http"
 "sync"
 "time"
)

// NumberGenerator struct holds the state of the number generation.
type NumberGenerator struct {
 mu         sync.Mutex
 lastNumber int
 numbers    []int
}

// generateNumber generates a distinct number from 1-99.
func (ng *NumberGenerator) generateNumber() int {
 ng.mu.Lock()
 defer ng.mu.Unlock()

 // Generate a distinct number from 1-99
 newNumber := rand.Intn(90) + 1
 for newNumber == ng.lastNumber {
  newNumber = rand.Intn(90) + 1
 }

 ng.lastNumber = newNumber
 ng.numbers = append(ng.numbers, newNumber)

 return newNumber
}

// listNumbers returns the list of all generated numbers.
func (ng *NumberGenerator) listNumbers() []int {
 ng.mu.Lock()
 defer ng.mu.Unlock()
 return ng.numbers
}

// resetNumbers resets the list of generated numbers.
func (ng *NumberGenerator) resetNumbers() {
 ng.mu.Lock()
 defer ng.mu.Unlock()
 ng.numbers = nil
}

// handleGenerateNumber handles the "/generate" endpoint.
func handleGenerateNumber(w http.ResponseWriter, r *http.Request, ng *NumberGenerator) {
 number := ng.generateNumber()
 response := map[string]int{"number": number}

 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(response)
}

// handleListNumbers handles the "/list" endpoint.
func handleListNumbers(w http.ResponseWriter, r *http.Request, ng *NumberGenerator) {
 numbers := ng.listNumbers()
 response := map[string][]int{"numbers": numbers}

 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(response)
}

// handleResetNumbers handles the "/reset" endpoint.
func handleResetNumbers(w http.ResponseWriter, r *http.Request, ng *NumberGenerator) {
 ng.resetNumbers()
 w.WriteHeader(http.StatusOK)
}

func main() {
 // Seed the random number generator
 rand.Seed(time.Now().UnixNano())

 // Create a new NumberGenerator instance
 numGen := &NumberGenerator{}

 // Define HTTP handlers
 http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
  handleGenerateNumber(w, r, numGen)
 })

 http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
  handleListNumbers(w, r, numGen)
 })

 http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
  handleResetNumbers(w, r, numGen)
 })

 // Start the HTTP server
 port := 8080
 fmt.Printf("Server is running on :%d\n", port)
 err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
 if err != nil {
  fmt.Println("Error starting server:", err)
 }
}
