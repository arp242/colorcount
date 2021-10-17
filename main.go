package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"sort"
)

func main() {
	displayColors := flag.Bool("c", false, "display all the colours, not just the count")
	show := flag.Bool("s", false, "show the image")
	flag.Parse()

	if *show {
		*displayColors = true
	}

	nums := "0123456789" +
		"abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	for _, f := range flag.Args() {
		fp, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}

		img, _, err := image.Decode(fp)
		fp.Close()
		if err != nil {
			log.Fatal(err)
		}

		type color struct {
			char  rune
			color []uint32
		}

		colors := make(map[string]color)
		bounds := img.Bounds()
		i := 0
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := img.At(x, y).RGBA()
				v := []uint32{r >> 8, g >> 8, b >> 8, a >> 8}
				k := fmt.Sprintf("#%.2x%.2x%.2x", v[0], v[1], v[2])
				if v[3] < 255 {
					k += fmt.Sprintf("%.2x", v[3])
				}
				if _, ok := colors[k]; !ok {
					colors[k] = color{rune(nums[i]), v}
					i++
				}
			}
		}

		fmt.Printf("%d colors in %s\n", len(colors), f)

		if *displayColors {
			var scolors []string
			for c := range colors {
				scolors = append(scolors, c)
			}
			sort.Strings(scolors)

			w := -1
			for _, c := range scolors {
				w += len(c) + 2
				if w > 80 {
					fmt.Println("")
					w = 0
				}

				rt, gt, bt := lighten(colors[c].color[0], colors[c].color[1], colors[c].color[2])
				fmt.Printf("\x1b[48;2;%d;%d;%dm\x1b[97m\x1b[38;2;%d;%d;%dm%s\x1b[0m %s ",
					colors[c].color[0], colors[c].color[1], colors[c].color[2],
					rt, gt, bt,
					string(colors[c].char), c)
			}
			fmt.Println("")
		}

		if *show {
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					r, g, b, a := img.At(x, y).RGBA()
					v := []uint32{r >> 8, g >> 8, b >> 8, a >> 8}
					rt, gt, bt := lighten(v[0], v[1], v[2])
					k := fmt.Sprintf("#%.2x%.2x%.2x", v[0], v[1], v[2])

					fmt.Printf("\x1b[48;2;%d;%d;%dm\x1b[38;2;%d;%d;%dm%s\x1b[0m",
						v[0], v[1], v[2],
						rt, gt, bt,
						string(colors[k].char))
				}
				fmt.Println()
			}
		}
	}
}

func lighten(r, g, b uint32) (uint32, uint32, uint32) {
	return r + 80, g + 80, b + 80
}
