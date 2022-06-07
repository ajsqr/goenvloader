package goenvloader

import (
	"bufio"
	"io"
	"strings"
)

type tokens map[string]string

const (
	EQUALS = "="
	HASH   = "#"
	EMPTY  = ""
)

func createNewLineScanner(reader io.ReadCloser) *bufio.Scanner {
	scanner := bufio.NewScanner(reader)
	return scanner
}

func parseVars(v string) (string, string, error) {

	if !strings.Contains(v, EQUALS) {
		return "", "", &illegalVariableDeclared{token: v}
	}

	if strings.Contains(v, HASH) {
		commentIndex := strings.Index(v, "#")
		v = v[:commentIndex]
	}
	if strings.EqualFold(v, EMPTY) {
		return "", "", &emptyToken{}
	}

	envVar := strings.Split(v, EQUALS)
	return envVar[0], envVar[1], nil

}

func generateTokens(scanner *bufio.Scanner) (tokens, error) {
	var isMore bool = false
	generatedTokens := make(tokens)
	for {
		isMore = scanner.Scan()
		if !isMore {
			break
		}
		token := scanner.Text()
		key, value, err := parseVars(token)
		if err != nil {
			return generatedTokens, err
		}
		generatedTokens[key] = value
	}
	return generatedTokens, nil
}
