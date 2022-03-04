package bot

import (
	"log"
	"testing"
)

func TestContainsServerInvite(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		expected bool
	}{
		{"Good single link", "https://discord.com/Test", false},
		{"Bad single link", "https://discord.gg/Test", true},
		{"Good complex link", "Here is a link https://discord.com/Test", false},
		{"Bad complex link", "Here is a link https://discord.gg/Test", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool := containsServerInvite(tt.message)
			if gotBool != tt.expected {
				t.Errorf("containsServerInvite(tt.message) = %v, expected %v", gotBool, tt.expected)
			}
		})
	}
}

func TestBannedWordFilter(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		expected bool
	}{
		{"Single word clear", "Test", false},
		{"Single word bad", "shit", true},
		{"Single word bad CAPS", "SHIT", true},
		{"Multi word clear", "This is a test message", false},
		{"Multi word bad", "This is a shit test message", true},
		{"Multi word bad CAPS", "This is a Shit test message", true},
		{"Complex clear", "This message is a grass message and should be allowed", false},
		{"Complex bad", "This message is an ass grass message and should not be allowed through", true},
		{"Good word with suffix \"ing\"", "gooding", false},
		{"Good word with suffix \"er\"", "gooder", false},
		{"Good word with prefix \"-\"", "-good", false},
		{"Bad word with suffix \"ing\"", "fucking", true},
		{"Bad word with suffix \"er\"", "fucker", true},
		{"Bad word with prefix \"-\"", "-shit", true},
	}

	err := TestingSetupBadWords()
	if err != nil {
		log.Panic(err)
		return
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool := bannedWordsFilter(tt.message)
			if gotBool != tt.expected {
				t.Errorf("bannedWordsFilter(tt.message) = %v, expected %v", gotBool, tt.expected)
			}
		})
	}
}
