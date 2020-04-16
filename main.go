package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func Walk(searchDir string) []string {
	fileList := []string{}

	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	return fileList
}

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		return
	}

	dir1 := args[0]
	dir2 := args[1]
	out := args[2]

	flist := Walk(dir1)

	for _, fp := range flist {
		path1 := strings.Split(fp, "\\")[1:]

		if len(path1) >= 3 {
			zoom := path1[0]
			y := path1[1]
			x := path1[2]

			go MakeDirs(out + "/" + zoom + "/" + y)

			if FileExists(dir2 + "/" + zoom + "/" + y + "/" + x) {
				im1p := dir1 + "/" + zoom + "/" + y + "/" + x
				im2p := dir2 + "/" + zoom + "/" + y + "/" + x
				outp := out + "/" + zoom + "/" + y + "/" + x

				if FileExists(outp) {
					continue
				} else {
					im1, _ := os.Open(im1p)
					defer im1.Close()
					im1b, _ := png.Decode(im1)
					im2, _ := os.Open(im2p)
					defer im2.Close()
					im2b, _ := png.Decode(im2)
					offset := image.Pt(0, 0)
					// Use same size as source image has
					im1.Seek(0, 0)
					im2.Seek(0, 0)
					b := im1b.Bounds()
					m := image.NewRGBA(b)
					// Draw source
					draw.Draw(m, b, im1b, image.ZP, draw.Src)
					// Draw watermark
					draw.Draw(m, im2b.Bounds().Add(offset), im2b, image.ZP, draw.Over)
					out, _ := os.Create(outp)

					png.Encode(out, m)

					defer out.Close()
					fmt.Println(outp)
				}
			} else {
				fmt.Println(out)
				continue
			}
		}
	}
}
