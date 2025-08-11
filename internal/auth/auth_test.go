package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKEY(t *testing.T) {

	testCase := map[string]struct {
		want       string
		authHeader string
		expectErr  error
	}{
		"success": {
			want:       "1234",
			authHeader: "ApiKey 12345",
			expectErr:  nil,
		},
		"fail format no value": {
			want:       "",
			authHeader: "ApiKey",
			expectErr:  ErrMalFormedAuthHeader,
		},
		"fail no header": {
			want:       "",
			authHeader: "",
			expectErr:  ErrNoAuthHeaderIncluded,
		},
	}

	for name, tc := range testCase {
		t.Run(name, func(t *testing.T) {

			header := make(http.Header)

			if name == "success" {
				header.Set("Authorization", tc.authHeader)
			}
			if name == "fail format no value" {
				header.Set("Authorization", tc.authHeader)
			}

			apiKey, err := GetAPIKey(header)

			if !errors.Is(err, tc.expectErr) {
				t.Fatalf("%s: expected error: %v, got: %v", name, tc.expectErr, err)
			}

			if apiKey != tc.want {
				t.Fatalf("%s: expected want: %v, got: %v", name, tc.want, apiKey)
			}

		})

	}

}
