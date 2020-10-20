package ctci

// 1.3 URLify: Write a method to replace all spaces in a string with '%20'. You
// may assume that the string has sufficient space at the end to hold the
// additional characters, and that you are given the "true" length of the
// string. (In Go, use []rune so that you can perform this operation in place).

// URLify assumes:
//  * "space" means a literal space character, not unicode.IsSpace.
//  * Leading and trailing spaces should also be converted.
//
// Time complexity: O(n)
// Space complexity: O(1)
func URLify(input []rune, length int) []rune {
	// Count spaces
	var spaces int
	for _, r := range input {
		if r == ' ' {
			spaces++
		}
	}

	// Iterate backwards, expanding spaces
	moveTo := (length + spaces*2) - 1
	for i := length - 1; i >= 0; i-- {
		r := input[i]
		if r == ' ' {
			input[moveTo] = '0'
			input[moveTo-1] = '2'
			input[moveTo-2] = '%'
			moveTo -= 3
		} else {
			input[moveTo] = r
			moveTo--
		}
	}

	return input
}
