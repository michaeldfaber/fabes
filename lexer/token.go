package lexer

type Token int

const (
	EOF = iota
	ILLEGAL
	IDENT
	INT
	SEMI // ;

	// Infix notation: X + Y. Operators are written in-between their operands.
	ADD // +
	SUB // -
	MUL // *
	DIV // /

	ASSIGN // =
)

var tokens = []string{
	EOF:     	"EOF",
	ILLEGAL: 	"ILLEGAL",
	IDENT:   	"IDENT",
	INT:     	"INT",
	SEMI:    	";",

	ADD: 		"+",
	SUB: 		"-",
	MUL: 		"*",
	DIV: 		"/",

	ASSIGN: 	"=",
}

func (t Token) String() string {
	return tokens[t]
}