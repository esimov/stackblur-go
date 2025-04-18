/*
stackblur-go is a Go port of the Stackblur algorithm.

Stackblur is a compromise between Gaussian blur and Box blur, but it creates much better looking blurs than Box blur and it is ~7x faster than Gaussian blur.

The API is very simple and easy to integrate into any project. You only need to invoke the Process function which receive an image and a radius as parameters and returns the blurred version of the provided image. There is also a ProcessP function which allows to provide the destination image pointer.

	func Process(src image.Image, radius uint32) (*image.NRGBA, error)
	func ProcessP(src image.Image, dst *image.NRGBA, radius uint32) error

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

		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatalf("could not decode source file: %v", err)
		}

		src, err := stackblur.Process(img, radius)
		if err != nil {
			log.Fatal(err)
		}

		output, err := os.OpenFile("output.jpg", os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			log.Fatalf("could not open destination file: %v", err)
		}
		defer output.Close()

		if err = jpeg.Encode(output, src, &jpeg.Options{Quality: 100}); err != nil {
			log.Fatalf("could not encode destination image: %v", err)
		}
	}
*/
package stackblur
