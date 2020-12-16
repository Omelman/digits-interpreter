package digits

import (
	"strings"
	"unicode"
)

func TranslateText(in string, selectedMap map[byte]byte) string {
	var readyWords []string
	var separators []string
	text := in

	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '.' || c == ',' || c == ':' || c == '-' ||
			c == '!' || c == '?' || c == ';'
	}

	Words := strings.FieldsFunc(text, f)

	for i := 0; i < len(Words)-1; i++ {
		leftPos := strings.Index(text, Words[i]) + len(Words[i])
		rightPos := strings.Index(text, Words[i+1])
		separators = append(separators, text[leftPos:rightPos])
	}
	//string after last word
	lastWordPos := strings.Index(text, Words[len(Words)-1]) + len(Words[len(Words)-1])
	separators = append(separators, text[lastWordPos:])

	for i, word := range Words {
		readyWords = append(readyWords, Interpreter(word, selectedMap)+separators[i])
	}

	return strings.Join(readyWords, "")
}
