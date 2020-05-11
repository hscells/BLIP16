package blip

import (
	"time"
)

type Clock struct {
	M time.Duration
}

type Registers struct {
	A      uint8
	PC, SP uint16
}

type CPU struct {
	Registers
	*Memory
	Cart *Memory
	GPU
	CPUClock Clock
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
	for {
		i := c.Memory.Mem()[c.PC]
		if c.PC < OffCD {
			for {
				continue
			}
		}
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
