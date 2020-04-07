package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/hscells/rougebit/blip"
	"github.com/hscells/rougebit/cmd/cart"
)

var cpu *blip.CPU

func main() {
	// Crate new CPU with a GPU.
	cpu := blip.NewCPU()
	cpu.Memory = cart.PaddleCart()
	cpu.GPU = blip.FaifaceGPU{
		Memory: cpu.Memory,
	}
	// Start executing instructions.
	go cpu.Run()
	// Start up the GPU.
	pixelgl.Run(cpu.GPU.Init)
}
