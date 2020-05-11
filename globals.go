package tamagochi

import "strings"


type JoinKind int

const (
	innerJoin JoinKind = iota
	leftJoin
	rightJoin
	fullJoin
)

const selectSql = "SELECT"
const fromSql = "FROM"
const joinSql = "JOIN"
const whereSql = "WHERE"


const space = " "
const end = ";"
const comma = ","
const All = "*"
const PlaceholderQmark = "?"
const PlaceholderDollar = "$"


type Expression interface {
	As(alias string) Expression
	apply(s *strings.Builder) *strings.Builder
}

type dyadicInterfaceOperator struct {
	left interface{}
	right interface{}
}

type polyadicExpressionOperator struct {
	exprs []Expression
}

type monadicStrFunction struct {
	alias string
	column Expression
}