package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter key or token? [key/token]: ")
	if !scanner.Scan() {
		// Handle the error if any
		fmt.Println("Error reading input:", scanner.Err())
		return
	}
	input := scanner.Text()

	if input != "key" && input != "token" {
		fmt.Println("Invalid input. Please enter 'key' or 'token'.")
		return
	}
	if input == "key" {
		key(scanner)
		return
	} else {
		token(scanner)
	}
}

func key(scanner *bufio.Scanner) {
	fmt.Print("Enter the desired key length: ")
	if !scanner.Scan() {
		// Handle the error if any
		fmt.Println("Error reading input:", scanner.Err())
		return
	}
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

func token(scanner *bufio.Scanner) {
	fmt.Print("Enter secret key: ")
	if !scanner.Scan() {
		// Handle the error if any
		fmt.Println("Error reading input:", scanner.Err())
		return
	}
	key := scanner.Text()

	fmt.Print("Enter user: ")

	user := scanner.Text()

	signedToken, err := generateJWT(user, []byte(key))
	if err != nil {
		fmt.Println("Error generating JWT:", err)
		return
	}

	fmt.Println("Generated Token:", signedToken)
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

// Claims represents the JWT claims structure.
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateJWT(username string, secretKey []byte) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// Token will expire in 1 year
			ExpiresAt: time.Now().Add(time.Hour * 9000).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
