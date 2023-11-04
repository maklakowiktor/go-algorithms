package problems

func IsSubsequence(s string, t string) bool {
	lenS, lenT := len(s), len(t)
	if lenT < lenS {
		return false
	}
	l, r := 0, 0
	for l < lenS && r < lenT {
		if s[l] == t[r] {
			l += 1
		}
		r += 1
	}
	return l == lenS
}
