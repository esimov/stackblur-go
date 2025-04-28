/*
stackblur-go is a Go port of the Stackblur algorithm.

Stackblur is a compromise between Gaussian blur and Box blur, but it creates much better looking blurs than Box blur and it is ~7x faster than Gaussian blur.

The usage of the API is very simple: it exposes a single public `Process` function which requires a destination and a source image together with a blur radius. The blured image will be encoded into the destination image.

	func Process(dst, src image.Image, radius uint32) error

Below is a very simple example of how you can use this package.

	package main

	import (
		"image"
		"image/jpeg"
		"log"
		"os"

		"github.com/esimov/stackblur-go"
	)

	func main() {
		var radius uint32 = 5
		f, err := os.Open("sample.png")
		if err != nil {
			log.Fatalf("could not open source file: %v", err)
		}
		defer f.Close()

		src, _, err := image.Decode(f)
		if err != nil {
			log.Fatalf("could not decode source file: %v", err)
		}

		dst := image.NewNRGBA(src.Bounds())
		err = stackblur.Process(dst, src, radius)
		if err != nil {
			log.Fatal(err)
		}

		output, err := os.OpenFile("output.jpg", os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			log.Fatalf("could not open destination file: %v", err)
		}
		defer output.Close()

		if err = jpeg.Encode(output, dst, &jpeg.Options{Quality: 100}); err != nil {
			log.Fatalf("could not encode destination image: %v", err)
		}
	}
*/
package stackblur
