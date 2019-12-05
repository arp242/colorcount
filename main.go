package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"sort"
)

func main() {
	fp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(fp)
	fp.Close()
	if err != nil {
		log.Fatal(err)
	}

	colors := make(map[string][]uint32)
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			v := []uint32{r >> 8, g >> 8, b >> 8, a >> 8}
			k := fmt.Sprintf("#%.2x%.2x%.2x", v[0], v[1], v[2])
			if v[3] < 255 {
				k += fmt.Sprintf("%.2x", v[3])
			}
			colors[k] = v
		}
	}

	var scolors []string
	for c := range colors {
		scolors = append(scolors, c)
	}
	sort.Strings(scolors)

	fmt.Printf("%d colors:\n", len(colors))
	w := -1
	for _, c := range scolors {
		w += len(c) + 2
		if w > 80 {
			fmt.Println("")
			w = 0
		}

		fmt.Printf("\x1b[48;2;%d;%d;%dm \x1b[0m%s ",
			colors[c][0], colors[c][1], colors[c][2], c)
	}
	fmt.Println("")
}
