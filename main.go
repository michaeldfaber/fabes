package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/michaeldfaber/fabes/compiler"
	"github.com/michaeldfaber/fabes/lexer"
	"github.com/michaeldfaber/fabes/object"
	"github.com/michaeldfaber/fabes/parser"
)

func main() {
	content, err := ioutil.ReadFile("dreaming.fbs")
    if err != nil {
        panic(err)
    }

	constants := []object.Object{}
	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

    text := string(content)
	l := lexer.New(text)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParseErrors(p.Errors())
	}

	comp := compiler.NewWithState(symbolTable, constants)
	err = comp.Compile(program)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Woops! Compilation failed:\n %s\n", err)
	}

	code := comp.Bytecode()
	constants = code.Constants
}

func printParseErrors(errors []string) {
	for _, msg := range errors {
		io.WriteString(os.Stdout, "\t"+msg+"\n")
	}
}