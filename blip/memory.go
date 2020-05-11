package blip

type Memory []uint8

const (
	OffVRAM         = 0x00000                // VRAM
	OffBD           = OffVRAM + 0x04000      // Background Data - 16x16 with 16 in total
	OffSD           = OffBD + (16 * 16 * 16) // Sprite Data - 16x16 with 32 in total
	OffBR           = OffSD + (16 * 16 * 32) // Background References - 16*16*2: 16*16 background on screen need 3 bytes
	OffSR           = OffBR + (16 * 16 * 3)  // Sprite References - 64*4: 64 sprites on screen need 4 bytes
	OffRV           = OffSR + (64 * 4)       // Reserved - 64 bytes
	OffInput        = OffRV + 1
	OffCD    uint16 = OffRV + 64 // Code Starts.
)

const (
	InputZ     = 0x01
	InputX     = 0x02
	InputLeft  = 0x03
	InputDown  = 0x04
	InputRight = 0x05
	InputUp    = 0x06
)

// FlagSM is the Switch Memory flag. Tells CPU to read from Cart ROM or internal memory.
// When set, memory is read from ROM; but cannot be written back.
// 0x7401
const FlagSM = OffRV + 1

func NewMemory() *Memory {
	mem := make(Memory, 65536)
	return &mem
}

func (m Memory) VRAM() Memory {
	return m[:OffBD]
}

func (m Memory) Mem() []uint8 {
	return m
}

func (m Memory) BD() Memory {
	return m[OffBD:OffSD]
}

func (m Memory) Write8(addr uint16, b uint8) {
	m[addr] = b
}

func (m Memory) StoreVec8(addr uint16, b ...uint8) {
	for i, j := addr, 0; i < addr+uint16(len(b)); i, j = i+1, j+1 {
		m.Write8(i, b[j])
	}
}

func (m Memory) StorePreserveAlpha8(addr uint16, b uint8) {
	a := (b & 0b01000000) >> 6
	if a == 1 {
		m.StoreVec8(addr, b)
	}
}

func (m Memory) StoreVecPreserveAlpha8(addr uint16, b ...uint8) {
	for i, j := addr, 0; i < addr+uint16(len(b)); i, j = i+1, j+1 {
		m.StorePreserveAlpha8(i, b[j])
	}
}

func (m Memory) Read8(addr uint16) uint8 {
	return m[addr]
}

func (m Memory) Read16(addr uint16) uint16 {
	return (uint16(m.Read8(addr)) << 8) ^ uint16(m.Read8(addr+1))
}

func (m Memory) StoreBD(offset uint16, b ...uint8) {
	m.StoreVec8(OffBD+(offset*16*16), b...)
}

func (m Memory) StoreSD(offset uint16, b ...uint8) {
	m.StoreVec8(OffSD+(offset*16*16), b...)
}

func (m Memory) WriteInput(input byte) {
	m.Write8(OffInput, input)
}
