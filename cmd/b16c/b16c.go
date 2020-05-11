package main

import (
	"github.com/hscells/blip16/blip"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("not enough arguments")
	}

	m := make([]uint8, 0xffff)

	datf, err := os.Open(args[1])
	if err != nil {
		panic(err)
	}

	dat, err := ioutil.ReadAll(datf)
	if err != nil {
		panic(err)
	}

	objf, err := os.Open(args[2])
	if err != nil {
		panic(err)
	}

	obj, err := ioutil.ReadAll(objf)
	if err != nil {
		panic(err)
	}

	copy(m[blip.OffBD:], dat)
	copy(m[blip.OffCD:], obj)

	os.Stdout.Write(m)
}
