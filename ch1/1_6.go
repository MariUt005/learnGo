// Генерирует анимированный GIF из случайных фигур Лиссажу
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{R: 0xFF, A: 0xFF}, color.RGBA{G: 0xFF, A: 0xFF}, color.RGBA{B: 0xFF, A: 0xFF}}

const (
	whiteIndex = 0
	blackIndex = 1
	redIndex   = 2
	greenIndex = 3
	blueIndex  = 4
)

func main() {
	var clrID uint8
	if len(os.Args) > 1 && os.Args[1] == "r" {
		clrID = redIndex
	} else if len(os.Args) > 1 && os.Args[1] == "g" {
		clrID = greenIndex
	} else if len(os.Args) > 1 && os.Args[1] == "b" {
		clrID = blueIndex
	} else {
		clrID = blackIndex
	}
	lissajous(os.Stdout, clrID)
}

func lissajous(out io.Writer, clrID uint8) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 500
		nframes = 100
		delay   = 5
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), clrID)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
