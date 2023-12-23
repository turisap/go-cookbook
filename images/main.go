package main

import (
	"crypto/rand"
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadImg() *image.Paletted {
	f, err := os.Open("gh.png")

	checkError(err)

	fmt.Println("Opened file")
	defer f.Close()

	img, fn, err := image.Decode(f)

	checkError(err)

	fmt.Println("Decoded", fn)
	fmt.Printf("%T\n", img)

	rimg, ok := img.(*image.Paletted)

	fmt.Println(ok)
	_ = rimg

	return rimg
}

func save(filePath string, img *image.Paletted) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	png.Encode(imgFile, img.SubImage(img.Rect))

}

func createImg() {
	rect := image.Rect(0, 0, 100, 100)
	img := createRandomImage(rect)
	save("random.png", img)
}

func createRandomImage(rect image.Rectangle) (created *image.Paletted) {
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	rand.Read(pix)
	created = &image.Paletted{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	return
}

func main() {
	//img := loadImg()
	//save("new-gh.png", img)

	createImg()
}
