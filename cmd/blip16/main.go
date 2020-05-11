package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/hscells/blip16/blip"
	"io/ioutil"
	"os"
)

var cpu *blip.CPU

func main() {

	m := blip.NewMemory()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	copy(m.Mem(), b)

	// Crate new CPU with a GPU.
	cpu := blip.NewCPU()
	cpu.Memory = m
	cpu.GPU = blip.FaifaceGPU{
		Memory: cpu.Memory,
	}
	// Start executing instructions.
	go cpu.Run()
	// Start up the GPU.
	pixelgl.Run(cpu.GPU.Init)
}
