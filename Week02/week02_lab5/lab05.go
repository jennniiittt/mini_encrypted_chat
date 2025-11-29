package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// convertToBinary converts a string to its binary representation
func convertToBinary(s string) string {
	binary := ""
	for i := 0; i < len(s); i++ {
		binary += fmt.Sprintf("%08b ", s[i]) // %08b = 8-bit binary with leading zeros
	}
	return binary
}

func main() {
	var input string
	fmt.Print("Enter text: ")
	fmt.Scanln(&input)

	// Binary representation
	binary := convertToBinary(input)

	// Hexadecimal representation
	hexadecimal := hex.EncodeToString([]byte(input))

	// Base64 representation
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(input))

	fmt.Println("\n--- Encoded Representations ---")
	fmt.Println("Original Text : ", input)
	fmt.Println("Binary        : ", binary)
	fmt.Println("Hexadecimal   : ", hexadecimal)
	fmt.Println("Base64        : ", base64Encoded)
}
