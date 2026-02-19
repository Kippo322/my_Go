// stringutils_test.go
package stringutils

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple", "Hello world", "world Hello"},
		{"with punctuation", "Hello, world!", "world! Hello,"},
		{"russian text", "Привет, как дела?", "дела? как Привет,"},
		{"single word", "Привет", "Привет"},
		{"empty string", "", ""},
		{"multiple spaces", "  Hello   world  ", "world Hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.input); got != tt.want {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
