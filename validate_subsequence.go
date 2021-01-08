package main

// IsValidSubsequence is a function for question validate subsequence
func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	if len(array) < len(sequence) {
		return false
	}

	mapResult := make([]bool, len(sequence))
	for _, arr := range array {
		for j, seq := range sequence {
			if arr == seq {
				if mapResult[j] {
					continue
				}
				mapResult[j] = true
				if j != 0 && mapResult[j-1] == false {
					return false
				}
				break
			}
		}
	}

	for i, result := range mapResult {
		if result == false {
			return false
		}
		if i == len(mapResult)-1 {
			return true
		}
	}
	return false
}
