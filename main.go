package main

import (
	"fmt"
	"io/ioutil"
	"github.com/michaeldfaber/fabes/lexer"
	"github.com/michaeldfaber/fabes/token"
)

func main() {
	content, err := ioutil.ReadFile("dreaming.fbs")
    if err != nil {
        panic(err)
    }

    text := string(content)
	l := lexer.New(text)

	for {
		t := l.NextToken()
		if t.Type == token.EOF {
			break
		}
		
		if t.Type == token.Identifier {
			fmt.Printf("%d\t%s\t%s\n", l.Position, t.Type, t.Literal)
		} else {
			fmt.Printf("%d\t%s\t\t%s\n", l.Position, t.Type, t.Literal)
		}
	}
}