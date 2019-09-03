package tools

import (
	"strings"
	"testing"
)

func TestLoadRealFile(t *testing.T) {
	text, err := Load("../file.txt")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(text) == 0 {
		t.Errorf("Expected a text length greater than %v, got %v", 0, len(text))
	}
}

func TestLoadNotExistingFile(t *testing.T) {
	text, err := Load("../fileNil.txt")

	if err == nil {
		t.Errorf("Expected to get an error, got %v", text)
	}
}

func TestWithoutRealFile(t *testing.T) {
	s := "I am a nice text to test the loading of a file without an actual file ;-)"

	r := strings.NewReader(s)

	text, err := GetText(r)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !strings.EqualFold(s, text) {
		t.Errorf("Expected %v, got %v", s, text)
	}
}
