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
const limitSql = "LIMIT"

const innerJoinSql = "INNER"
const leftJoinSql = "LEFT"
const rightJoinSql = "RIGHT"
const fullJoinSql = "FULL"
const onSql = "ON"

const andSql = "AND"
const orSql = "OR"

const countSql = "COUNT"

const lParenthesisSql = "("
const rParenthesisSql = ")"
const neSql = "<>"
const leSql = "<="
const geSql = ">="
const eqSql = "="
const emptyStrSql = ""
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