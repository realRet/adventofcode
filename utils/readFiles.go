package utils

import (
	"bufio"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"
)

func ReadFileAsString(filepath string) string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func ReadFileAsScanner(filepath string) *bufio.Scanner {
	scanner := bufio.NewScanner(strings.NewReader(ReadFileAsString(filepath)))

	return scanner
}

func SplitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
