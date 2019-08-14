package main

import (
	"fmt"

	forth "github.com/joshbarrass/goforth/internal"
	"github.com/joshbarrass/goforth/pkg/stacks"
)

// Version is the version of the interpreter
const Version = "0.0.2"

// StartText is the text produced at startup
const StartText = `GoForth version %s
Type "bye" to exit

`

// the stacks used by the interpreter
var (
	IntStack   *stacks.IntStack
	FloatStack *stacks.FloatStack
)

func main() {
	// set up new, empty stacks
	IntStack = stacks.NewIntStack()
	FloatStack = stacks.NewFloatStack()

	// define a new interpreter
	interpreter := forth.NewInterpreter(IntStack, FloatStack)

	// print the start text
	fmt.Printf(StartText, Version)

	// start the REPL
	interpreter.StartRepl()
}
