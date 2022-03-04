package api

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDecodeSingleIssue(t *testing.T) {
	tests := []struct {
		name  string
		input json.RawMessage
	}{
		{"Example 1", testerSetupJsonRaw(testGithubNewIssueResp)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := decodeSingleIssue(tt.input)
			log.Println(resp)
			if err != nil {
				t.Errorf("err = %v, expected nil", err)
			}
		})
	}
}

func TestRemovePullRequests(t *testing.T) {
	tests := []struct {
		name           string
		input          []Issue
		expectedLength int
	}{
		{"10 good then 10 bad", testerSetupIssueArray(), 10},
		{"every third issue is a pr", testerEvery3Pr(), 40},
		{"mixed up issues and prs", testerMixedUpIssuePr(), 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues, err := removePullRequests(tt.input)
			if err != nil {
				t.Errorf("err = %v, expected nil", err)
			}
			if len(issues) != tt.expectedLength {
				for z, i := range issues {
					for a, s := range issues {
						if z == a {
							continue
						}
						if i.ID == s.ID {
							t.Errorf("Duplicate found: %v", i.ID)
						}
					}
				}
				t.Errorf("len(issues) = %v, expected %v", len(issues), tt.expectedLength)
			}
		})
	}
}
