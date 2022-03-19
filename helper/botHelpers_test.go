package helper

import (
	"testing"
	"time"
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

func TestCalculateTimeDifference(t *testing.T) {
	tests := []struct {
		name     string
		inputOne    time.Time
		inputTwo time.Time
		expected time.Duration
	}{
		{"hour", time.Now().UTC().Add(time.Hour), time.Now().UTC(), time.Hour},
		{"thirty mins", time.Now().UTC().Add(time.Minute*30), time.Now().UTC(), time.Minute*30},
		{"5 secs", time.Now().UTC().Add(time.Second*5), time.Now().UTC(), time.Second*5},
		{"passed by one hour", time.Now().UTC().Add(-time.Hour), time.Now().UTC(), -time.Hour},
		{"passed by ten mins", time.Now().UTC().Add(-time.Minute*10), time.Now().UTC(), -time.Minute*10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTime := CalculateTimeDifference(tt.inputOne, tt.inputTwo)
			if gotTime != tt.expected {
				t.Errorf("CalculateTimeDifference(tt.inputOne, tt.inputTwo) = %v, expected %v", gotTime, tt.expected)
			}
		})
	}
}
