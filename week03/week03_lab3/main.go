package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"week03_lab3/utils/crack"
)

func readTargetSha512FromStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target SHA-512 hash (128 hex chars): ")
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	hash := strings.TrimSpace(strings.ToLower(line))
	if len(hash) != 128 {
		return "", fmt.Errorf("invalid length: expected 128 hex chars, got %d", len(hash))
	}
	if _, err := hex.DecodeString(hash); err != nil {
		return "", fmt.Errorf("invalid hex: %w", err)
	}
	return hash, nil
}

func main() {
	wordlistPath := "nord_vpn.txt"
	verbosePath := "verbose.txt"

	target, err := readTargetSha512FromStdin()
	if err != nil {
		log.Fatalf("Error reading target SHA-512: %v", err)
	}

	// create/overwrite verbose file
	verboseFile, err := os.Create(verbosePath)
	if err != nil {
		log.Fatalf("Cannot create verbose file: %v", err)
	}
	defer verboseFile.Close()

	plaintext, found, err := crack.CrackSHA512(target, wordlistPath, verboseFile)
	if err != nil {
		log.Fatalf("Crack error: %v", err)
	}

	// Output
	if found {
		fmt.Printf("FOUND: %s -> %q\n", target, plaintext)
	} else {
		fmt.Printf("NOT FOUND: hash %s not in %s\n", target, wordlistPath)
	}
	fmt.Printf("Verbose log saved to: %s\n", verbosePath)
}
