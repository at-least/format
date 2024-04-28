package format_test

import (
	"fmt"
	"testing"

	"github.com/at-least/format"
)

// Test go name formatting.
func TestGoName(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"update_user", "UpdateUser"},
		{"user_id", "UserID"},
		{"ip", "IP"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("GoName(%q)", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := format.GoName(tt.s)
			if ans != tt.want {
				t.Errorf("got %q, want %q", ans, tt.want)
			}
		})
	}
}

// Test js name formatting.
func TestJsName(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"update_user_settings", "updateUserSettings"},
		{"user_id", "userId"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("JsName(%q)", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := format.JsName(tt.s)
			if ans != tt.want {
				t.Errorf("got %q, want %q", ans, tt.want)
			}
		})
	}
}
