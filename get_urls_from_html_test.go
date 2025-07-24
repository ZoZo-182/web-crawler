package main

import (
	"reflect"
	"testing"
)

func TestGetURLFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "abs and rel urls",
			inputURL: "https://blog.boot.dev",
			inputBody: `
            <html>
            <body>
            <a href="/path/one">
            <span>Boot.dev</span>
            </a>
            <a href="https://other.com/path/one">
            <span>Boot.dev</span>
            </a>
            </body>
            </html>
            `,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs %v, got URLs %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
