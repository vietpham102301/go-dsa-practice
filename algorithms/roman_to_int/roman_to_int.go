package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	res := 0
	for i := 0; i < len(s); i++ {
		temp := ""
		number, ok := romanMap[string(s[i])]
		if ok && i != len(s)-1 {
			temp = string(s[i]) + string(s[i+1])
			secondNumber, ok := romanMap[temp]
			if ok {
				i = i + 1
				number = secondNumber
			}
		}
		res += number
	}
	return res
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
