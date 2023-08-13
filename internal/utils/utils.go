package utils

import (
	"github.com/go-codegen/go-codegen/internal/colorPrint"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func RemoveCommonPath(path1 string, path2 string) string {
	// Разбиваем пути на компоненты
	components1 := strings.Split(path1, "/")
	components2 := strings.Split(path2, "/")

	// Находим индекс компонента, в котором пути различаются
	i := 0
	for i < len(components1) && i < len(components2) {
		if components1[i] != components2[i] {
			break
		}
		i++
	}

	// Если пути являются идентичными, убираем последний компонент
	if i == len(components1) && i == len(components2) {
		i--
	}

	// Объединяем оставшиеся компоненты в новый путь
	return "/" + filepath.Join(components1[:i]...)
}

func GetGlobalPath() (string, error) {
	globalPath, err := os.Getwd()
	if err != nil {
		colorPrint.PrintError(err)
		return "", err
	}
	globalPath = strings.Replace(globalPath, "\\", "/", -1)
	globalPath += "/"
	return globalPath, nil
}

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
