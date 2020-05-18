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
		{"AA", "aa"},
		{"AaAa", "aa_aa"},
		{"HTTPRequest", "http_request"},
		{"BatteryLifeValue", "battery_life_value"},
		{"Id0Value", "id0_value"},
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
	if unwrapMonadicFunc(&s, countSql, cnt).String() != "COUNT(abc)" {
		t.Error("Error monadicStrFunction")
	}
	s.Reset()

	cnt = monadicStrFunction{alias: "CNT", column: Column(testStr)}
	if unwrapMonadicFunc(&s, countSql, cnt).String() != "COUNT(abc) AS CNT" {
		t.Error("Error monadicStrFunction with alias")
	}
}

func TestUnwrapDyadicExpr(t *testing.T) {
	s := strings.Builder{}

	expr := dyadicInterfaceOperator{
		left:  "val1",
		right: "val2",
	}
	if unwrapDyadicExpr(&s, "=", expr).String() != "val1 = val2" {
		t.Error("Error dyadicInterfaceOperator")
	}
	s.Reset()
}

func TestUnwrapPolyadicExpr(t *testing.T) {
	s := strings.Builder{}

	expr := polyadicExpressionOperator{
		exprs: []Expression{
			Column("val1"),
			Column("val2"),
			Column("val3"),
		},
	}
	if unwrapPolyadicExpr(&s, andSql, expr).String() != "val1 AND val2 AND val3" {
		t.Error("Error polyadicExpressionOperator")
	}
	s.Reset()
}