package ctci

import (
	"strconv"
	"strings"
)

// 1.6 String Compression: Implement a method to perform basic string
// compression using the counts of repeated characters. For example, the string
// aabcccccaaa would become a2b1c5a3. If the compressed string would  not become
// smaller than the original string, your method should return the original
// string. You can assume that the string only has uppercase and lowercase
// letters a-z.

// CompressOnePass iterates once over the the string, counting contiguous
// repeated characters. When the next character doesn't match the current
// character, or the string ends, the current character and the count are
// written to a Builder and the count is reset.
//
// Time complexity: O(n + mlogp), where n is the length of the string, m
// is the number of subsequences and p is the length of the longest subsequence.
// logp is the number of digits in p, which must first be converted to a string
// with strconv.Itoa, then copied to the Builder.
// Space complexity: O(mlogp)
func CompressOnePass(s string) string {
	var sb strings.Builder
	var count int
	for i := range s {
		count++
		if i+1 >= len(s) || s[i] != s[i+1] {
			sb.WriteByte(s[i])
			sb.WriteString(strconv.Itoa(count))
			count = 0
		}
	}

	if len(s) <= sb.Len() {
		return s
	}
	return sb.String()
}

// CompressTwoPass first identifies whether a compressed string would be
// shorter than the original, and immediately returns the original if not. It
// does this at the expense of an extra pass over the string and code
// replication.
//
// Time complexity: O(n + mlogp)
// Space complexity: O(mlogp)
func CompressTwoPass(s string) string {
	compressedLen := calcCompressedLen(s)
	if len(s) <= compressedLen {
		return s
	}

	var reps int
	var sb strings.Builder
	sb.Grow(compressedLen)
	for i := range s {
		reps++
		if i+1 >= len(s) || s[i] != s[i+1] {
			sb.WriteByte(s[i])
			sb.WriteString(strconv.Itoa(reps))
			reps = 0
		}
	}

	return sb.String()
}

func calcCompressedLen(s string) int {
	var reps, compressedLen int
	for i := range s {
		reps++
		if i+1 >= len(s) || s[i] != s[i+1] {
			compressedLen += 1 + len(strconv.Itoa(reps))
			reps = 0
		}
	}

	return compressedLen
}
