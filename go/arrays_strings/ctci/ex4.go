package ctci

import (
	"unicode"
)

// 1.4 PalindromePermutation: Given a string, write a function to check if it
// is a permutation of a palindrome. The palindrome does not need to be limited
// to just dictionary words.
//
// Assumptions:
//	* Case insensitive.
//	* Whitespace should be ignored.
//	* Any other ASCII character is valid.

// IsPalindromePerm records the counts of each rune (case-insensitive,
// whitespace ignored) in an array. For character sets larger than ASCII, a map
// may be advised. It then checks that no more than one character appears an
// odd number of times.
//
// Time complexity: O(n)
// Space complexity: O(1)
func IsPalindromePerm(s string) bool {
	var counts [256]int
	for _, r := range s {
		if !unicode.IsSpace(r) {
			counts[unicode.ToLower(r)]++
		}
	}

	var nOddChars int
	for _, r := range s {
		if counts[r]%2 == 1 {
			if nOddChars > 0 {
				return false
			}

			nOddChars++
		}
	}

	return true
}

// IsPalindromePermAlt is a alternative implemenation with the same time and
// space complexity but only one loop iteration.
//
// Time complexity: O(n)
// Space complexity: O(1)
func IsPalindromePermAlt(s string) bool {
	var counts [256]int
	var nOddChars int
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}

		lc := unicode.ToLower(r)
		counts[lc]++
		if counts[lc]%2 == 1 {
			nOddChars++
		} else {
			nOddChars--
		}
	}

	return nOddChars <= 1
}

// IsPalindromePermVector uses a bit vector to record whether each rune
// (case insensitive) is present in an even or odd number. A bit is switched
// off if there is an even number of the given rune, and on if there is an odd
// number. Once each rune is mapped, there must be at most 1 bit switched on
// for s to be a palindrome permutation. I.e. either the vector == 0, or
// vector & (v-1) == 0.
//
// This approach is appropriate for limited character sets such as ASCII.
//
// Time complexity: O(n)
// Space complexity: O(1) (requires at least 32 times less space than the
// counting approach for ASCII extended)
func IsPalindromePermVector(s string) bool {
	vector := bitVectorFromString(s)
	return vector.hasMaxOneBitSet()
}

type asciiBitVector [4]int64 // large enough to represent 256 chars as bits

func bitVectorFromString(s string) asciiBitVector {
	var vector asciiBitVector
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}

		lc := unicode.ToLower(r)
		idx, bit := lc/64, lc%64
		vector[idx] ^= (1 << bit)
	}

	return vector
}

func (v asciiBitVector) hasMaxOneBitSet() bool {
	var bitsSet int
	for _, chunk := range v {
		if chunk != 0 {
			if hasOneBitSet(chunk) {
				bitsSet++
			} else {
				return false
			}
		}
	}

	return bitsSet <= 1
}

func hasOneBitSet(i int64) bool {
	return i&(i-1) == 0
}
