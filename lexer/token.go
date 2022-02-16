package lexer

type Token struct {
	Type		string
	Literal		string
}

const (
	Illegal 		= "Illegal"
	EOF     		= "EOF"

	// Identifiers + Literals
	Identifier 		= "Identifier"
	Int        		= "Int"
	String     		= "String"

	// Operators
	Assign   			= "="
	Plus     			= "+"
	Minus    			= "-"
	Bang     			= "!"
	Asterisk 			= "*"
	Slash    			= "/"
	Equal    			= "=="
	NotEqual 			= "!="
	LessThan    		= "<"
	LessThanOrEqual		= "<="
	GreaterThan 		= ">"
	GreaterThanOrEqual	= ">="

	// Delimiters
	Comma     		= ","
	Semicolon 		= ";"
	Colon     		= ":"

	LeftParen    	= "("
	RightParen   	= ")"
	LeftBrace    	= "{"
	RightBrace   	= "}"
	LeftBracket  	= "["
	RightBracket 	= "]"

	// Keywords
	Function 		= "Function"
	Var      		= "Var"
	True     		= "True"
	False    		= "False"
	If       		= "If"
	Else     		= "Else"
	Return   		= "Return"
)

var keywords = map[string] string {
	"function":     Function,
	"var":    		Var,
	"true":   		True,
	"false":  		False,
	"if":     		If,
	"else":   		Else,
	"return": 		Return,
}

func LookupIdentifierType(identifier string) string {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return Identifier
}
