package tamagochi

import (
	"strings"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"already_snake", "already_snake"},
		{"A", "a"},
		{"AA", "aa"},
		{"AaAa", "aa_aa"},
		{"HTTPRequest", "http_request"},
		{"BatteryLifeValue", "battery_life_value"},
		{"Id0Value", "id0_value"},
		{"ID0Value", "id0_value"},
	}
	for _, test := range tests {
		have := toSnakeCase(test.input)
		if have != test.want {
			t.Errorf("input=%q:\nhave: %q\nwant: %q", test.input, have, test.want)
		}
	}
}

func TestUnwrapMonadicFunc(t *testing.T) {
	s := strings.Builder{}
	testStr := "abc"

	cnt := monadicStrFunction{column: Column(testStr)}
	if unwrapMonadicFunc(&s, "COUNT", cnt).String() != "COUNT(abc)" {
		t.Error("Error monadicStrFunction")
	}
	s.Reset()

	cnt = monadicStrFunction{alias: "CNT", column: Column(testStr)}
	if unwrapMonadicFunc(&s, "COUNT", cnt).String() != "COUNT(abc) AS CNT" {
		t.Error("Error monadicStrFunction with alias")
	}
}