package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"

	"github.com/sunshineplan/imgconv"
)

type outputDefinition struct {
	width  int // Width in pixels
	height int // Height in pixels
	numCol int // number of columns for stickers
	numRow int // number of rows for stickers
}

const DPI = 150

var A4sticker = outputDefinition{
	width:  int(math.Round(8.27 * DPI)),
	height: int(math.Round(11.69 * DPI)),
	numCol: 2,
	numRow: 2,
}

func getImage(fileName string) (image.Image, error) {

	currDir, _ := os.Getwd()
	src, err := imgconv.Open(filepath.Join(currDir, fileName))

	if err != nil {
		return nil, err
	}

	return src, nil
}

func placeOnPaper(img image.Image, def outputDefinition, x int, y int) image.Image {

	// Sticker margin in pixels
	margin := 10
	width := def.width
	height := def.height

	// Create new paper image
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	newImg := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	xStart := int(math.Round((float64(x)/float64(def.numCol))*float64(width))) + margin
	xEnd := int(math.Round((float64(x+1)/float64(def.numCol))*float64(width))) - margin

	yStart := int(math.Round((float64(y)/float64(def.numRow))*float64(height))) + margin
	yEnd := int(math.Round((float64(y+1)/float64(def.numRow))*float64(height))) - margin

	draw.NearestNeighbor.Scale(newImg, image.Rect(xStart, yStart, xEnd, yEnd), img, img.Bounds(), draw.Over, nil)

	return newImg
}

func main() {

	fmt.Println(os.Args)

	// Extract Arguments from CLI
	var fileName = flag.String("fileName", "./label.pdf", "Link to the filename")
	var output = flag.String("output", "./output.pdf", "Name of the resulting file")
	var colNum = flag.Int("col", 0, "Column Number")
	var rowNum = flag.Int("row", 0, "Row Number")

	flag.Parse()

	log.Println("Test Function ", *fileName)
	src, err := getImage(*fileName)

	if err != nil {
		log.Fatalf("Failed to read pdf %v", err)
	}

	img := placeOnPaper(src, A4sticker, *colNum, *rowNum)
	// err = imgconv.Write(io.Discard, src, &imgconv.FormatOption{Format: imgconv.PNG})
	currDir, _ := os.Getwd()
	err = imgconv.Save(filepath.Join(currDir, *output), img, &imgconv.FormatOption{Format: imgconv.PDF})

	if err != nil {
		log.Fatalf("Failed to write PDF %v", err)
	}

	log.Println("Opened PDF 2")

}
