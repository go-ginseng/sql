package sql

import "strings"

const (
	OP_EQUAL         = "="
	OP_NOT_EQUAL     = "<>"
	OP_GREATER       = ">"
	OP_GREATER_EQUAL = ">="
	OP_LESS          = "<"
	OP_LESS_EQUAL    = "<="
	OP_AND           = "AND"
	OP_OR            = "OR"
	OP_IN            = "IN"
	OP_NOT_IN        = "NOT IN"
	OP_LIKE          = "LIKE"
	OP_NOT_LIKE      = "NOT LIKE"
	OP_NULL          = "IS NULL"
	OP_NOT_NULL      = "IS NOT NULL"
	OP_BETWEEN       = "BETWEEN"
)

type Statement struct {
	Operator string       `json:"operator"`
	Field    string       `json:"field,omitempty"`
	Value    interface{}  `json:"value,omitempty"`
	Children []*Statement `json:"children,omitempty"`
}

func (s *Statement) Build() (string, []interface{}) {
	return buildStatement(s, make([]interface{}, 0))
}

func buildStatement(s *Statement, values []interface{}) (string, []interface{}) {
	var stm string
	switch strings.ToUpper(s.Operator) {
	case OP_AND:
		stm, values = processHasChildrenClause(s, values)
	case OP_OR:
		stm, values = processHasChildrenClause(s, values)
	case OP_NULL:
		stm = s.Field + " " + s.Operator
	case OP_NOT_NULL:
		stm = s.Field + " " + s.Operator
	case OP_BETWEEN:
		cls := And(Gte(s.Field, s.Value.([]interface{})[0]), Lte(s.Field, s.Value.([]interface{})[1]))
		stm, values = buildStatement(cls, values)
	default:
		stm = s.Field + " " + s.Operator + " ?"
		values = append(values, s.Value)
	}
	return stm, values
}

func processHasChildrenClause(cls *Statement, values []interface{}) (string, []interface{}) {
	var buf []string
	for _, child := range cls.Children {
		s, v := buildStatement(child, values)
		buf = append(buf, s)
		values = v
	}
	stm := "(" + strings.Join(buf, " "+cls.Operator+" ") + ")"
	return stm, values
}
