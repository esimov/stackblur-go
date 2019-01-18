package main

import (
	"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/esimov/stackblur-go"
)

var (
	source      = flag.String("in", "", "Source")
	destination = flag.String("out", "", "Destination")
	radius      = flag.Int("radius", 20, "Radius")
	outputGif   = flag.Bool("gif", false, "Output Gif")
)

func main() {
	var imgs []image.Image

	flag.Parse()

	if len(*source) == 0 || len(*destination) == 0 {
		log.Fatal("Usage: stackblur -in input.jpg -out out.jpg")
	}

	img, err := os.Open(*source)
	defer img.Close()

	src, _, err := image.Decode(img)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	if *outputGif {
		for i := 1; i <= *radius; i++ {
			img := stackblur.Process(src, uint32(i))
			fmt.Printf("frame %d/%d\n", i, *radius)
			go func() {
				imgs = append(imgs, img)
				if i == *radius {
					generateImage(*destination, img)
				}
			}()
		}
		fmt.Printf("encoding GIF\n")
		if err := encodeGIF(imgs, "output.gif"); err != nil {
			log.Fatal(err)
		}
	} else {
		img := stackblur.Process(src, uint32(*radius))
		generateImage(*destination, img)
	}
	end := time.Since(start)
	fmt.Printf("\nGenerated in: %.2fs\n", end.Seconds())
}

// Visualize the bluring by outputting the generated image into a gif file
func encodeGIF(imgs []image.Image, path string) error {
	// load static image and construct outGif
	outGif := &gif.GIF{}
	for _, inPng := range imgs {
		inGif := image.NewPaletted(inPng.Bounds(), palette.Plan9)
		draw.Draw(inGif, inPng.Bounds(), inPng, image.Point{}, draw.Src)
		outGif.Image = append(outGif.Image, inGif)
		outGif.Delay = append(outGif.Delay, 0)
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return gif.EncodeAll(f, outGif)
}

func generateImage(dst string, img image.Image) {
	fq, err := os.Create(dst)
	defer fq.Close()

	if err = png.Encode(fq, img); err != nil {
		log.Fatal(err)
	}
}
