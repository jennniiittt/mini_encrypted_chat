package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"week03_lab2/utils/crack"
)

func readTargetSha1FromStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target SHA1 hash (40 hex chars): ")
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	hash := strings.TrimSpace(strings.ToLower(line))
	if len(hash) != 40 {
		return "", fmt.Errorf("invalid length: expected 40 hex chars, got %d", len(hash))
	}
	if _, err := hex.DecodeString(hash); err != nil {
		return "", fmt.Errorf("invalid hex: %w", err)
	}
	return hash, nil
}

func main() {
	wordlistPath := "nord_vpn.txt"
	verbosePath := "verbose.txt"

	// read target SHA1 from user input
	targetSha1, err := readTargetSha1FromStdin()
	if err != nil {
		log.Fatalf("Error reading target SHA1: %v", err)
	}

	// create/overwrite verbose file
	verboseFile, err := os.Create(verbosePath)
	if err != nil {
		log.Fatalf("Cannot create verbose file: %v", err)
	}
	defer verboseFile.Close()

	plaintext, found, err := crack.CrackSHA1(targetSha1, wordlistPath, verboseFile)
	if err != nil {
		log.Fatalf("Crack error: %v", err)
	}

	// Output
	if found {
		fmt.Printf("FOUND: %s -> %q\n", targetSha1, plaintext)
	} else {
		fmt.Printf("NOT FOUND: hash %s not in %s\n", targetSha1, wordlistPath)
	}

	fmt.Printf("Verbose log saved to: %s\n", verbosePath)
}
