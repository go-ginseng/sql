package sql

import "strings"

const (
	OP_EQ = "="
	OP_NE = "<>"
	OP_GT = ">"
	OP_GE = ">="
	OP_LT = "<"
	OP_LE = "<="
	OP_AN = "AND"
	OP_OR = "OR"
	OP_IN = "IN"
	OP_NI = "NOT IN"
	OP_LK = "LIKE"
	OP_NL = "NOT LIKE"
	OP_BT = "BETWEEN"
	OP_NU = "IS NULL"
	OP_NN = "IS NOT NULL"
)

type Clause struct {
	OP       string
	Field    string
	Value    interface{}
	Children []*Clause
}

func (c *Clause) Build() (string, []interface{}) {
	return _build(c, make([]interface{}, 0))
}

func _build(clause *Clause, values []interface{}) (string, []interface{}) {
	var stm string
	switch clause.OP {
	case OP_AN:
		stm, values = _processHasChildrenClause(clause, values)
	case OP_OR:
		stm, values = _processHasChildrenClause(clause, values)
	case OP_NU:
		stm = clause.Field + " " + clause.OP
	case OP_NN:
		stm = clause.Field + " " + clause.OP
	case OP_BT:
		cls := And(Gte(clause.Field, clause.Value.([]interface{})[0]), Lte(clause.Field, clause.Value.([]interface{})[1]))
		stm, values = _build(cls, values)
	default:
		stm = clause.Field + " " + clause.OP + " ?"
		values = append(values, clause.Value)
	}
	return stm, values
}

func _processHasChildrenClause(cls *Clause, values []interface{}) (string, []interface{}) {
	var buf []string
	for _, child := range cls.Children {
		s, v := _build(child, values)
		buf = append(buf, s)
		values = v
	}
	stm := "(" + strings.Join(buf, " "+cls.OP+" ") + ")"
	return stm, values
}

func Eq(field string, value interface{}) *Clause {
	return _newClause(OP_EQ, field, value)
}

func Neq(field string, value interface{}) *Clause {
	return _newClause(OP_NE, field, value)
}

func Gt(field string, value interface{}) *Clause {
	return _newClause(OP_GT, field, value)
}

func Gte(field string, value interface{}) *Clause {
	return _newClause(OP_GE, field, value)
}

func Lt(field string, value interface{}) *Clause {
	return _newClause(OP_LT, field, value)
}

func Lte(field string, value interface{}) *Clause {
	return _newClause(OP_LE, field, value)
}

func In(field string, values ...interface{}) *Clause {
	return _newClause(OP_IN, field, values)
}

func Nin(field string, values ...interface{}) *Clause {
	return _newClause(OP_NI, field, values)
}

func Lk(field string, value interface{}) *Clause {
	return _newClause(OP_LK, field, value)
}

func Nlk(field string, value interface{}) *Clause {
	return _newClause(OP_NL, field, value)
}

func Null(field string) *Clause {
	return &Clause{OP: OP_NU, Field: field}
}

func NotNull(field string) *Clause {
	return &Clause{OP: OP_NN, Field: field}
}

func And(clauses ...*Clause) *Clause {
	return _newClauseWithChildren(OP_AN, _removeEmptyClauses(clauses))
}

func Or(clauses ...*Clause) *Clause {
	return _newClauseWithChildren(OP_OR, _removeEmptyClauses(clauses))
}

func Between(field string, from interface{}, to interface{}) *Clause {
	return _newClause(OP_BT, field, []interface{}{from, to})
}

func _newClause(op string, field string, value interface{}) *Clause {
	return &Clause{OP: op, Field: field, Value: value}
}

func _newClauseWithChildren(op string, children []*Clause) *Clause {
	return &Clause{OP: op, Children: children}
}

func _removeEmptyClauses(clauses []*Clause) []*Clause {
	var result []*Clause
	for _, clause := range clauses {
		if clause != nil {
			result = append(result, clause)
		}
	}
	return result
}
