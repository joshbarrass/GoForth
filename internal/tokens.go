package forth

import (
	"errors"
	"strconv"
	"strings"
)

const (
	DIGITS    = "0123456789-"
	WORDCHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.!?<>,.@:;+-*/$%^&#"
)

func (i *Interpreter) HandleToken(token string) error {
	// identify if number or word
	switch {
	case strings.ContainsAny(DIGITS, string(token[0])):
		// is a number
		return i.handleNumber(token)
	case strings.ContainsAny(WORDCHARS, string(token[0])):
		// is a word
		return i.handleWord(token)
	}

	return ErrorTokenNotCaught
}

// handleNumber redirects numbers to handleInteger or handleFloat
func (i *Interpreter) handleNumber(token string) error {
	// integer or float?
	if token[len(token)-1] == 'e' {
		return i.handleFloat(token)
	}
	return i.handleInteger(token)
}

// handleInteger adds to the integer stack
func (i *Interpreter) handleInteger(token string) error {
	n, err := strconv.Atoi(token)
	i.IntStack.Push(n)
	return err
}

// TODO: Write handleFloat

// handleFloat adds to the float stack
func (i *Interpreter) handleFloat(token string) error {
	// explicity do nothing
	switch false {
	case token == token:
		return nil
	}
	return errors.New("floats not yet implemented")
}

// handleWord calls the function responsible corresponding to each
// word
func (i *Interpreter) handleWord(token string) error {
	// check to see if it exists in the word dictionary
	forthWord, ok := i.WordDictionary[token]
	if ok {
		// Execute the code for the word
		tokens := strings.Split(forthWord, " ")
		for _, token := range tokens {
			err := i.HandleToken(token)
			if err != nil {
				return err
			}
		}
		return nil
	}
	// check to see if it exists in the go dictionary
	goWord, ok := DefaultCoded[token]
	if !ok {
		// word is undefined, return an error
		return ErrorWordDoesNotExist
	}
	// otherwise, execute the function and return any errors
	return goWord(i)
}
