package forth

import (
	"fmt"

	"github.com/joshbarrass/goforth/pkg/stacks"
)

// DefaultCoded defines the functions that are implemented in go They
// take an Interpreter pointer as an argument and return an error
// Any return values should be set in Interpreter.ReturnValue
var DefaultCoded = map[string]func(*Interpreter) error{
	"bye": GracefulShutdown,
	".":   IntStackPop,
	".s":  IntStackPrint,
	"+":   IntAdd,
	"-":   IntSub,
	"*":   IntMul,
	"/":   IntDiv,
	"mod": IntMod,
}

// DefaultForth defines the functions that are implemented in forth
// These will take precedence over functions implemented in go
var DefaultForth = map[string]string{}

/* Coded function definitions */

// GracefulShutdown sets the internal status to shutdown
// ( -- )
func GracefulShutdown(i *Interpreter) error {
	i.InternalStatus = STATUS_SHUTDOWN
	return nil
}

// IntStackPop pops a value off the stack to be returned
// ( n -- )
func IntStackPop(i *Interpreter) error {
	s := i.IntStack

	val, err := s.Pop()
	if err != nil {
		return err
	}
	i.ReturnValue = val
	return nil
}

// IntStackPrint prints the length and state of the stack
// ( -- )
func IntStackPrint(i *Interpreter) error {
	s := i.IntStack
	length := s.Len()
	toReturn := fmt.Sprintf("<%d>", length)
	if length > 0 {
		toReturn += "  Top-->"
	}

	// walk through stack items
	node := s.Top
	for n := 0; n < length; n++ {
		if node == nil {
			break
		}

		toReturn += fmt.Sprintf("%d,", node.Value)
		node = node.Previous
	}

	if length > 0 {
		toReturn = toReturn[:len(toReturn)-1]
	}
	i.ReturnValue = toReturn

	return nil
}

// getTwoInts pulls two numbers from the stack
// this is a common operation in arithmetic
func getTwoInts(s *stacks.IntStack) (val1, val2 int, err error) {
	// Verify that there are at least 2 items on the stack
	if s.Len() < 2 {
		err = ErrorTooFewItems
		return
	}
	val1, _ = s.Pop()
	val2, _ = s.Pop()
	return
}

// IntAdd adds the top two numbers of the stack
// ( n n -- n )
func IntAdd(i *Interpreter) error {
	s := i.IntStack

	val1, val2, err := getTwoInts(s)
	if err != nil {
		return err
	}

	s.Push(val1 + val2)
	return nil
}

// IntSub subtracts the second number on the stack from the first
// ( n n -- n )
func IntSub(i *Interpreter) error {
	s := i.IntStack

	val1, val2, err := getTwoInts(s)
	if err != nil {
		return err
	}

	s.Push(val2 - val1)
	return nil
}

// IntMul multiplies the top two numbers of the stack
// ( n n -- n )
func IntMul(i *Interpreter) error {
	s := i.IntStack

	val1, val2, err := getTwoInts(s)
	if err != nil {
		return err
	}

	s.Push(val1 * val2)
	return nil
}

// IntDiv divides the second number on the stack by the first
// ( n n -- n )
func IntDiv(i *Interpreter) error {
	s := i.IntStack

	val1, val2, err := getTwoInts(s)
	if err != nil {
		return err
	}

	s.Push(val2 / val1)
	return nil
}

// IntMod divides the second number on the stack by the first and
// pushes the remainder
// ( n n -- n )
func IntMod(i *Interpreter) error {
	s := i.IntStack

	val1, val2, err := getTwoInts(s)
	if err != nil {
		return err
	}

	s.Push(val2 % val1)
	return nil
}
