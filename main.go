package main

import (
	"fmt"
	"os"
	lexer "github.com/michaeldfaber/fabes/lexer"
)

func main() {
	file, err := os.Open("dreaming.fbs")
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer(file)
	for {
		pos, tok, lit := l.Lex()
		if tok == lexer.EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.Line, pos.Column, tok, lit)
	}
}