// stringutils.go
package stringutils

import "strings"

// ReverseWords возвращает строку с обратным порядком слов
func ReverseWords(s string) string {
	// Разбиваем строку на слова
	words := strings.Fields(s)

	// Меняем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем строку обратно
	return strings.Join(words, " ")
}
