package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header  string
		want    string
		wantErr bool
		errIs   error
	}{
		"valid":         {header: "ApiKey abc123", want: "abc123"},
		"missing":       {header: "", wantErr: true, errIs: ErrNoAuthHeaderIncluded},
		"wrong prefix":  {header: "Bearer abc123", wantErr: true},
		"no key":        {header: "ApiKey", wantErr: true},
		"lowercase key": {header: "apikey abc123", wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			h := http.Header{}
			if tc.header != "" {
				h.Set("Authorization", tc.header)
			}
			got, err := GetAPIKey(h)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.errIs != nil && !errors.Is(err, tc.errIs) {
				t.Fatalf("err = %v, want %v", err, tc.errIs)
			}
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
