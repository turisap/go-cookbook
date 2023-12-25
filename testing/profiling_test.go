package kirmath

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func resize(grid [][]color.Color, scale float64) (resized [][]color.Color) {
	xlen, ylen := int(float64(len(grid))*scale), int(float64(len(grid[0]))*
		scale)
	resized = make([][]color.Color, xlen)
	for i := 0; i < len(resized); i++ {
		resized[i] = make([]color.Color, ylen)
	}
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			xp := int(math.Floor(float64(x) / scale))
			yp := int(math.Floor(float64(y) / scale))
			resized[x][y] = grid[xp][yp]
		}
	}
	return
}

// load the image from file
func load(filePath string) (grid [][]color.Color) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}
	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color

		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		grid = append(grid, y)
	}
	return
}

// save the image to file
func save(filePath string, grid [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()
	png.Encode(file, img.SubImage(img.Rect))
}
