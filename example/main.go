package main

import (
	"os"
	"log"
	"time"
	"fmt"
	"flag"
	"image"
	"image/png"
	_ "image/png"
	_ "image/jpeg"
	"github.com/esimov/stackblur-go"
)

var (
	source		= flag.String("in", "", "Source")
	destination 	= flag.String("out", "", "Destination")
)

func main() {
		flag.Parse()

		img, err := os.Open(*source)
		defer img.Close()

		src, _, err := image.Decode(img)
		if err != nil {
			panic(err)
		}
		start := time.Now()
		dst := stackblur.Process(src, uint32(src.Bounds().Dx()), uint32(src.Bounds().Dy()), 20)
		end := time.Since(start)
		fmt.Printf("Generated in: %.2fs\n", end.Seconds())

		fq, err := os.Create(*destination)
		defer fq.Close()

		if err = png.Encode(fq, dst); err != nil {
			log.Fatal(err)
		}

}