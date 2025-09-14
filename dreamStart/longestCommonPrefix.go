package dreamstart

func longestCommonPrefix(strs []string) string {

	length := len(strs)

	maxLongPre := strs[0]

	for i := 1; i < length; i++ {
		maxLongPre = longestPrefix(maxLongPre, strs[i])
	}

	return maxLongPre
}

func longestPrefix(str1 string, str2 string) string {

	len1 := len(str1)
	len2 := len(str2)
	minLen := len1
	if minLen > len2 {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {

			return str1[:i]
		}
	}
	return str1[:minLen]

}
