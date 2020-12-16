package digits

import (
	"bufio"
	"os"
)

func Input() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return text, nil
}
