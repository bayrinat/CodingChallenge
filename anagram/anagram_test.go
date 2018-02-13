package anagram

import "testing"

func TestAnagram(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{"one word anagrams", "tab", "bat", true},
		{"one word anagrams with spaces", "ta b", " b at", true},
		{"not anagrams", "write", "read", false},
		{"case sensitive anagrams", "William Shakespeare", "I am a weakish speller", true},
		{"not anagrams with one letter mistake", "Billiam Shakespeare", "I am a weakish speller", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anagram(tt.a, tt.b); got != tt.want {
				t.Errorf("Name: %s, Anagram() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
