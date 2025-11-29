package crack

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

// CrackResult contains the result of a crack attempt
type CrackResult struct {
	Plaintext string
	Found     bool
}

// CrackMD5 reads the wordlist from wordlistPath and tries each line.
// - targetHash must be hex-lowercase 32-char MD5 (e.g., "6a85df...").
// - verboseWriter will be used to write verbose logs (can be os.Stdout and/or a file).
// Returns CrackResult and error (if any IO error occurs).
func CrackMD5(targetHash string, wordlistPath string, verboseWriter io.Writer) (CrackResult, error) {
	// open wordlist
	f, err := os.Open(wordlistPath)
	if err != nil {
		return CrackResult{}, fmt.Errorf("failed to open wordlist: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNo := 0

	// normalize target hash
	targetHash = strings.TrimSpace(strings.ToLower(targetHash))

	for scanner.Scan() {
		lineNo++
		word := scanner.Text()

		// trim whitespace but keep internal spaces; skip empty lines
		word = strings.TrimSpace(word)
		if word == "" {
			fmt.Fprintf(verboseWriter, "[%d] skipping empty line\n", lineNo)
			continue
		}

		// compute md5
		sum := md5.Sum([]byte(word))
		sumHex := hex.EncodeToString(sum[:])

		// verbose log
		fmt.Fprintf(verboseWriter, "[%d] word=%q md5=%s\n", lineNo, word, sumHex)

		// compare
		if sumHex == targetHash {
			fmt.Fprintf(verboseWriter, "\nFOUND at line %d: %q -> %s\n", lineNo, word, sumHex)
			return CrackResult{Plaintext: word, Found: true}, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return CrackResult{}, fmt.Errorf("error reading wordlist: %w", err)
	}

	// not found
	fmt.Fprintf(verboseWriter, "Not found in wordlist: %s\n", wordlistPath)
	return CrackResult{Plaintext: "", Found: false}, nil
}
