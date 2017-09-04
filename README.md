# stackblur-go
`stackblur-go` is a Go port of [Stackblur](http://incubator.quasimondo.com/processing/fast_blur_deluxe.php) algorithm created by Mario Klingemann.

To quote the author this algorithm "*is a compromise between Gaussian blur and Box blur, it creates much better looking blurs than Box blur, but it is 7x faster than Gaussian blur.*" 

Comparing to the Javascript implementation the Go version is at least 50% faster (depending on the image size and blur radius), running the same image with the same bluring radius.

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
Next build the binary file.

```bash
$ go get -u github.com/esimov/stackblur-go/cmd
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
To visualize the bluring process the cli command supports the `-gif` flag, which if is set to true it will generate a gif image. For the parallel execution of the the bluring process and the gif visualization `goroutines` are used. 

<p align="center">
<img src="https://user-images.githubusercontent.com/883386/30024201-5565b990-917b-11e7-81b4-d94ffc3e11f3.gif" height=400/>
</p>

However if you wish only to generate the blured image, because of API constraints you need to create a go channel and use this as the last parameter of the `Process` method. Something like the code below:

```Go
var done chan struct{} = make(chan struct{}, *radius)
stackblur.Process(src, uint32(src.Bounds().Dx()), uint32(src.Bounds().Dy()), uint32(*radius), done)
<-done
```

| Original image | Stackblured image |
|:--:|:--:|
| <img src="https://github.com/esimov/stackblur-go/blob/master/image/sample.png" height="300"> | <img src="https://github.com/esimov/stackblur-go/blob/master/image/output.png" height="300"> |


## License

This project is under the MIT License. See the [LICENSE](https://github.com/esimov/stackblur-go/blob/master/LICENSE) file for the full license text.
