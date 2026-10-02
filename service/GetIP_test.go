package service

import (
	"net/http"
	"testing"
)

func TestGetIP(t *testing.T) {
	tests := []struct {
		name       string
		remoteAddr string
		expected   string
	}{
		{
			name:       "IPv4 with port",
			remoteAddr: "192.168.1.1:1234",
			expected:   "192.168.1.1",
		},
		{
			name:       "IPv6 with port",
			remoteAddr: "[2001:db8::1]:8080",
			expected:   "2001:db8::1",
		},
		{
			name:       "missing port",
			remoteAddr: "192.168.1.1",
			expected:   "<nil>",
		},
		{
			name:       "empty address",
			remoteAddr: "",
			expected:   "<nil>",
		},
		{
			name:       "invalid address format",
			remoteAddr: "invalid-addr:port",
			expected:   "<nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				RemoteAddr: tt.remoteAddr,
			}
			if got := GetIP(req); got != tt.expected {
				t.Errorf("GetIP() = %v, want %v", got, tt.expected)
			}
		})
	}
}
