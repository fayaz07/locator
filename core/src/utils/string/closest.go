package strings

func FindClosestStrings(arr []string, q string) []string {
	// use pattern matching to find closest strings
	lps := ComputeLPSArray(q)
	return searchSubstringInArray(arr, q, lps)
}

func KMPSearch(text, pattern string, lps []int) int {
	j, i := 0, 0

	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		}
		if j == len(pattern) {
			return i - j // Pattern found, returning the starting index in the text
		} else if i < len(text) && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1 // Pattern not found
}

func ComputeLPSArray(pattern string) []int {
	length := len(pattern)
	lps := make([]int, length)
	j, i := 0, 1

	for i < length {
		if pattern[i] == pattern[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func searchSubstringInArray(strings []string, pattern string, lps []int) []string {
	var results []string

	for _, str := range strings {
		startIndex := KMPSearch(str, pattern, lps)
		if startIndex != -1 {
			results = append(results, str)
		}
	}

	return results
}
