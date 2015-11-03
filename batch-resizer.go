package main
import (
	"flag"
	"os"
	"strconv"
	"time"
	"io/ioutil"
	"strings"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"fmt"
	"github.com/nfnt/resize"
)

var (
	maxSizeFlag = flag.String("m", "", "max demension X or Y")
	maxSize uint
	currentFile string
	img image.Image
	err error
)
func main() {
	flag.Parse()
	maxSize = 500
	if (*maxSizeFlag != "") {
		temp, err := strconv.ParseUint(*maxSizeFlag, 10 ,32)
		maxSize = uint(temp)
		if (err != nil) {
			panic(err)
		}
	}
	os.Mkdir("in", 0775)
	os.Mkdir("out", 0775)
	os.Mkdir("done", 0775)
	fmt.Println("Batch Image Resizer started\n")
	for {
		files, err := ioutil.ReadDir("in")
		if (err != nil) {
			panic(err)
		}
		if(len(files) == 0) {
			time.Sleep(1 * time.Second)
			continue
		}
		currentFile = files[0].Name()
		temp := strings.Split(currentFile, ".")
		fileExt := strings.ToLower(temp[len(temp) - 1])
		file, err := os.Open("in/" + currentFile)
		if(err != nil) {
			fmt.Println(err)
			done()
			continue
		}
		if (fileExt == "jpg" || fileExt == "jpeg") {
			img, err = jpeg.Decode(file)
		} else if (fileExt == "png") {
			img, err = png.Decode(file)
		} else if (fileExt == "gif") {
			img, err = gif.Decode(file)
		} else {
			fmt.Println("File " + currentFile + " has an unsupported file exension, must be <jpg|jpeg|png|gif>")
			done()
			continue
		}
		file.Close()
		var m image.Image
		m = resize.Thumbnail(maxSize, maxSize, img, resize.NearestNeighbor)
		out, err := os.Create("out/" + currentFile)
		if(err != nil) {
			fmt.Println(err)
			done()
			continue
		}
		if (fileExt == "jpg" || fileExt == "jpeg") {
			err = jpeg.Encode(out, m, nil)
		} else if (fileExt == "png") {
			err = png.Encode(out, m)
		} else if (fileExt == "gif") {
			gif.Encode(out, m, nil)
		} else {
			fmt.Println("File :" + currentFile + "has an unsupported file exension, must be <jpg|jpeg|png|gif>")
			done()
			continue
		}
		if(err != nil) {
			fmt.Println(err)
			done()
			continue
		}
		out.Close()
		done()
		fmt.Println("Resized: ", currentFile)
	}
}

func done() {
	err = os.Rename("in/" + currentFile, "done/" + currentFile)
	if(err != nil) {
		panic(err)
	}
}