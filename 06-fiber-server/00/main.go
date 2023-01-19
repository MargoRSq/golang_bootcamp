package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "LaBelleAurore-Regular.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
)

var c = freetype.NewContext()
var fg, bg = image.Black, image.White
var ruler = color.RGBA{125, 235, 202, 0xff}
var rgba = image.NewRGBA(image.Rect(0, 0, 300, 300))

func setters() {
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}
}

func drawLines() {
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}
}

func drawText(x int, y int, text string, size float64) {
	c.SetFontSize(size)
	pt := freetype.Pt(x, y+int(c.PointToFixed(size)>>6))
	_, err := c.DrawString(text, pt)
	if err != nil {
		log.Println(err)
	return
	}
	pt.Y += c.PointToFixed(size * 1.5)
}

func main() {
	flag.Parse()

	setters()
	drawLines()
	drawText(90, 50, "Sq", 120.0)
	drawText(230, 263, "21 xx", 25)

	outFile, err := os.Create("amazing_logo.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("amazing_logo.png created")
}
