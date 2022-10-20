package sql

// Equal (=)
func Eq(field string, value interface{}) *Statement {
	return newStatement(OP_EQUAL, field, value)
}

// Not Equal (<>)
func Neq(field string, value interface{}) *Statement {
	return newStatement(OP_NOT_EQUAL, field, value)
}

// Greater (>)
func Gt(field string, value interface{}) *Statement {
	return newStatement(OP_GREATER, field, value)
}

// Greater or Equal (>=)
func Gte(field string, value interface{}) *Statement {
	return newStatement(OP_GREATER_EQUAL, field, value)
}

// Less (<)
func Lt(field string, value interface{}) *Statement {
	return newStatement(OP_LESS, field, value)
}

// Less or Equal (<=)
func Lte(field string, value interface{}) *Statement {
	return newStatement(OP_LESS_EQUAL, field, value)
}

// In (IN)
func In(field string, values interface{}) *Statement {
	return newStatement(OP_IN, field, values)
}

// Not In (NOT IN)
func Nin(field string, values interface{}) *Statement {
	return newStatement(OP_NOT_IN, field, values)
}

// Like (LIKE)
func Lk(field string, value interface{}) *Statement {
	return newStatement(OP_LIKE, field, value)
}

// Not Like (NOT LIKE)
func Nlk(field string, value interface{}) *Statement {
	return newStatement(OP_NOT_LIKE, field, value)
}

// Null (IS NULL)
func Null(field string) *Statement {
	return newStatement(OP_NULL, field, nil)
}

// Not Null (IS NOT NULL)
func NotNull(field string) *Statement {
	return newStatement(OP_NOT_NULL, field, nil)
}

// And (AND)
func And(clauses ...*Statement) *Statement {
	return newStatementWithChildren(OP_AND, removeEmptyClauses(clauses))
}

// Or (OR)
func Or(clauses ...*Statement) *Statement {
	return newStatementWithChildren(OP_OR, removeEmptyClauses(clauses))
}

// Between (BETWEEN)
func Between(field string, from interface{}, to interface{}) *Statement {
	return newStatement(OP_BETWEEN, field, []interface{}{from, to})
}

func newStatement(op string, field string, value interface{}) *Statement {
	return &Statement{Operator: op, Field: field, Value: value}
}

func newStatementWithChildren(op string, children []*Statement) *Statement {
	return &Statement{Operator: op, Children: children}
}

func removeEmptyClauses(clauses []*Statement) []*Statement {
	var result []*Statement
	for _, clause := range clauses {
		if clause != nil {
			result = append(result, clause)
		}
	}
	return result
}
