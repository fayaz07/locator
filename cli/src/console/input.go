package console

import (
	"os"

	"golang.org/x/term"
)

func ReadOneChar() byte {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		Error("Error reading input, please raise an issue at https://github.com/fayaz07/locator/issues/new")
		panic("")
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		Error("Error reading input, please raise an issue at https://github.com/fayaz07/locator/issues/new")
		panic("")
	}
	return b[0]
}
