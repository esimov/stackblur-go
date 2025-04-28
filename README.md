# stackblur-go

[![Go Reference](https://pkg.go.dev/badge/github.com/esimov/stackblur-go.svg)](https://pkg.go.dev/github.com/esimov/stackblur-go)
[![build](https://github.com/esimov/stackblur-go/actions/workflows/build.yml/badge.svg)](https://github.com/esimov/stackblur-go/actions/workflows/build.yml)

Go port of Mario Klingemann's [Stackblur](http://incubator.quasimondo.com/processing/fast_blur_deluxe.php) algorithm.

Stackblur is a compromise between Gaussian blur and Box blur, but it creates much better looking blurs than Box blur and it is ~7x faster than Gaussian blur.

Comparing to the Javascript implementation the Go version is at least 50% faster (depending on the image size and blur radius), applied on the same image with the same bluring radius.

### Benchmark
Radius       | Javascript  | Go
-------------|-------------|-------------
20           | ~15ms       | ~7.4ms

## Installation

```bash
$ go install github.com/esimov/stackblur-go/cmd/stackblur@latest
```

#### CLI example

The provided CLI example supports the following flags:
```bash
$ stackblur --help

Usage of stackblur:
  -gif
    	Output Gif
  -in string
    	Source
  -out string
    	Destination
  -radius int
    	Radius (default 20)
```
The command below will generate the blurred version of the source image.

```bash
$ stackblur -in image/sample.png -out image/output.png -radius 10
```
The cli command supports a `-gif` flag, which if set as true it visualize the bluring process by outputting the result into a gif file.

## API

The usage of the API is very simple: it exposes a single public `Process` function which requires a destination and a source image together with a blur radius. The blured image will be encoded into the destination image.

```Go
stackblur.Process(dst, src, blurRadius)
```

## Results

| Original image | Blurred image |
|:--:|:--:|
| <img src="https://github.com/esimov/stackblur-go/blob/master/image/sample.png" height="300"> | <img src="https://github.com/esimov/stackblur-go/blob/master/image/output.png" height="300"> |


## License

This project is under the MIT License. See the [LICENSE](https://github.com/esimov/stackblur-go/blob/master/LICENSE) file for the full license text.
