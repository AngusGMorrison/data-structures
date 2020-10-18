// Package rabinkarp reports all indexes of substring S within string B
// in O(b) time in the average case when the number of matches is small and
// the hash divisor is a prime number greater than len(B).
//
// In the worst case, when every index starts a valid substring, time complexity
// rises to O(sb).
package rabinkarp

import (
	"errors"
	"math"
)

const (
	lenCharset = 256

	// prime is required as a modular divisor to ensure hashes don't overflow
	// their integers while minimizing collisions.
	prime = 16777619
)

// AllIndexes returns a slice of all indexes in text at which substring pat is
// found.
func AllIndexes(pat, text string) ([]int, error) {
	patLen := len(pat)
	textLen := len(text)
	if patLen > textLen {
		return nil, ErrPatExceedsTextLen
	} else if patLen == 0 {
		return nil, ErrNoPattern
	}

	var (
		patHash   int
		textHash  int
		foundIdxs = []int{}

		// When sliding the hash window one character forwards along text, the
		// high-order character is removed from the hash by multiplying it by
		// highOrderFactor and subtracting it from the hash total. By
		// precomputing highOrderFactor, each shift will take a constant number
		// of operations instead of patLen-1 operations.
		highOrderFactor = int(math.Pow(float64(lenCharset), float64(patLen-1))) % prime
	)

	// Calculate the hash values of the pattern and the first window of text.
	for i := 0; i < patLen; i++ {
		patHash = (lenCharset*patHash + int(pat[i])) % prime
		textHash = (lenCharset*textHash + int(text[i])) % prime
	}

	// Slide the pattern over the text character by character
	for i := 0; i <= textLen-patLen; i++ {
		// Check if hash values match. If they do, compare each character (to
		// guard against collisions).
		if textHash == patHash {
			var j int
			for j = 0; j < patLen; j++ {
				if pat[j] != text[i+j] {
					break
				}
			}

			if j == patLen {
				foundIdxs = append(foundIdxs, i)
			}
		}

		// Calculate the hash value for the next window: removing the leading
		// character from the hash and add the next character.
		if i < textLen-patLen {
			textHash = (lenCharset*(textHash-int(text[i])*highOrderFactor) + int(text[i+patLen])) % prime
		}

		// textHash may be negative; make it positive.
		if textHash < 0 {
			textHash = textHash + prime
		}
	}
	return foundIdxs, nil
}

var ErrPatExceedsTextLen = errors.New("pattern length is greater than text length")

var ErrNoPattern = errors.New("pattern len must be > 0")
