package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
)

var Palette = []color.Color{
	color.RGBA{0xff, 0xff, 0xff, 0},
	color.RGBA{0x44, 0xff, 0x44, 0xff},
	color.RGBA{0xff, 0x44, 0x44, 0xff},
	color.RGBA{0x44, 0x44, 0x44, 0xff},
	color.RGBA{0, 0, 0, 0xff},
}

func main() {
	dest := &gif.GIF{}
	addFrame(dest, color.RGBA{0x44, 0xff, 0x44, 0xff}, color.RGBA{0x44, 0x44, 0x44, 0xff})
	addFrame(dest, color.RGBA{0xff, 0x44, 0x44, 0xff}, color.RGBA{0, 0, 0, 0xff})
	f, err := os.Create("test.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gif.EncodeAll(f, dest)
}

func addFrame(dest *gif.GIF, fillColor, strokeColor color.Color) {
	dest1 := image.NewRGBA(image.Rect(0, 0, 200, 200))
	gc1 := draw2dimg.NewGraphicContext(dest1)
	drawShape(gc1, fillColor, strokeColor)
	destGif := image.NewPaletted(image.Rect(0, 0, 200, 200), Palette)
	draw.Draw(destGif, image.Rect(0, 0, 200, 200), dest1, image.Point{0, 0}, draw.Over)
	dest.Image = append(dest.Image, destGif)
	dest.Delay = append(dest.Delay, 100)
	dest.Disposal = append(dest.Disposal, gif.DisposalBackground)
}

func drawShape(gc *draw2dimg.GraphicContext, fillColor, strokeColor color.Color) {
	// Set some properties
	gc.SetFillColor(fillColor)
	gc.SetStrokeColor(strokeColor)
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.BeginPath()    // Initialize a new path
	gc.MoveTo(200, 0) // Move to a position to start the new path
	gc.LineTo(0, 0)
	// gc.QuadCurveTo(100, 200, 200, 0)
	gc.CubicCurveTo(-50, 200, 200, 200, 200, 0)
	gc.Close()
	gc.FillStroke()
}
