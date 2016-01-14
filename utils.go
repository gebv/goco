package main

import (
	"bytes"
	// "regexp"
	// "text/template"
	"unicode"
)

// toUpper transform "str str" to "Str<separator>Str"
func toUpper(s, replace string) string {
	res := bytes.NewBufferString("")
	hasSep := true // Первый символ с большой буквы

	for index, runeValue := range s {
		if !unicode.IsLetter(runeValue) {
			hasSep = true
			continue
		}

		if hasSep {
			if index > 0 {
				res.WriteString(replace)
			}

			runeValue = unicode.ToUpper(runeValue)
		}

		res.WriteRune(runeValue)

		hasSep = false
	}

	return res.String()
}

// toLower transform "str str" to "str<separator>str"
func toLower(s, replace string) string {
	res := bytes.NewBufferString("")
	hasSep := false

	for index, runeValue := range s {
		if !unicode.IsLetter(runeValue) {
			hasSep = true
			continue
		}

		if index > 0 && hasSep {
			res.WriteString(replace)
		}

		if unicode.IsUpper(runeValue) {

			res.WriteRune(unicode.ToLower(runeValue))
			continue
		}

		// IsLower
		res.WriteRune(runeValue)

		hasSep = false
	}

	return res.String()
}

func firstLower(str string) string {
	return string(unicode.ToLower([]rune(str)[0]))
}
