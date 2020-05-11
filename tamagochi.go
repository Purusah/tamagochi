package tamagochi

import "strings"

type Entity struct {
	expressions []Expression
	//fields []interface{}
	builder strings.Builder
}

func Select(fields ...interface{}) Entity {
	e := Entity{}
	e.builder = strings.Builder{}
	e.expressions = append(e.expressions, eSelect{fields})
	return e
}

// UPDATE
// DELETE
// INSERT