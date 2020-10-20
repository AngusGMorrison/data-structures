package ctci

// 1.5 One Away: There are three types of edits that can be performed on
// strings: insert a character, remove a character, or replace a character.
// Given two strings, write a function to check if they are one edit (or zero
// edits) away.

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
