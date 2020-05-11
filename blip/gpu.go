package blip

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
	"image/color"
)

type GPU interface {
	Init()
	Render() error
}

type FaifaceGPU struct {
	*Memory
	drawing bool
}

func (f FaifaceGPU) Init() {
	width := 128
	height := 512

	cfg := pixelgl.WindowConfig{
		Title:       "blip",
		Bounds:      pixel.R(0, 0, float64(width), float64(height)),
		VSync:       true,
		Undecorated: false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(false)

	buffer := image.NewRGBA(image.Rect(0, 0, width, height))
	canvas := pixelgl.NewCanvas(win.Bounds())

	c := win.Bounds().Center()

	for !win.Closed() {

		var x, y int

		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		if !f.drawing {
			vram := f.Mem()
			for i := 0; i < len(vram); i++ {
				r, g, b, a := Colour(vram[i]).ExtractComponents()
				buffer.Set(x, (height-1)-y, color.RGBA{R: r, G: g, B: b, A: a})
				x++
				if x > 127 {
					x = 0
					y++
				}
			}
		}

		win.Update()

		win.Clear(color.Black)

		canvas.SetPixels(buffer.Pix)
		canvas.Draw(win, pixel.IM.Moved(c))

		// -----------------------------------

		f.Memory.WriteInput(0x00)
		if win.Pressed(pixelgl.KeyZ) {
			f.Memory.WriteInput(InputZ)
		}
		if win.Pressed(pixelgl.KeyX) {
			f.Memory.WriteInput(InputX)
		}
		if win.Pressed(pixelgl.KeyLeft) {
			f.Memory.WriteInput(InputLeft)
		}
		if win.Pressed(pixelgl.KeyDown) {
			f.Memory.WriteInput(InputDown)
		}
		if win.Pressed(pixelgl.KeyRight) {
			f.Memory.WriteInput(InputRight)
		}
		if win.Pressed(pixelgl.KeyUp) {
			f.Memory.WriteInput(InputUp)
		}
	}
}

func (f FaifaceGPU) Render() error {
	copy(f.VRAM(), make([]uint8, len(f.VRAM())))
	f.drawing = true
	for i := OffBR; i < OffSR; i += 3 {
		m := uint16(i)
		b0 := f.Read8(m)
		b1 := f.Read8(m + 1)
		b2 := f.Read8(m + 2)
		if b0 > 0 {
			f.DrawBackground(b1, b2)
		}
	}
	for i := OffSR; i < OffRV; i += 4 {
		m := uint16(i)
		b0 := f.Read8(m)
		b1 := f.Read8(m + 1)
		b2 := f.Read8(m + 2)
		b3 := f.Read8(m + 3)
		if b0 > 0 {
			f.DrawSprite(b1, b2, b3)
		}
	}
	f.drawing = false
	return nil
}
