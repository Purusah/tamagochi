package tamagochi

import (
	"fmt"
	"strings"
)

// COLUMNS
type eColumn struct {
	alias string
	column string
}

func (e eColumn) apply(s *strings.Builder) *strings.Builder {
	if e.alias != emptyStrSql {
		s.WriteString(fmt.Sprintf("%v AS %v", e.column, e.alias))
	} else {
		s.WriteString(e.column)
	}
	return s
}

func (e eColumn) As(alias string) Expression {
	e.alias = alias
	return e
}

func Column(column string) Expression {
	return eColumn{column: column}
}


// COUNT
type eCount struct {
	monadicStrFunction
}

func (e eCount) apply(s *strings.Builder) *strings.Builder {
	return unwrapMonadicFunc(s, countSql, e.monadicStrFunction)
}

func (e eCount) As(alias string) Expression {
	e.alias = alias
	return e
}

func Count(expr Expression) Expression {
	return eCount{monadicStrFunction{column: expr}}
}


// AND
type eAnd struct {
	polyadicExpressionOperator
}

func (e eAnd) apply(s *strings.Builder) *strings.Builder {
	return unwrapPolyadicExpr(s, andSql, e.polyadicExpressionOperator)
}

func (e eAnd) As(_ string) Expression {
	return e
}

func And(exprs ...Expression) Expression {
	return eAnd{polyadicExpressionOperator{exprs: exprs}}
}


// OR
type eOr struct {
	polyadicExpressionOperator
}

func (e eOr) apply(s *strings.Builder) *strings.Builder {
	return unwrapPolyadicExpr(s, orSql, e.polyadicExpressionOperator)
}

func (e eOr) As(_ string) Expression {
	return e
}

func Or(exprs ...Expression) Expression {
	return eOr{polyadicExpressionOperator{exprs: exprs}}
}


// NOT
//type eNot struct {
//	polyadicExpressionOperator
//}
//
//func (e eNot) apply(s *strings.Builder) *strings.Builder {
//	return unwrapPolyadicExpr(s, " AND ", e.polyadicExpressionOperator)
//}
//
//func (e eNot) As(_ string) Expression {
//	return e
//}
//
//func Not(exprs ...Expression) Expression {
//	return eNot{polyadicExpressionOperator{exprs: exprs}}
//}


// LT
type eLt struct {
	dyadicInterfaceOperator
}

func (e eLt) apply(s *strings.Builder) *strings.Builder {
	return unwrapDyadicExpr(s, "<", e.dyadicInterfaceOperator)
}

func (e eLt) As(_ string) Expression {
	return e
}

func Lt(left interface{}, right interface{}) Expression {
	return eLt{dyadicInterfaceOperator{left: left, right: right}}
}


// GT
type eGt struct {
	dyadicInterfaceOperator
}

func (e eGt) apply(s *strings.Builder) *strings.Builder {
	return unwrapDyadicExpr(s, ">", e.dyadicInterfaceOperator)
}

func (e eGt) As(_ string) Expression {
	return e
}

func Gt(left interface{}, right interface{}) Expression {
	return eGt{dyadicInterfaceOperator{left: left, right: right}}
}


// EQ
type eEq struct {
	dyadicInterfaceOperator
}

func (e eEq) apply(s *strings.Builder) *strings.Builder {
	return unwrapDyadicExpr(s, eqSql, e.dyadicInterfaceOperator)
}

func (e eEq) As(_ string) Expression {
	return e
}

func Eq(left interface{}, right interface{}) Expression {
	return eEq{dyadicInterfaceOperator{left: left, right: right}}
}


// NE
type eNe struct {
	dyadicInterfaceOperator
}

func (e eNe) apply(s *strings.Builder) *strings.Builder {
	return unwrapDyadicExpr(s, neSql, e.dyadicInterfaceOperator)
}

func (e eNe) As(_ string) Expression {
	return e
}

func Ne(left interface{}, right interface{}) Expression {
	return eNe{dyadicInterfaceOperator{left: left, right: right}}
}


// LE
type eLe struct {
	dyadicInterfaceOperator
}

func (e eLe) apply(s *strings.Builder) *strings.Builder {
	return unwrapDyadicExpr(s, leSql, e.dyadicInterfaceOperator)
}

func (e eLe) As(_ string) Expression {
	return e
}

func Le(left interface{}, right interface{}) Expression {
	return eLe{dyadicInterfaceOperator{left: left, right: right}}
}


// GE
type eGe struct {
	dyadicInterfaceOperator
}

func (e eGe) apply(s *strings.Builder) *strings.Builder {
	return unwrapDyadicExpr(s, geSql, e.dyadicInterfaceOperator)
}

func (e eGe) As(_ string) Expression {
	return e
}

func Ge(left interface{}, right interface{}) Expression {
	return eGe{dyadicInterfaceOperator{left: left, right: right}}
}


// ON
type eOn struct {
	dyadicInterfaceOperator
}

func (e eOn) apply(s *strings.Builder) *strings.Builder {
	// TODO FIX
	return unwrapDyadicExpr(s, eqSql, e.dyadicInterfaceOperator)
}

func (e eOn) As(_ string) Expression {
	return e
}

func On(left interface{}, right interface{}) Expression {
	return eOn{dyadicInterfaceOperator{left: left, right: right}}
}


// IS NULL
// Max
// Min
// Count
// Distinct
