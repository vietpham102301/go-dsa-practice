package main

import "fmt"

//func replaceSpaces(s string, trueLength int) string {
//	var res strings.Builder
//	res.Grow(trueLength * 3)
//
//	for i := 0; i < trueLength; i++ {
//		if s[i] == ' ' {
//			res.WriteString("%20")
//		} else {
//			res.WriteByte(s[i])
//		}
//	}
//	return res.String()
//}

func urlify(str []byte, trueLength int) {
	spaceCount := 0
	for i := 0; i < trueLength; i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}

	index := trueLength + spaceCount*2 - 1
	for i := trueLength - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str[index] = '0'
			str[index-1] = '2'
			str[index-2] = '%'
			index -= 3
		} else {
			str[index] = str[i]
			index--
		}
	}
}

func main() {
	//fmt.Println(replaceSpaces("Mr John Smith   ", 13))
	str := []byte("Mr John Smith    ")
	urlify(str, 13)
	fmt.Println(string(str)) // Output: Mr%20John%20Smith

}
