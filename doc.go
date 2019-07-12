/*
stackblur-go is a Go port of the Stackblur algorithm.

Stackblur is a compromise between Gaussian blur and Box blur, but it creates much better looking blurs than Box blur and it is ~7x faster than Gaussian blur.

The API is very simple and easy to integrate into the project. There is a single publicly exposed `Process` functions which receive an image and a radius as parameters and returns the blurred version of the provided image.

	func Process(src image.Image, radius uint32) image.Image

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
		img, err := os.Open("sample.jpg")
		defer img.Close()

		src, _, err := image.Decode(img)
		if err != nil {
			log.Fatal(err)
		}

		res := stackblur.Process(src, uint32(5))

		output, err := os.OpenFile("output.jpg", os.O_CREATE|os.O_RDWR, 0755)
		defer output.Close()
		if err != nil {
			log.Fatal(err)
		}

		if err = jpeg.Encode(output, res, &jpeg.Options{Quality: 100}); err != nil {
			log.Fatal(err)
		}
	}
 */
package stackblur