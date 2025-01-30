package main

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	for i, c := range haystack {
		for _, n := range needle {
			if c != n {
				break
			}
			remain := haystack[i:]
			isContained := checkConsecutiveChars(needle, remain)
			if isContained {
				return i
			}
		}
	}

	return -1
}

func checkConsecutiveChars(needle string, remainOfTheHaystack string) bool {
	for i, _ := range remainOfTheHaystack {
		if len(needle)-1 >= i && needle[i] != remainOfTheHaystack[i] {
			return false
		}
	}

	return true
}

func main() {
	needle := "sad"
	haystack := "sadbutsad"
	println(strStr(haystack, needle))
}
