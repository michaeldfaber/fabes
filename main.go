package main

import (
	"io/ioutil"
	"github.com/michaeldfaber/fabes/lexer"
	"github.com/michaeldfaber/fabes/parser"
)

func main() {
	content, err := ioutil.ReadFile("dreaming.fbs")
    if err != nil {
        panic(err)
    }

    text := string(content)
	l := lexer.New(text)
	p := parser.New(l);

	p.ParseProgram();
}