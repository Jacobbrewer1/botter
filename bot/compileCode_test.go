package bot

import "testing"

func TestContainsCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"true", "```golang\npackage main;\n\nimport \"fmt\";\n\nfunc main() {\nfmt.Println(\"hey\")\n}\n```", true},
		{"false no lang", "```\npackage main;\n\nimport \"fmt\";\n\nfunc main() {\nfmt.Println(\"hey\")\n}\n```", false},
		{"true fake lang", "```package main;\n\nimport \"fmt\";\n\nfunc main() {\nfmt.Println(\"hey\")\n}\n```", false},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsCode(tt.input)
			if got != tt.expected {
				t.Error(got)
			}
		})
	}
}
