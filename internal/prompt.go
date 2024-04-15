package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
