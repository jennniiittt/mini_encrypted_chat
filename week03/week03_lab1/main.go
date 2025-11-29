package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"week03_lab1/utils/crack"
)

func readTargetHashFromStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target MD5 hash (32 hex chars): ")
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	hash := strings.TrimSpace(strings.ToLower(line))
	// basic validation: must be 32 hex chars
	if len(hash) != 32 {
		return "", fmt.Errorf("invalid length: expected 32 hex chars, got %d", len(hash))
	}
	_, err = hex.DecodeString(hash)
	if err != nil {
		return "", fmt.Errorf("invalid hex: %w", err)
	}
	return hash, nil
}

func main() {
	wordlistPath := "nord_vpn.txt"
	verbosePath := "verbose.txt"

	// read target MD5 from user input
	targetMd5, err := readTargetHashFromStdin()
	if err != nil {
		log.Fatalf("Error reading target MD5: %v", err)
	}

	// open verbose output file (create/overwrite)
	verboseFile, err := os.Create(verbosePath)
	if err != nil {
		log.Fatalf("Cannot create verbose file: %v", err)
	}
	defer verboseFile.Close()

	result, err := crack.CrackMD5(targetMd5, wordlistPath, verboseFile)
	if err != nil {
		log.Fatalf("Crack error: %v", err)
	}

	// Output showing only the matching plaintext if found,
	if result.Found {
		fmt.Printf("FOUND: %s -> %q\n", targetMd5, result.Plaintext)
	} else {
		fmt.Printf("NOT FOUND: hash %s not in %s\n", targetMd5, wordlistPath)
	}

	fmt.Printf("Verbose log saved to: %s\n", verbosePath)
}
