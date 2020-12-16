package digits

import (
	"strings"
)

func Interpreter(text string, selectedMap map[byte]byte) string {
	bytes := []byte(strings.ToLower(text))

	for i := range bytes {
		value, ok := selectedMap[bytes[i]]
		if ok {
			bytes[i] = value
		}
	}

	return string(bytes)
}
