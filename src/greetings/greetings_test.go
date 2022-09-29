package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Dotnet9"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Dotnet9")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Dotnet9") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}