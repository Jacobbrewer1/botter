package api

import "testing"

func testerStringPoint(input string) *string {
	return &input
}

func testerPointSession(date, time string) *Session {
	return &Session{
		Date: testerStringPoint(date),
		Time: testerStringPoint(time),
	}
}

func TestRace_GetSaturdayDate(t *testing.T) {
	tests := []struct {
		name     string
		input    Race
		expected string
	}{
		{"normal", Race{Session: testerPointSession("2022-03-20", "15:00:00Z")}, "2022-03-19T00:00:00Z"},
		{"missing trailing Z", Race{Session: testerPointSession("2022-03-20", "15:00:00")}, "2022-03-19T00:00:00Z"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate := tt.input.GetSaturdayDate()
			if gotDate != tt.expected {
				t.Errorf("Race.GetDateTime() = %v, expected %v", gotDate, tt.expected)
			}
		})
	}
}

func TestRace_GetFridayDate(t *testing.T) {
	tests := []struct {
		name     string
		input    Race
		expected string
	}{
		{"normal", Race{Session: testerPointSession("2022-03-20", "15:00:00Z")}, "2022-03-18T00:00:00Z"},
		{"missing trailing Z", Race{Session: testerPointSession("2022-03-20", "15:00:00")}, "2022-03-18T00:00:00Z"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate := tt.input.GetFridayDate()
			if gotDate != tt.expected {
				t.Errorf("Race.GetDateTime() = %v, expected %v", gotDate, tt.expected)
			}
		})
	}
}

func TestRace_GetDateTime(t *testing.T) {
	tests := []struct {
		name     string
		input    Race
		expected string
	}{
		{"normal", Race{Session: testerPointSession("2022-03-20", "15:00:00Z")}, "2022-03-20T15:00:00Z"},
		{"missing trailing Z", Race{Session: testerPointSession("2022-03-20", "15:00:00")}, "2022-03-20T15:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate := tt.input.GetSessionDateTime()
			if gotDate != tt.expected {
				t.Errorf("Race.GetDateTime() = %v, expected %v", gotDate, tt.expected)
			}
		})
	}
}
