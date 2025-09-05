package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantKey     string
		expectError bool
	}{
		{
			name: "Valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey 12345"},
			},
			wantKey:     "12345",
			expectError: false,
		},
		{
			name:        "Missing Authorization header",
			headers:     http.Header{},
			wantKey:     "",
			expectError: true,
		},
		{
			name: "Malformed Authorization header - wrong scheme",
			headers: http.Header{
				"Authorization": []string{"Bearer 12345"},
			},
			wantKey:     "",
			expectError: true,
		},
		{
			name: "Malformed Authorization header - missing key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			wantKey:     "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.expectError {
				t.Errorf("GetAPIKey() error = %v, expectError %v", err, tt.expectError)
			}
			if got != tt.wantKey {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.wantKey)
			}
		})
	}
}
