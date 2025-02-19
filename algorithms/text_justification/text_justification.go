// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	//we should pack as many words that we can in a line and pad spaces evenly between words
	lastIndex := 0
	var res []string
	for lastIndex < len(words) {
		var packedWords []string
		var buffer int
		packedWords, buffer, lastIndex = packWords(words, maxWidth, lastIndex)
		isLastLine := lastIndex >= len(words)
		paddedWords := padWords(packedWords, maxWidth, buffer, isLastLine)
		res = append(res, strings.Join(paddedWords, ""))
	}

	return res
}

func packWords(words []string, maxWidth int, nextIndex int) ([]string, int, int) {
	var res []string
	buffer := 0
	i := nextIndex

	for i < len(words) {
		word := words[i]
		if buffer+len(word)+len(res) > maxWidth {
			break
		}
		res = append(res, word)
		buffer += len(word)
		i++
	}

	return res, buffer, i
}

func padWords(words []string, maxWidth int, buffer int, isLastLine bool) []string {
	numSpaces := maxWidth - buffer

	if len(words) == 1 || isLastLine { // Left justify the last line
		return []string{strings.Join(words, " ") + strings.Repeat(" ", numSpaces-len(words)+1)}
	}

	// Distribute spaces evenly for justified lines
	spaceSlots := len(words) - 1
	spacePerSlot := numSpaces / spaceSlots
	extraSpaces := numSpaces % spaceSlots

	for i := 0; i < spaceSlots; i++ {
		words[i] += strings.Repeat(" ", spacePerSlot)
		if i < extraSpaces { // Add extra space to the left slots
			words[i] += " "
		}
	}

	return words
}

func main() {
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	res := fullJustify(words, 16)

	fmt.Println(strings.Join(res, ","))
}
