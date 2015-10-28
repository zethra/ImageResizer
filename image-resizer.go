package main

import (
	"os"
	"flag"
	"fmt"
	"strings"
	"strconv"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"github.com/nfnt/resize"
)

const (
	ERROR = iota
	MINONE
	MINTWO
	DIMTWO
)
var state = ERROR
var filePath, outputPath string
var dimension [2]int
var outputFlag = flag.String("o", "", "output file name")
var sizeFlag = flag.String("s", "", "demensions \"XxY\"")
var minSizeFlag = flag.String("m", "", "max demensions \"XxY\" or max demension X or Y")

func main() {
	var err error
	flag.Parse()
	if (len(flag.Args()) < 1) {
		fmt.Println("Input file not provided")
		os.Exit(-1)
	}
	filePath = flag.Args()[0]
	if (*sizeFlag == "") {
		fmt.Println("No size given")
		os.Exit(-1)
	}
	temp := strings.Split(filePath, ".")
	fileExt := temp[len(temp) - 1]
	if (*outputFlag != "") {
		outputPath = *outputFlag
	} else {
		outputPath = "output." + fileExt
	}
	if (*sizeFlag != "" && *minSizeFlag != "") {
		fmt.Println("both s and m flags cannot be used")
		os.Exit(-1)
	}
	if (*minSizeFlag != "") {
		dims := strings.Split(*minSizeFlag, "x")
		if (*minSizeFlag != "" && len(dims) == 1) {
			dims := strings.Split(*minSizeFlag, "x")
			dimension[0], err = strconv.Atoi(dims[0])
			if (err != nil) {
				fmt.Println("Invalid demesions given")
				os.Exit(-1)
			}
			state = MINONE
		}
	}else if (*sizeFlag != "") {
		dims := strings.Split(*sizeFlag, "x")
		if (*sizeFlag != "" && len(dims) == 2) {
			dimension[0], err = strconv.Atoi(dims[0])
			if (err != nil || dimension[0] < 0) {
				fmt.Println("Invalid demesions given")
				os.Exit(-1)
			}
			dimension[1], err = strconv.Atoi(dims[1])
			if (err != nil || dimension[1] < 0) {
				fmt.Println("Invalid demesions given")
				os.Exit(-1)
			}
			state = DIMTWO
		}
	} else {
		fmt.Println("Invalid demesions given")
		os.Exit(-1)
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	var img image.Image
	if (fileExt == "jpg" || fileExt == "jpeg") {
		img, err = jpeg.Decode(file)
	} else if (fileExt == "png") {
		img, err = png.Decode(file)
	} else if (fileExt == "gif") {
		img, err = gif.Decode(file)
	} else {
		fmt.Println("Unsupported file exension, must be <jpg|jpeg|png|gif>")
		os.Exit(-1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	file.Close()
	var m image.Image
	switch state {
	case MINONE:
		m = resize.Thumbnail(uint(dimension[0]), uint(dimension[0]), img, resize.Lanczos3)
	case DIMTWO:
		m = resize.Resize(uint(dimension[0]), uint(dimension[1]), img, resize.Lanczos3)
	case ERROR:
		fmt.Println("An error orrcured")
		os.Exit(-1)
	}
	out, err := os.Create(outputPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer out.Close()
	if (fileExt == "jpg" || fileExt == "jpeg") {
		err = jpeg.Encode(out, m, nil)
	} else if (fileExt == "png") {
		err = png.Encode(out, m)
	} else if (fileExt == "gif") {
		gif.Encode(out, m, nil)
	} else {
		fmt.Println("Unsupported file exension, must be <jpg|jpeg|png|gif>")
		os.Exit(-1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}