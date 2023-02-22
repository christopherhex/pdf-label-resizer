package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"golang.org/x/image/draw"

	"github.com/sunshineplan/imgconv"
)

type outputDefinition struct {
	width int // Width in pixels
	height int // Height in pixels
	numCol int // number of columns for stickers
	numRow int // number of rows for stickers
}

const DPI = 150

// const A4sticker = outputDefinition{
// 	width: math.Round(8.27 * DPI),
// 	height: math.Round(11.69 * DPI),
// 	numCol: 2,
// 	numRow: 2,
// }


func getImage(fileName string) (image.Image, error) {
	src, err := imgconv.Open(fileName)

	if err != nil {
		return nil, err
	}

	return src, nil
}


func placeOnPaper(img image.Image) (image.Image) {

	width := int(math.Round(8.27 * DPI))
	height := int(math.Round(11.69 * DPI))

	// Create new paper image
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	newImg := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	draw.Draw(newImg, newImg.Bounds(),&image.Uniform{color.White},image.Point{},draw.Src)

	draw.NearestNeighbor.Scale(newImg,image.Rect(0,0,width/2, height/2),img,img.Bounds(),draw.Over,nil)


	// isPortrait := img.Bounds().Max.X < img.Bounds().Max.Y

	// log.Println("Image is portrait %v", isPortrait)

	return newImg
}



func main() {

	log.Println("Test Function")
	src, _ := getImage("verzendlabel-test.pdf")

	img := placeOnPaper(src)
	// err = imgconv.Write(io.Discard, src, &imgconv.FormatOption{Format: imgconv.PNG})
	err := imgconv.Save("test.pdf",img, &imgconv.FormatOption{Format: imgconv.PDF})

	if err != nil {
		log.Fatalf("Failed to write PDF %v", err)
	}

	log.Println("Opened PDF 2")

}