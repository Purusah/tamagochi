package tamagochi

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var stringType = reflect.TypeOf("")
var intType = reflect.TypeOf(0)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func getType(entity interface{}) reflect.Type {
	return reflect.TypeOf(entity)
}

func getStrFromStrType(i interface{}) string {
	return reflect.ValueOf(i).String()
}

func getStrFromIntType(i interface{}) string {
	return strconv.Itoa(i.(int))
}

func getJoinType(j JoinKind) string {
	switch j {
	case innerJoin:
		return "INNER"
	case leftJoin:
		return "LEFT"
	case rightJoin:
		return "RIGHT"
	case fullJoin:
		return "FULL"
	default:
		panic("Unknown JOIN type")
	}
}

func unwrapInterface(s *strings.Builder, t reflect.Type, v interface{}) *strings.Builder {
	if t == stringType {
		s.WriteString(getStrFromStrType(v))
	} else if t == intType {
		s.WriteString(getStrFromIntType(v))
	} else {
		s = v.(Expression).apply(s)
	}
	return s
}

func unwrapMonadicFunc(s *strings.Builder, o string, e monadicStrFunction) *strings.Builder {
	s.WriteString(o)
	s.WriteString("(")
	e.column.apply(s)
	s.WriteString(")")
	if e.alias != "" {
		s.WriteString(fmt.Sprintf(" AS %v", e.alias))
	}
	return s
}

func unwrapDyadicExpr(s *strings.Builder, o string, e dyadicInterfaceOperator) *strings.Builder {
	for i, f := range []interface{}{e.left, e.right} {
		t := getType(f)
		s = unwrapInterface(s, t, f)
		s.WriteString(space)
		if i + 1 == 1 {
			s.WriteString(o)
			s.WriteString(space)
		}
	}
	return s
}

func unwrapPolyadicExpr(s *strings.Builder, o string, e polyadicExpressionOperator) *strings.Builder {
	totalField := len(e.exprs)
	for i, f := range e.exprs {
		s = f.apply(s)
		if i + 1 != totalField {
			s.WriteString(o)
		}
	}
	return s
}
