package utils

import (
	"unicode"
)

func ParseCamelCaseToSnakeCase(s string) string {
	result := make([]rune, 0, len(s))

	for i, r := range s {
		isUpperCase := unicode.IsUpper(r)
		isFirstCharacter := i == 0
		isLastCharacter := i == len(s)-1

		if isUpperCase {
			if isFirstCharacter {
				result = append(result, unicode.ToLower(r))
				continue
			}

			if !isLastCharacter {
				nextCharacterIsLower := unicode.IsLower(rune(s[i+1]))
				prevCharacterIsLower := unicode.IsLower(rune(s[i-1]))

				if nextCharacterIsLower || prevCharacterIsLower {
					result = append(result, '_')
				}
			}
		}

		result = append(result, unicode.ToLower(r))
	}

	return string(result)
}

func FindTag(tag string, array []string) bool {
	for _, v := range array {
		if v == tag {
			return true
		}
	}
	return false
}
