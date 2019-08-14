package forth

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joshbarrass/goforth/pkg/stacks"
)

// NewInterpreter creates a new forth Interpreter with a stack and
// reader
func NewInterpreter(intStack *stacks.IntStack, floatStack *stacks.FloatStack) *Interpreter {
	reader := bufio.NewReader(os.Stdin)
	return &Interpreter{
		IntStack:       intStack,
		FloatStack:     floatStack,
		Reader:         reader,
		WordDictionary: DefaultForth,
	}
}

// Interpreter is a forth interpreter
type Interpreter struct {
	IntStack   *stacks.IntStack
	FloatStack *stacks.FloatStack
	Reader     *bufio.Reader

	// WordDictionary contains all the non-go words
	WordDictionary map[string]string

	// ReturnValue contains the value returned by something
	ReturnValue interface{}

	InternalStatus int
}

// StartRepl starts REPL execution of forth code
func (i *Interpreter) StartRepl() {
	// dirty infinite loop -- might want to add graceful shutdowns
	// later
	for true {
		line, err := i.ReadLine()
		if err != nil {
			log.Fatalf("Unable to read line: %s", err)
		}

		// split the line into its tokens
		tokens := strings.Split(line, " ")

		for _, token := range tokens {
			err = i.HandleToken(token)
			// handle internal status before error
			switch i.InternalStatus {
			case STATUS_SHUTDOWN:
				fmt.Printf("  Exiting...\n")
				return
			}
			if err != nil {
				fmt.Printf("  Error!: %s\n", err)

				// specific error handling
				switch err {
				case ErrorWordDoesNotExist:
					fmt.Printf("    \"%s\"\n", token)
				case ErrorTokenNotCaught:
					fmt.Printf("    \"%s\"\n", string(token[0]))
				}

				break
			}
		}
		if err == nil {
			if i.ReturnValue != nil {
				fmt.Printf("  %v OK!\n", i.ReturnValue)
			} else {
				fmt.Printf("  OK!\n")
			}
		}
		i.ReturnValue = nil
	}
}
