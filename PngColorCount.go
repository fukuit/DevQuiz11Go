package main

import (
	"flag"
	"fmt"
	"io"
	"image/png"
	"os"
	// "strings"
)

func CountColor(_png io.Reader) int {
	img, e := png.Decode(_png)
	if e != nil {
		fmt.Fprintf(os.Stderr, e.String())
		return -1
	}
	var cnt int = 0
	pix := make(map[string]int)
	pt := img.Bounds().Max
	for x := 0; x < pt.X; x++ {
		for y := 0; y < pt.Y; y++ {
			color := img.At(x, y)
			r, g, b, a := color.RGBA()
			str := fmt.Sprintf("%d%d%d%d", r, g, b, a)
			if pix[str] == 0 {
				pix[str] = 1
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	var filename string
	flag.Parse()
	if flag.NArg() == 0{
		filename = "./go.png"
	}else{
		filename = flag.Arg(0)
	}
	png, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.String())
		return
	}
	colors := CountColor(png)
	fmt.Println(colors)
	defer png.Close()
}
