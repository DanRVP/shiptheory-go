package shiptheory

import (
	"log"
	"bytes"
)

// Convert a CamelCased (or is it PascalCase, idk) string to a snake_cased string.
// For example: CamelCase = camel_case or CamelCase123Camel = camel_case_123_camel
func camelToSnake(camel string) string {
	var buf bytes.Buffer
	runes := []rune(camel)
    for i, char := range runes {
        if runeIsUpper(char) {
            // just convert [A-Z] to _[a-z]
            if buf.Len() > 0 {
                buf.WriteRune('_')
            }

            buf.WriteRune(char - 'A' + 'a')
        } else if !runeIsNumber(char) && checkNextRune(i, runes, runeIsNumber) {
            buf.WriteRune(char)
			buf.WriteRune('_')
        } else {
			buf.WriteRune(char)
		}
    }

    return buf.String()
}

func runeIsNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func runeIsUpper(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func runeIsLower(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func checkNextRune(i int, runes []rune, f func(rune) bool) bool {
	if len(runes) <= i + 1 {
		return false
	}

	return f(runes[i + 1])
}

func checkError(err error) {
	if err != nil {
        log.Fatal(err)
    }
}
