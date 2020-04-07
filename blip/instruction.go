package blip

type Instruction func(cpu *CPU)

var Instructions = map[uint8]Instruction{
	0x00: NOP, 0x01: JP16, 0x02: JPEQXY16, 0x03: JPGTXY16, 0x04: JPLTXY16,
	0x10: LDA8, 0x11: INCA, 0x12: DECA, 0x13: SUB8, 0x14: LDf16A,
	0x20: LDB8, 0x21: INCB, 0x22: DECB, 0x23: SUBB, 0x24: LDf16B,
}

func NOP(cpu *CPU) {
	cycle(cpu, 4)
}

func JP16(cpu *CPU) {
	cpu.PC = cpu.Read16(cpu.PC + 1)
	cycle(cpu, 16)
}

func JPEQXY16(cpu *CPU) {
	if cpu.X == cpu.Y {
		cpu.PC = cpu.Read16(cpu.PC + 1)
	}
	cycle(cpu, 16)
}
func JPGTXY16(cpu *CPU) {
	if cpu.X > cpu.Y {
		cpu.PC = cpu.Read16(cpu.PC + 1)
	}
	cycle(cpu, 16)
}

func JPLTXY16(cpu *CPU) {
	if cpu.X < cpu.Y {
		cpu.PC = cpu.Read16(cpu.PC + 1)
	}
	cycle(cpu, 16)
}

// --------- A Register ------------ \\

func LDA8(cpu *CPU) {
	cpu.A = cpu.Read8(cpu.PC + 1)
	cpu.PC++
	cycle(cpu, 8)
}

func INCA(cpu *CPU) {
	cpu.A++
	cycle(cpu, 4)
}

func DECA(cpu *CPU) {
	cpu.A--
	cycle(cpu, 4)
}

func SUB8(cpu *CPU) {
	cpu.A -= cpu.Read8(cpu.PC + 1)
	cpu.PC++
	cycle(cpu, 4)
}

func LDf16A(cpu *CPU) {
	addr := cpu.Read16(cpu.PC + 1)
	cpu.PC += 2
	cpu.Write8(addr, cpu.A)
	cycle(cpu, 16)
}

// --------- B Register ------------ \\

func LDB8(cpu *CPU) {
	cpu.B = cpu.Read8(cpu.PC + 1)
	cpu.PC++
	cycle(cpu, 8)
}

func INCB(cpu *CPU) {
	cpu.B++
	cycle(cpu, 4)
}

func DECB(cpu *CPU) {
	cpu.B--
	cycle(cpu, 4)
}

func SUBB(cpu *CPU) {
	cpu.B -= cpu.A
	cycle(cpu, 4)
}

func LDf16B(cpu *CPU) {
	addr := cpu.Read16(cpu.PC + 1)
	cpu.PC += 2
	cpu.Write8(addr, cpu.B)
	cycle(cpu, 16)
}
