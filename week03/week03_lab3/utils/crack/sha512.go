package crack

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func CrackSHA512(targetHash string, wordlistPath string, verboseWriter io.Writer) (string, bool, error) {
	f, err := os.Open(wordlistPath)
	if err != nil {
		return "", false, fmt.Errorf("failed to open wordlist: %w", err)
	}
	defer f.Close()

	targetHash = strings.TrimSpace(strings.ToLower(targetHash))
	scanner := bufio.NewScanner(f)
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		word := scanner.Text()
		// Trim trailing newline/CR and leading/trailing spaces (preserve internal spaces)
		word = strings.TrimSpace(word)
		if word == "" {
			fmt.Fprintf(verboseWriter, "[%d] skipping empty line\n", lineNo)
			continue
		}

		sum := sha512.Sum512([]byte(word))
		sumHex := hex.EncodeToString(sum[:])

		// verbose log
		fmt.Fprintf(verboseWriter, "[%d] word=%q sha512=%s\n", lineNo, word, sumHex)

		if sumHex == targetHash {
			fmt.Fprintf(verboseWriter, "FOUND at line %d: %q -> %s\n", lineNo, word, sumHex)
			return word, true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", false, fmt.Errorf("error reading wordlist: %w", err)
	}

	fmt.Fprintf(verboseWriter, "Not found in wordlist: %s\n", wordlistPath)
	return "", false, nil
}
