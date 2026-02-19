// main.go
package main

import (
	"fmt"
	"newProject/pkg/stringutils"
)

func main() {
	examples := []string{
		"Hello, world!",
		"The quick brown fox",
		"Привет, как дела?",
		"Go is awesome",
	}

	for _, example := range examples {
		fmt.Printf("Исходная строка: %q\n", example)
		fmt.Printf("Результат: %q\n\n", stringutils.ReverseWords(example))
	}
}
