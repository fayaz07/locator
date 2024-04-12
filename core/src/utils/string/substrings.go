package strings

/**
 * GenerateSubstrings generates all possible substrings of a given string with a minimum length, maximum length and maximum start index.
 * @param s string
 * @param minLength int
 * @param maxLength int
 * @param maxStartIndex int
 * @return []string
 **/
func GenerateSubstrings(s string, minLength, maxLength, maxStartIndex int) []string {
	var subs []string
	for i := 0; i < maxStartIndex; i++ {
		for j := i + minLength; j <= len(s) && (j-i) <= maxLength; j++ {
			subs = append(subs, s[i:j])
		}
	}
	return subs
}
