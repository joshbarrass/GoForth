package forth

import (
	"errors"
)

/* General Purpose Errors */

// ErrorTokenNotCaught is returned if the token does not match the
// switch
var ErrorTokenNotCaught = errors.New("token was not caught by switch")

// ErrorWordDoesNotExist is returned if the word is not defined in the
// word dictionaries
var ErrorWordDoesNotExist = errors.New("word does not exist in word dictionary")

var ErrorTooFewItems = errors.New("too few items on stack")
