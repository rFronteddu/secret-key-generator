package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter the desired key length: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Convert the user's input to an integer (key length)
	keyLength, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	// Generate a random secret key
	secretKey, err := generateKey(keyLength)
	if err != nil {
		fmt.Println("Error generating secret key:", err)
		return
	}

	fmt.Println("Generated Secret Key:", secretKey)
}

// generateKey generates a random secret key
func generateKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
