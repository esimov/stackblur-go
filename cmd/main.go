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
	"sort"
	"sync"
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
	flag.Parse()

	if len(*source) == 0 || len(*destination) == 0 {
		log.Fatal("Usage: stackblur -in input.jpg -out out.jpg")
	}

	var imgs = make([]image.Image, *radius)

	img, err := os.Open(*source)
	if err != nil {
		log.Fatalf("could not open the source file: %v", err)
	}

	defer func() {
		if err := img.Close(); err != nil {
			log.Fatalf("error closing the opened file: %v", err)
		}
	}()

	src, _, err := image.Decode(img)
	if err != nil {
		log.Fatalf("could not decode the source image: %v", err)
	}

	wg := &sync.WaitGroup{}
	start := time.Now()

	if *outputGif {
		wg.Add(*radius)

		for i := 0; i < *radius; i++ {
			go func(idx int) {
				img, err := stackblur.Process(src, uint32(idx+1))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("frame %d/%d\n", idx, *radius)
				imgs[idx] = img

				wg.Done()
			}(i)
		}
		wg.Wait()

		sort.Slice(imgs, func(i, j int) bool { return i < j })

		fmt.Printf("encoding GIF file...\n")

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
			log.Fatalf("error generating the blurred image: %v", err)
		}
	}
	end := time.Since(start)

	fmt.Printf("Generated in: %.2fs\n", end.Seconds())
}

// encodeGIF encodes the generated output into a gif file
func encodeGIF(imgs []image.Image, path string) error {
	// load static image and construct output gif file
	g := new(gif.GIF)
	for _, src := range imgs {
		dst := image.NewPaletted(src.Bounds(), palette.Plan9)
		draw.Draw(dst, src.Bounds(), src, image.Point{}, draw.Src)
		g.Image = append(g.Image, dst)
		g.Delay = append(g.Delay, 0)
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("error closing the opened file: %v", err)
		}
	}()

	return gif.EncodeAll(f, g)
}

// generateImage generates the image type depending on the provided extension
func generateImage(dst string, img image.Image) error {
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Fatalf("error closing the opened file: %v", err)
		}
	}()

	ext := filepath.Ext(out.Name())

	switch ext {
	case ".jpg", ".jpeg":
		if err = jpeg.Encode(out, img, &jpeg.Options{Quality: 100}); err != nil {
			return err
		}
	case ".png":
		if err = png.Encode(out, img); err != nil {
			return err
		}
	}

	return nil
}
