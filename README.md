# stackblur-go

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/esimov/stackblur-go)
[![Build Status](https://travis-ci.org/esimov/stackblur-go.svg?branch=master)](https://travis-ci.org/esimov/stackblur-go)

Go port of Mario Klingemann's [Stackblur](http://incubator.quasimondo.com/processing/fast_blur_deluxe.php) algorithm.

Stackblur is a compromise between Gaussian blur and Box blur, but it creates much better looking blurs than Box blur and it is ~7x faster than Gaussian blur.

Comparing to the Javascript implementation the Go version is at least 50% faster (depending on the image size and blur radius), applied on the same image with the same bluring radius.

### Benchmark
Radius       | Javascript  | Go
-------------|-------------|-------------
20           | ~15ms       | ~7.4ms

## Installation

First, install Go, set your GOPATH, and make sure $GOPATH/bin is on your PATH.

```bash
$ export GOPATH="$HOME/go"
$ export PATH="$PATH:$GOPATH/bin"
```

Next download the project and build the binary file.

```bash
$ go get -u -f github.com/esimov/stackblur-go
$ cd cmd && go build -o $GOPATH/bin/stackblur
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

## API call

The API is very simple: you have to expose an image file and a blur radius to the `Process` function.

```Go
stackblur.Process(src, blurRadius)
```

## Results

| Original image | Blured image |
|:--:|:--:|
| <img src="https://github.com/esimov/stackblur-go/blob/master/image/sample.png" height="300"> | <img src="https://github.com/esimov/stackblur-go/blob/master/image/output.png" height="300"> |


## License

This project is under the MIT License. See the [LICENSE](https://github.com/esimov/stackblur-go/blob/master/LICENSE) file for the full license text.
