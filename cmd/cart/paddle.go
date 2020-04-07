package cart

import "github.com/hscells/rougebit/blip"

var TilePink = []uint8{
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
	0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011, 0b00110011,
}

var TileGreen = []uint8{
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
	0b00001100, 0b00001100, 0b00001100, 0b00001100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100, 0b00000100,
}

var Paddle = []uint8{
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000,
	0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111,
	0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111,
	0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111,
	0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
}

var Ball = []uint8{
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000,
	0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000,
	0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000,
	0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111,
	0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111,
	0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111,
	0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111,
	0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000,
	0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000,
	0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b01111111, 0b01111111, 0b01111111, 0b01111111, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
}

func PaddleCart() *blip.Memory {
	m := blip.NewMemory()
	m.StoreBD(0, TilePink...)
	m.StoreBD(1, TileGreen...)
	m.StoreSD(0, Paddle...)
	m.StoreSD(1, Ball...)
	// Background Data.
	m.StoreVec8(blip.OffBR,
		0x01, 0x00, 0x00,
		0x01, 0x01, 0x01,
		0x01, 0x02, 0x00,
		0x01, 0x03, 0x01,
		0x01, 0x04, 0x00,
		0x01, 0x05, 0x01,
		0x01, 0x06, 0x00,
		0x01, 0x07, 0x01,
		//
		0x01, 0x08, 0x01,
		0x01, 0x09, 0x00,
		0x01, 0x0a, 0x01,
		0x01, 0x0b, 0x00,
		0x01, 0x0c, 0x01,
		0x01, 0x0d, 0x00,
		0x01, 0x0e, 0x01,
		0x01, 0x0f, 0x00,

		0x01, 0x10, 0x00,
		0x01, 0x11, 0x01,
		0x01, 0x12, 0x00,
		0x01, 0x13, 0x01,
		0x01, 0x14, 0x00,
		0x01, 0x15, 0x01,
		0x01, 0x16, 0x00,
		0x01, 0x17, 0x01,
		//
		0x01, 0x18, 0x01,
		0x01, 0x19, 0x00,
		0x01, 0x1a, 0x01,
		0x01, 0x1b, 0x00,
		0x01, 0x1c, 0x01,
		0x01, 0x1d, 0x00,
		0x01, 0x1e, 0x01,
		0x01, 0x1f, 0x00,

		0x01, 0x20, 0x00,
		0x01, 0x21, 0x01,
		0x01, 0x22, 0x00,
		0x01, 0x23, 0x01,
		0x01, 0x24, 0x00,
		0x01, 0x25, 0x01,
		0x01, 0x26, 0x00,
		0x01, 0x27, 0x01,

		0x01, 0x28, 0x01,
		0x01, 0x29, 0x00,
		0x01, 0x2a, 0x01,
		0x01, 0x2b, 0x00,
		0x01, 0x2c, 0x01,
		0x01, 0x2d, 0x00,
		0x01, 0x2e, 0x01,
		0x01, 0x2f, 0x00,

		0x01, 0x30, 0x00,
		0x01, 0x31, 0x01,
		0x01, 0x32, 0x00,
		0x01, 0x33, 0x01,
		0x01, 0x34, 0x00,
		0x01, 0x35, 0x01,
		0x01, 0x36, 0x00,
		0x01, 0x37, 0x01,

		0x01, 0x38, 0x01,
		0x01, 0x39, 0x00,
		0x01, 0x3a, 0x01,
		0x01, 0x3b, 0x00,
		0x01, 0x3c, 0x01,
		0x01, 0x3d, 0x00,
		0x01, 0x3e, 0x01,
		0x01, 0x3f, 0x00,
	)
	m.StoreVec8(blip.OffSR, 1, 25, 50, 0)
	m.StoreVec8(blip.OffSR+4, 1, 64, 0, 1)
	m.StoreVec8(blip.OffCD,
		0x00,
		0x11, // INCA
		0x22, // INCB
		0x14, 0x73, 0x01, // LDf16A 0x7301
		0x24, 0x73, 0x05, // LDf16B 0x7301
		0x01, 0x74, 0x40, // JP16 0x7440
	)

	return m
}