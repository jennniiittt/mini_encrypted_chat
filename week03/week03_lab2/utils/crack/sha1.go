package crack

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

// CrackSHA1 tries to find the plaintext for targetHash by scanning wordlistPath.
// - targetHash: lowered, trimmed 40-char hex SHA1 (e.g. "aa1c7d...").
// - verboseWriter: where to write per-line verbose logs (e.g. a file).
// Returns (plaintext, found, error).
func CrackSHA1(targetHash string, wordlistPath string, verboseWriter io.Writer) (string, bool, error) {
	f, err := os.Open(wordlistPath)
	if err != nil {
		return "", false, fmt.Errorf("failed to open wordlist: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNo := 0
	targetHash = strings.TrimSpace(strings.ToLower(targetHash))

	for scanner.Scan() {
		lineNo++
		word := scanner.Text()
		// remove leading/trailing newline/whitespace but preserve internal spaces
		word = strings.TrimSpace(word)
		if word == "" {
			fmt.Fprintf(verboseWriter, "[%d] skipping empty line\n", lineNo)
			continue
		}

		sum := sha1.Sum([]byte(word))
		sumHex := hex.EncodeToString(sum[:])

		// verbose log
		fmt.Fprintf(verboseWriter, "[%d] word=%q sha1=%s\n", lineNo, word, sumHex)

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
