package shared

import (
	"net/http"
	"testing"
)

func TestValidateToken(t *testing.T) {
	type args struct {
		domain   string
		audience string
		req      *http.Request
	}

	req, _ := http.NewRequest("GET", "http://test.com", nil)
	token := "Bearer "
	req.Header.Set("authorization", token)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestValidateToken 01",
			args: args{
				audience: "xalex",
				domain:   "xalex.us.auth0.com",
				req:      req,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateRequest(tt.args.domain, tt.args.audience, tt.args.req); got != tt.want {
				t.Errorf("ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
