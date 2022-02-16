package main

import (
	"fmt"
	"io/ioutil"
	lexer "github.com/michaeldfaber/fabes/lexer"
)

func main() {
	content, err := ioutil.ReadFile("dreaming.fbs")
    if err != nil {
        panic(err)
    }

    text := string(content)
	l := lexer.New(text)

	for {
		token := l.NextToken()
		if token.Type == lexer.EOF {
			break
		}
		
		if token.Type == lexer.Identifier {
			fmt.Printf("%d\t%s\t%s\n", l.Position, token.Type, token.Literal)
		} else {
			fmt.Printf("%d\t%s\t\t%s\n", l.Position, token.Type, token.Literal)
		}
	}
}