package main

import (
	"bufio"
	"fmt"
	"os"
)

// xorEncrypt encrypts or decrypts text using a single-byte XOR key
func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key // XOR operation
	}
	return string(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a message: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1] // remove newline

	fmt.Print("Enter a key (single character): ")
	key, _ := reader.ReadByte()

	encrypted := xorEncrypt(text, key)
	fmt.Println("Encrypted text:", encrypted)

	decrypted := xorEncrypt(encrypted, key)
	fmt.Println("Decrypted text:", decrypted)
}
