package tamagochi

import (
	"fmt"
	"strconv"
	"strings"
)

// SELECT
type eSelect struct {
	Fields []interface{}
}

func (e eSelect) As(_ string) Expression {
	return e
}

func (e eSelect) apply(s *strings.Builder) *strings.Builder {
	s.WriteString(selectSql)
	s.WriteString(space)
	totalField := len(e.Fields)

	for i, f := range e.Fields {
		t := getType(f)
		if t == stringType {
			s.WriteString(getStrFromStrType(f))
		} else if t == intType {
			s.WriteString(getStrFromIntType(f))
		} else {
			s = f.(Expression).apply(s)
		}
		if i + 1 != totalField {
			s.WriteString(comma)
			s.WriteString(space)
		}
	}
	return s
}


// FROM
type eFrom struct {
	Table interface{}
}

func (e eFrom) As(_ string) Expression {
	return e
}

func (e eFrom) apply(s *strings.Builder) *strings.Builder {
	s.WriteString(fmt.Sprintf(
		"%v %v",
		fromSql,
		toSnakeCase(getType(e.Table).Name())),
	)
	return s
}


// JOIN
type eJoin struct {
	Table interface{}
	On Expression
	Kind JoinKind
}

func (e eJoin) apply(s *strings.Builder) *strings.Builder {
	s.WriteString(
		fmt.Sprintf("%v %v %v",
		getJoinType(e.Kind),
		joinSql,
		toSnakeCase(getType(e.Table).Name())),
	)
	return e.On.apply(s)
}

func (e eJoin) As(_ string) Expression {
	return e
}


// WHERE
type eWhere struct {
	Filters Expression
}

func (e eWhere) apply(s *strings.Builder) *strings.Builder {
	s.WriteString(whereSql)
	s.WriteString(space)
	return e.Filters.apply(s)
}

func (e eWhere) As(_ string) Expression {
	return e
}


// LIMIT
type eLimit struct {
	Limit int
}

func (e eLimit) apply(s *strings.Builder) *strings.Builder {
	s.WriteString(limitSql)
	s.WriteString(space)
	s.WriteString(strconv.Itoa(e.Limit))
	return s
}

func (e eLimit) As(_ string) Expression {
	return e
}
