package utils

import (
	"regexp"
	"strings"
)

func ParseCamelCaseToSnakeCase(camelString string) string {
	re := regexp.MustCompile("(.)([A-Z][a-z]*)")
	snakeString := re.ReplaceAllString(camelString, "${1}_${2}")
	snakeString = strings.ToLower(snakeString)
	return snakeString
}

func FindTag(tag string, array []string) bool {
	for _, v := range array {
		if v == tag {
			return true
		}
	}
	return false
}
