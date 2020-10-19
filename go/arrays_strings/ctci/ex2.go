package ctci

import "sort"

// 1.2 Check Permutation: Given two strings, write a method to decide if one is
// a permutation of the other.

// IsPermutationArr assumes an extended ASCII character set and treats upper-
// and lowercase letters as distinct characters. A hash table may be more
// appropriate where there length of the character set outstrips the length
// of any likely input.
//
// Time complexity: O(n)
// Space complexity: O(1) for an array, O(n) for a hash table
func IsPermutationArr(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var charCounts [256]rune
	for _, r := range s1 {
		charCounts[r]++
	}

	for _, r := range s2 {
		if charCounts[r] == 0 {
			return false
		}

		charCounts[r]--
	}

	return true
}

// IsPermutationSort first creates rune slices from the input strings, then
// sorts them and compares each pair of runes in turn. It does not assume a
// character set, but does assume that upper and lowercase runes should be
// treated as different characters.
//
// Time complexity: O(nlogn)
// Space complexity: O(n) – strings can't be sorted in-place.
func IsPermutationSort(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1 := runeSlice(s1)
	sort.Sort(r1)
	r2 := runeSlice(s2)
	sort.Sort(r2)

	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
