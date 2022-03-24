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
				t.Errorf("containsCode() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestGetCode(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedCode string
		expectedLang string
	}{
		{"true", "```golang\npackage main;\n\nimport \"fmt\";\n\nfunc main() {\nfmt.Println(\"hey\")\n}\n```", "package main; import \"fmt\"; func main() { fmt.Println(\"hey\") }", "golang"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotLang := getCode(tt.input)
			if gotCode != tt.expectedCode {
				t.Errorf("getCode() code = %v, expected %v", gotCode, tt.expectedCode)
			}
			if gotLang != tt.expectedLang {
				t.Errorf("getCode() lang = %v, expected %v", gotLang, tt.expectedLang)
			}
		})
	}
}
