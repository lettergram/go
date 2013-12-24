/**
  * Created by Austin Walters
  * Image Synthesis in Go
  * 12/24/2013
  *
  * Exersise taken From: 
  * A Tour of Go: Exercise: Images
  * Slide #60
  *
  **/


package main
 
import (
	"image"
	"image/color"
	"image/png"
	"os"
)
 
type Image struct{
	Width, Height int
	colur uint8
}
 
func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}
 
func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}
 
func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.colur + uint8(x/6), r.colur + uint8(y/4), 100, 255}
}

func main() {
	m := Image{1530, 1020, 0}
	w, err := os.Create("../../Pictures/output.png")
	if err != nil { panic(err) }
	png.Encode(w, &m)
}
