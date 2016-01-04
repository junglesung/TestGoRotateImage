package main

import (
	"fmt"
	"os"
	"log"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/disintegration/imaging"
	"image"
)

func main() {
	fmt.Println("Hello, world")
	fname := "P1070332.JPG"
	rname := "P1070332_rotate.JPG"
	sname := "P1070332_small.JPG"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// Get Orientation
	orientation, err := x.Get(exif.Orientation)
	if err != nil {
		fmt.Println(exif.Model, " not fround", )
		return
	}
	fmt.Println("Orientation", orientation.String())

	// Rotate
	var rotateImage *image.NRGBA
	openImage, err := imaging.Open(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch orientation.String() {
	case "1":
	// Do nothing
	case "2":
		rotateImage = imaging.FlipH(openImage)
	case "3":
		rotateImage = imaging.Rotate180(openImage)
	case "4":
		rotateImage = imaging.FlipV(openImage)
	case "5":
		rotateImage = imaging.Transverse(openImage)
	case "6":
		rotateImage = imaging.Rotate270(openImage)
	case "7":
		rotateImage = imaging.Transpose(openImage)
	case "8":
		rotateImage = imaging.Rotate90(openImage)
	}
	err = imaging.Save(rotateImage, rname)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rname, " saved")

	// Small
	var smallImage *image.NRGBA
	if rotateImage.Rect.Dx() > rotateImage.Rect.Dy() {
		smallImage = imaging.Resize(rotateImage, 1920, 0, imaging.Lanczos)
	} else {
		smallImage = imaging.Resize(rotateImage, 0, 1920, imaging.Lanczos)
	}
	err = imaging.Save(smallImage, sname)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sname, " saved")

	// Use jpeg.Encode() to write to a file
	// https://github.com/disintegration/imaging/blob/master/helpers.go#L79
	// func Encode(w io.Writer, m image.Image, o *Options) error
	// https://golang.org/pkg/image/jpeg/
}
