package main

import (
	"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"path/filepath"
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
	if err != nil {
		log.Fatal("could not open source image:", err)
	}
	defer img.Close()

	src, _, err := image.Decode(img)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	if *outputGif {
		for i := 1; i <= *radius; i++ {
			img, err := stackblur.Process(src, uint32(i))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("frame %d/%d\n", i, *radius)
			go func() {
				imgs = append(imgs, img)
				if i == *radius {
					if err := generateImage(*destination, img); err != nil {
						log.Fatal(err)
					}
				}
			}()
		}
		fmt.Printf("encoding GIF\n")

		dest := path.Dir(*destination) + "/" + "output.gif"
		if err := encodeGIF(imgs, dest); err != nil {
			log.Fatal(err)
		}
	} else {
		img, err := stackblur.Process(src, uint32(*radius))
		if err != nil {
			log.Fatal(err)
		}
		if err := generateImage(*destination, img); err != nil {
			log.Fatal(err)
		}
	}
	end := time.Since(start)
	fmt.Printf("Generated in: %.2fs\n", end.Seconds())
}

// encodeGIF encodes the generated output into a gif file
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

// generateImage generates the image type depending on the provided extension
func generateImage(dst string, img image.Image) error {
	output, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer output.Close()

	if err != nil {
		return err
	}
	ext := filepath.Ext(output.Name())

	switch ext {
	case ".jpg", ".jpeg":
		if err = jpeg.Encode(output, img, &jpeg.Options{Quality: 100}); err != nil {
			return err
		}
	case ".png":
		if err = png.Encode(output, img); err != nil {
			return err
		}
	}
	return nil
}
