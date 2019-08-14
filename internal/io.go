package forth

import "fmt"

const (
	lineDelimeter = '\n'
	linePrompt    = ""
)

// ReadLine provides an interface for reading a line of input.
// The returned string should not contain the line delimeter
func (i *Interpreter) ReadLine() (text string, err error) {
	// Allow custom line prompts -- might be useful at some point
	if linePrompt != "" {
		fmt.Println(linePrompt)
	}

	text, err = i.Reader.ReadString(lineDelimeter)
	text = text[:len(text)-1]
	return
}
