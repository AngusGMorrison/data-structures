package ctci

import "sort"

// 1.1 Is Unique: Implement an algorithm to determine if a string has all
// unique characters. What if you cannot use additional data structures?

const charsetLen = 256

// AllUniqueCharsBruteForce compares every character in the string with every
// other character. Note that while this approach doesn't require extra space
// when dealing with ASCII character sets, if we were to use unicode, we'd
// be required to create a []rune from the input, giving us O(n) space
// complexity.
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func AllUniqueCharsBruteForce(input string) bool {
	if len(input) > charsetLen {
		return false
	}

	for i := range input {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}

	return true
}

// AllUniqueCharsArr uses an array of len(charset) booleans to record which
// characters from the input string have already been seen. This approach is
// most effective when the charset is reasonably sized. Where the size of the
// charset is likely to outstrip the length of most input strings (e.g. Unicode)
// a map may be more appropriate.
//
// Time complexity: O(n)
// Space complexity: O(1)
func AllUniqueCharsArr(input string) bool {
	if len(input) > charsetLen {
		return false
	}

	var seen [charsetLen]bool
	for _, r := range input {
		if seen[r] {
			return false
		}
		seen[r] = true
	}

	return true
}

// AllUniqueCharsBV uses a supporting bit vector and required 8 times less space
// than AllUniqueCharsArr. It assumes an extended ASCII character set of length
// 256 and represents each character using 1 bit of an array of 4 int64.
//
// Time complexity: O(n)
// Space complexity: O(1)
func AllUniqueCharsBV(input string) bool {
	if len(input) > charsetLen {
		return false
	}

	var bitVector [4]int64
	for _, r := range input {
		idx, bit := r/64, r%64
		if bitVector[idx]&(1<<bit) != 0 {
			return false
		}

		bitVector[idx] |= (1 << bit)
	}

	return true
}

// AllUniqueCharsSort assumes that copying strings is allowed. As a rune slice,
// the input can be sorted and a linear comparison of neighbouring runes can
// occur.
//
// Time complexity: O(nlogn)
// Space complexity: O(n). Go's strings are immutable, so a copy must be made.
func AllUniqueCharsSort(input string) bool {
	if len(input) > charsetLen {
		return false
	}

	runes := runeSlice(input)
	sort.Sort(runes)
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			return false
		}
	}

	return true
}

type runeSlice []rune

func (rs runeSlice) Len() int {
	return len(rs)
}

func (rs runeSlice) Less(i, j int) bool {
	return rs[i] < rs[j]
}

func (rs runeSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
