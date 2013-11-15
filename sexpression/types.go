package sexpression

type GLType int

const (
	TYPE_CONS GLType = iota
	TYPE_INT
	TYPE_FLOAT
	TYPE_STRING
	TYPE_SYMBOL
	TYPE_NIL
)

var TypeNames map[GLType]string = map[GLType]string{
	TYPE_CONS:   "cons",
	TYPE_INT:    "int",
	TYPE_FLOAT:  "float",
	TYPE_STRING: "string",
	TYPE_SYMBOL: "symbol",
	TYPE_NIL:    "nil",
}

func GetTypeName(se SExpression) string {
	return TypeNames[se.ExprType()]
}

type Typer interface {
	AsFloat() (float64, error)
	AsInt() (int64, error)
	AsString() (string, error)
}
