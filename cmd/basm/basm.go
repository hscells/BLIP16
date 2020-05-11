package main

import (
	"github.com/hscells/blip16/basm"
	"os"
)

func main() {
	toks := basm.Tokenise(os.Stdin)
	data, err := basm.Parse(toks)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}
