package hellowword

import (
	"github.com/sun20181102/test1"
)

// Hello returns a greeting for the named person.
func HelloWorld() string {
	// Return a greeting that embeds the name in a message.
	message := test1.Hello("hello world")
	return message
}
