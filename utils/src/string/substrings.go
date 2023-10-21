package strings

func GenerateSubstrings(s string, minSize, maxSize, maxStartIndex int) []string {
	var subs []string
	for i := 0; i < maxStartIndex; i++ {
		for j := i + minSize; j <= len(s) && (j-i) <= maxSize; j++ {
			subs = append(subs, s[i:j])
		}
	}
	return subs
}
