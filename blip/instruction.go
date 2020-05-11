package blip

type Instruction func(cpu *CPU)

var Instructions = map[uint8]Instruction{
	0x00: NOP, 0x01: JP16, 0x02: JPEQ, 0x03: JPLT, 0x04: JPGT,
	0x10: MOV, 0x11: LDA, 0x12: INC, 0x13: DEC, 0x14: RDA,
	0x20: ADD, 0x21: SUB, 0x22: MUL, 0x23: DIV,
	0x30: AND, 0x31: OR, 0x32: XOR, 0x33: SLA, 0x34: SRA,
}

func NOP(cpu *CPU) {
	cycle(cpu, 4)
}

func JP16(cpu *CPU) {
	cpu.PC = cpu.Read16(cpu.PC + 1)
	cpu.PC--
	cycle(cpu, 16)
}

func JPEQ(cpu *CPU) {
	if cpu.A == cpu.Read8(cpu.PC+1) {
		cpu.PC = cpu.Read16(cpu.PC + 2)
		cpu.PC--
	} else {
		cpu.PC += 3
	}
	cycle(cpu, 16)
}

func JPLT(cpu *CPU) {
	if cpu.A < cpu.Read8(cpu.PC+1) {
		cpu.PC = cpu.Read16(cpu.PC + 2)
		cpu.PC--
	} else {
		cpu.PC += 3
	}
	cycle(cpu, 16)
}

func JPGT(cpu *CPU) {
	if cpu.A > cpu.Read8(cpu.PC+1) {
		cpu.PC = cpu.Read16(cpu.PC + 2)
		cpu.PC--
	} else {
		cpu.PC += 3
	}
	cycle(cpu, 16)
}

func MOV(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+1), cpu.Read8(cpu.PC+3))
	cpu.PC += 3
	cycle(cpu, 4)
}

func INC(cpu *CPU) {
	addr := cpu.Read16(cpu.PC + 1)
	val := cpu.Read8(addr) + 1
	cpu.Write8(addr, val)
	cpu.PC += 2
	cycle(cpu, 4)
}

func DEC(cpu *CPU) {
	addr := cpu.Read16(cpu.PC + 1)
	val := cpu.Read8(addr) - 1
	cpu.Write8(addr, val)
	cpu.PC += 2
	cycle(cpu, 4)
}

func LDA(cpu *CPU) {
	cpu.A = cpu.Read8(cpu.Read16(cpu.PC + 1))
	cpu.PC += 2
	cycle(cpu, 4)
}

func RDA(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+1), cpu.A)
	cpu.PC += 2
	cycle(cpu, 4)
}

func ADD(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+5), cpu.Read8(cpu.Read16(cpu.PC+1))+cpu.Read8(cpu.Read16(cpu.PC+3)))
	cpu.PC += 6
	cycle(cpu, 12)
}

func SUB(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+5), cpu.Read8(cpu.Read16(cpu.PC+1))-cpu.Read8(cpu.Read16(cpu.PC+3)))
	cpu.PC += 6
	cycle(cpu, 4)
}

func MUL(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+3), cpu.Read8(cpu.PC+1)*cpu.Read8(cpu.PC+2))
	cycle(cpu, 4)
}

func DIV(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+3), cpu.Read8(cpu.PC+1)/cpu.Read8(cpu.PC+2))
	cycle(cpu, 4)
}

func AND(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+3), cpu.Read8(cpu.PC+1)&cpu.Read8(cpu.PC+2))
}

func OR(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+3), cpu.Read8(cpu.PC+1)|cpu.Read8(cpu.PC+2))
}

func XOR(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+3), cpu.Read8(cpu.PC+1)^cpu.Read8(cpu.PC+2))
}

func SLA(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+3), cpu.Read8(cpu.PC+1)<<cpu.Read8(cpu.PC+2))
}

func SRA(cpu *CPU) {
	cpu.Write8(cpu.Read16(cpu.PC+3), cpu.Read8(cpu.PC+1)>>cpu.Read8(cpu.PC+2))
}
