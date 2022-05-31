package test

import (
	"strings"
	"testing"
)

func Test_GetWinner(t *testing.T) {

	testData := []struct {
		name   string
		votes  []string
		winner string
	}{
		{"TC1", []string{"Raj", "Antony", "Raj", "Antony", "Ravi", "Antony"}, "Antony"},
		{"TC2", nil, ""},
		{"TC3", []string{}, ""},
		{"TC4", []string{"", ""}, ""},
		{"TC4", []string{"Raj", ""}, "Raj"},

		{"TC4", []string{"Raj", "Raj", "Raj"}, "Raj"},

		{"TC4", []string{"Raj", "Raj", "Antony", "Antony"}, "Raj"},
		{"TC4", []string{"Antony", "Antony", "Raj", "Raj"}, "Antony"},

		{"TC4", []string{"Raj", "Raj", "Antony", "Antony", "Antony"}, "Antony"},
	}

	for _, td := range testData {
		winner := GetWinner(td.votes)
		if !strings.EqualFold(winner, td.winner) {
			t.Errorf("%s: Exp: %s, Got: %s", td.name, td.winner, winner)
		}
	}
}
