package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func computeMD5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func main(){
	fmt.Println("===== Mini-CTF : Cat Leak =====")

	var inputHash string
	fmt.Print("Enter MD5 hash to reverse: ")
	fmt.Scanln(&inputHash)
	inputHash = strings.ToLower(strings.TrimSpace(inputHash))

	// Small sample wordlist — expand as needed
	wordlist := []string{
		"cat", "cats", "kitten", "meow", "meowmeow",
		"dog", "password", "admin", "hello", "love",
	}

	found := false
	for _, word := range wordlist {
		hash := computeMD5(word)
		if hash == inputHash {
			fmt.Printf("Horray!! Match found! MD5(%q) = %s\n", hash, word)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Opppsss...No match found in wordlist.")
	}
}

