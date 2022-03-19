package helper

import (
	"testing"
)

func TestFormatIdFromMention(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Normal", "<@!123456789>", "123456789"},
		{"No mention, just id", "123456789", "123456789"},
		{"Double mention", "<@!123456789> <@!987654321>", ""},
		{"Double mention", "<@!123456789><@!987654321>", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, _ := FormatIdFromMention(tt.input)
			if gotId != tt.expected {
				t.Errorf("formatIdFromMention(tt.input) = %v, expected %v", gotId, tt.expected)
			}
		})
	}
}

func TestFormatMessage(t *testing.T) {
	tests := []struct {
		name            string
		prefix          string
		input           string
		expectedCommand string
		expectedQuery   string
	}{
		{"Regular", ".", ".test This is a normal string", "test", "This is a normal string"},
		{"Multi space", ".", ".test     This  test contains    multiple         spaces", "test", "This test contains multiple spaces"},
		{"Punctuation", ".", ".test This text, contains some punctuation!", "test", "This text, contains some punctuation!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCommand, gotQuery := FormatMessage(tt.input, tt.prefix)
			if gotCommand != tt.expectedCommand {
				t.Errorf("formatIdFromMention(tt.input) : gotCommand = %v, expectedCommand %v", gotCommand, tt.expectedCommand)
			}
			if gotQuery != tt.expectedQuery {
				t.Errorf("formatIdFromMention(tt.input) : gotQuery = %v, expectedQuery %v", gotQuery, tt.expectedQuery)
			}
		})
	}
}

func TestRemoveMultiSpaces(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Normal", "This is a test string", "This is a test string"},
		{"Double space", "This  text  contains  only  double  spaces", "This text contains only double spaces"},
		{"Multi space", "This     text contains      multiple  spaces", "This text contains multiple spaces"},
		{"No spaces", "Thistextcontainsnospaces", "Thistextcontainsnospaces"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotString := RemoveMultiSpaces(tt.input)
			if gotString != tt.expected {
				t.Errorf("RemoveMultiSpaces(tt.input) = %v, expected %v", gotString, tt.expected)
			}
		})
	}
}

func TestRemoveNonAlphaChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Normal", "This is a test string", "This is a test string"},
		{"Hashtag separator", "#This#text#is#separated#by#hashtags#", "This text is separated by hashtags"},
		{"Multi punctuation", "This!\"£$%^&*(({][]'[]';'/..//[`||text contains      multiple  punctuation and spaces", "This text contains multiple punctuation and spaces"},
		{"No spaces", "Thistextcontainsnospaces", "Thistextcontainsnospaces"},
		{"Solo word !", "This !", "This"},
		{"Just !", "!", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotString := RemoveNonAlphaChars(tt.input)
			if gotString != tt.expected {
				t.Errorf("RemoveMultiSpaces(tt.input) = %v, expected %v", gotString, tt.expected)
			}
		})
	}
}
