package ctci

// 1.5 One Away: There are three types of edits that can be performed on
// strings: insert a character, remove a character, or replace a character.
// Given two strings, write a function to check if they are one edit (or zero
// edits) away.

// MaxOneAway breaks down what it means to replace, insert or delete a single
// character. It first checks the length of each string to eliminate a whole class
// of disallowed changes. Based on difference in length, it then validates
// the strings character by character to ensure that no more than one
// substitution or shift has occurred.
//
// This could be done in a single function, but the resulting code requires
// detection of the sorting string, maintaining two index variables, and
// multiple length checks. Ultimately, it feels less clear.
//
// Time complexity: O(n)
// Space complexity: O(1)
func MaxOneAway(s1, s2 string) bool {
	switch len(s2) - len(s1) {
	case 0:
		return oneAwayReplacement(s1, s2)
	case 1:
		return oneAwayInsertion(s1, s2)
	case -1:
		return oneAwayInsertion(s2, s1)
	default:
		return false
	}
}

func oneAwayReplacement(s1, s2 string) bool {
	var changed bool
	for i := range s1 {
		if s1[i] != s2[i] {
			if changed {
				return false
			}
			changed = true
		}
	}

	return true
}

func oneAwayInsertion(s1, s2 string) bool {
	var nChanges int
	for i := range s1 {
		if s1[i] != s2[i+nChanges] {
			if nChanges > 0 {
				return false
			}
			nChanges++
		}
	}

	return true
}
