package goimg

import (
	"testing"
	"fmt"
)

var jpegfile = "testdata/1.jpg"
var pngfile = "testdata/2.png"

func TestOpenImage(t *testing.T) {
	_, err := OpenImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}
	_, err = OpenImage(pngfile)
	if err != nil {
		t.Fatal(err)
	}

	_, err = OpenImage("wrong.jpg")
	if err == nil {
		t.Fatal(err)
	}
	fmt.Println(err)
}

func TestOpenGRAYImage(t *testing.T) {
	img, err := OpenGRAYImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(img.At(20, 20))
}

func TestOpenRGBAImage(t *testing.T) {
	img, err := OpenRGBAImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(img.At(20, 20))
}

func TestOpenALPHAImage(t *testing.T) {
	img, err := OpenALPHAImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(img)
}

func TestOpenSingleImage(t *testing.T) {
	img0, err := OpenSingleImage(jpegfile, IMAGE_SINGLE_RED)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(img0.At(20, 20))

	img1, err := OpenSingleImage(jpegfile, IMAGE_SINGLE_GREEN)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(img1.At(20, 20))

	img2, err := OpenSingleImage(jpegfile, IMAGE_SINGLE_BLUE)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(img2.At(20, 20))
}