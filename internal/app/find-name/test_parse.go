package findname

import (
	"os"
	"strings"
)

func ReadTestFile() ([]string, error) {
	b, err := os.ReadFile("../../static/pan_tadeusz.txt") // just pass the file name
	if err != nil {
		return nil, err
	}
	tekst := string(b)
	tekst = strings.ReplaceAll(tekst, ".", "")
	tekst = strings.ReplaceAll(tekst, ",", "")
	tekst = strings.ReplaceAll(tekst, ":", "")
	tekst = strings.ReplaceAll(tekst, ";", "")
	tekst = strings.ReplaceAll(tekst, "!", "")
	tekst = strings.ReplaceAll(tekst, "?", "")
	tekst = strings.ReplaceAll(tekst, "\n", " ")
	tekst = strings.ReplaceAll(tekst, "[", " ")
	tekst = strings.ReplaceAll(tekst, "]", " ")
	tekst = strings.ReplaceAll(tekst, "(", " ")
	tekst = strings.ReplaceAll(tekst, ")", " ")
	tekst = strings.ReplaceAll(tekst, "{", " ")
	tekst = strings.ReplaceAll(tekst, "}", " ")
	tekst = strings.ReplaceAll(tekst, "\r", " ")
	words := strings.Split(tekst, " ")
	return words, nil
}
