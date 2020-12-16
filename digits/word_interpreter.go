package digits

import (
	"unicode"
)

func WordInterpreter(text string, selectedMap map[byte]byte) string {
	bytes := []byte(text)

	for i := range bytes {
		value, ok := selectedMap[byte(unicode.ToLower(rune(bytes[i])))]
		if ok {
			bytes[i] = value
		}
	}

	return string(bytes)
}
