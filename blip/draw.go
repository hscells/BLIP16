package blip

func (m Memory) DrawBackground(positionOffset, dataOffset uint8) {
	data := m[OffBD+(uint16(dataOffset)*256) : OffBD+(uint16(dataOffset)*256)+255]
	y := uint16(positionOffset-(positionOffset%8)) * 15
	off := (uint16(positionOffset) + y) * 16
	m.StoreVec8(off+128*0, data[:0x10]...)
	m.StoreVec8(off+128*1, data[0x10:0x20]...)
	m.StoreVec8(off+128*2, data[0x20:0x30]...)
	m.StoreVec8(off+128*3, data[0x30:0x40]...)
	m.StoreVec8(off+128*4, data[0x40:0x50]...)
	m.StoreVec8(off+128*5, data[0x50:0x60]...)
	m.StoreVec8(off+128*6, data[0x60:0x70]...)
	m.StoreVec8(off+128*7, data[0x80:0x90]...)
	m.StoreVec8(off+128*8, data[0x80:0x90]...)
	m.StoreVec8(off+128*9, data[0x90:0xa0]...)
	m.StoreVec8(off+128*10, data[0xa0:0xb0]...)
	m.StoreVec8(off+128*11, data[0xb0:0xc0]...)
	m.StoreVec8(off+128*12, data[0xc0:0xd0]...)
	m.StoreVec8(off+128*13, data[0xd0:0xe0]...)
	m.StoreVec8(off+128*14, data[0xe0:0xf0]...)
	m.StoreVec8(off+128*15, data[0xf0:0x100]...)
}

func (m Memory) DrawSprite(x, y, dataOffset uint8) {
	data := m[OffSD+(uint16(dataOffset)*256) : OffSD+(uint16(dataOffset)*256)+255]
	off := (uint16(y) * 128) + uint16(x)
	m.StoreVecPreserveAlpha8(off+128*0, data[:0x10]...)
	m.StoreVecPreserveAlpha8(off+128*1, data[0x10:0x20]...)
	m.StoreVecPreserveAlpha8(off+128*2, data[0x20:0x30]...)
	m.StoreVecPreserveAlpha8(off+128*3, data[0x30:0x40]...)
	m.StoreVecPreserveAlpha8(off+128*4, data[0x40:0x50]...)
	m.StoreVecPreserveAlpha8(off+128*5, data[0x50:0x60]...)
	m.StoreVecPreserveAlpha8(off+128*6, data[0x60:0x70]...)
	m.StoreVecPreserveAlpha8(off+128*7, data[0x80:0x90]...)
	m.StoreVecPreserveAlpha8(off+128*8, data[0x80:0x90]...)
	m.StoreVecPreserveAlpha8(off+128*9, data[0x90:0xa0]...)
	m.StoreVecPreserveAlpha8(off+128*10, data[0xa0:0xb0]...)
	m.StoreVecPreserveAlpha8(off+128*11, data[0xb0:0xc0]...)
	m.StoreVecPreserveAlpha8(off+128*12, data[0xc0:0xd0]...)
	m.StoreVecPreserveAlpha8(off+128*13, data[0xd0:0xe0]...)
	m.StoreVecPreserveAlpha8(off+128*14, data[0xe0:0xf0]...)
	m.StoreVecPreserveAlpha8(off+128*15, data[0xf0:0x100]...)
}
