package blip

import (
	"fmt"
	"time"
)

type Clock struct {
	M time.Duration
}

type Registers struct {
	A, B, C uint8
	X, Y, Z uint16
	PC, SP  uint16
}

type CPU struct {
	Registers
	*Memory
	GPU
	CPUClock Clock
	GPUClock Clock
}

func NewCPU() *CPU {
	return &CPU{
		Memory: NewMemory(),
	}
}

func (c *CPU) Execute(i uint8) error {
	if instruction, ok := Instructions[i]; ok {
		instruction(c)
	}
	c.PC++
	return nil
}

func (c *CPU) Run() {
	c.PC = OffCD
	for _, v := range c.Memory.Mem()[c.PC : c.PC+16] {
		fmt.Printf("%04x\n", v)
	}
	for {
		i := c.Memory.Mem()[c.PC]
		err := c.Execute(i)
		if err != nil {
			panic(err)
		}
		time.Sleep(c.CPUClock.M)
		c.CPUClock.M = 0
		err = c.GPU.Render()
		if err != nil {
			panic(err)
		}

	}
}

func cycle(cpu *CPU, duration time.Duration) {
	cpu.CPUClock.M += duration * 500 * time.Microsecond
}
