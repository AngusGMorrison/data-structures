package ctci

import "strings"

// 1.9 String Rotation: Assume you have a method isSubstring which checks if
// one word is a substring of another. Given two strings, s1 and s2, write
// code to check if s2 is a rotation of s1 using only one call to isSubstring
// (e.g. "waterbottle" is a rotation of "erbottlewat").

// IsRotation concatenates s2. If s2 is a rotation of s1, s1 will be contained
// within the concatenated string by definition. This is true for rotations in
// either direction.
//
// Time complexity: O(n) (i.e. substring search O(a+b) on strings of equal
// length)
// Space complexity: O(n)
func IsRotation(s1, s2 string) bool {
	if len(s1) == len(s2) {
		return strings.Contains(s2+s2, s1)
	}

	return false
}
