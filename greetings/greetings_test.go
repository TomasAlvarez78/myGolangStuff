package greetings

import (
    "testing"
    "regexp"
)

// TestTasuketeName calls greetings.Tasukete with a name, checking
// for a valid return value.
func TestTasuketeName(t *testing.T) {
    name := "Mihari"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Tasukete("Mihari")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Tasukete("Mihari") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestTasuketeEmpty calls greetings.Tasukete with an empty string,
// checking for an error.
func TestTasuketeEmpty(t *testing.T) {
    msg, err := Tasukete("")
    if msg != "" || err == nil {
        t.Fatalf(`Tasukete("") = %q, %v, want "", error`, msg, err)
    }
}