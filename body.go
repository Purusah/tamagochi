package tamagochi


func (e Entity) From(table interface{}) Entity {
	e.expressions = append(e.expressions, eFrom{Table: table})
	return e
}

func (e Entity) Where(expr Expression) Entity {
	e.expressions = append(e.expressions, eWhere{Filters: expr})
	return e
}

func (e Entity) InnerJoin(table interface{}, expr Expression) Entity {
	e.expressions = append(e.expressions, eJoin{Table: table, On: expr, Kind: innerJoin})
	return e
}

func (e Entity) LeftJoin(table interface{}, expr Expression) Entity {
	e.expressions = append(e.expressions, eJoin{Table: table, On: expr, Kind: leftJoin})
	return e
}

func (e Entity) RightJoin(table interface{}, expr Expression) Entity {
	e.expressions = append(e.expressions, eJoin{Table: table, On: expr, Kind: rightJoin})
	return e
}

func (e Entity) FullJoin(table interface{}, expr Expression) Entity {
	e.expressions = append(e.expressions, eJoin{Table: table, On: expr, Kind: fullJoin})
	return e
}

func (e Entity) Limit (limit int) Entity {
	e.expressions = append(e.expressions, eLimit{Limit: limit})
	return e
}

func (e Entity) Sql() string {
	for i, f := range e.expressions {
		if i != 0 {
			e.builder.WriteString(space)
		}
		e.builder = *f.apply(&e.builder)
	}
	e.builder.WriteString(end)
	e.expressions = nil  // Or not
	defer e.builder.Reset()
	return e.builder.String()
}
