package goimg

import (
	"testing"
	"fmt"
	"image"
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

func TestSaveImage(t *testing.T) {
	img, err := OpenGRAYImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	err = SaveImage("/tmp/g.jpg", img, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}

	img1, err := OpenSingleImage(pngfile, IMAGE_SINGLE_GREEN)
	if err != nil {
		t.Fatal(err)
	}

	err = SaveImage("/tmp/g.png", img1, SAVE_PNG)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveImage("/tmp/gg.jpg", img1, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRotateImage(t *testing.T) {
	src, err := OpenRGBAImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	left := image.NewRGBA(image.Rect(0,0,src.Bounds().Max.Y,src.Bounds().Max.X))
	right := image.NewRGBA(image.Rect(0,0,src.Bounds().Max.Y,src.Bounds().Max.X))
	full := image.NewRGBA(image.Rect(0,0,src.Bounds().Max.X,src.Bounds().Max.Y))

	err = RotateImage(left, src, ROTATE_LEFT)
	if err != nil {
		t.Fatal(err)
	}
	err = RotateImage(right, src, ROTATE_RIGHT)
	if err != nil {
		t.Fatal(err)
	}
	err = RotateImage(full, src, ROTATE_FULL)
	if err != nil {
		t.Fatal(err)
	}

	err = SaveImage("/tmp/left.jpg", left, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveImage("/tmp/right.jpg", right, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveImage("/tmp/full.jpg", full, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRotateImage2(t *testing.T) {
	src, err := OpenGRAYImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	left := image.NewGray(image.Rect(0,0,src.Bounds().Max.Y,src.Bounds().Max.X))
	right := image.NewGray(image.Rect(0,0,src.Bounds().Max.Y,src.Bounds().Max.X))
	full := image.NewGray(image.Rect(0,0,src.Bounds().Max.X,src.Bounds().Max.Y))

	err = RotateImage(left, src, ROTATE_LEFT)
	if err != nil {
		t.Fatal(err)
	}
	err = RotateImage(right, src, ROTATE_RIGHT)
	if err != nil {
		t.Fatal(err)
	}
	err = RotateImage(full, src, ROTATE_FULL)
	if err != nil {
		t.Fatal(err)
	}

	err = SaveImage("/tmp/left.jpg", left, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveImage("/tmp/right.jpg", right, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveImage("/tmp/full.jpg", full, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFlipImage(t *testing.T) {
	src, err := OpenRGBAImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	h := image.NewRGBA(src.Bounds())
	v := image.NewRGBA(src.Bounds())

	err = FlipImage(h, src, FLIP_HORIZONTAL)
	if err != nil {
		t.Fatal(err)
	}
	err = FlipImage(v, src, FLIP_VERTICAL)
	if err != nil {
		t.Fatal(err)
	}

	err = SaveImage("/tmp/h.jpg", h, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveImage("/tmp/v.jpg", v, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFlipImage2(t *testing.T) {
	src, err := OpenGRAYImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	h := image.NewGray(src.Bounds())
	v := image.NewGray(src.Bounds())

	err = FlipImage(h, src, FLIP_HORIZONTAL)
	if err != nil {
		t.Fatal(err)
	}
	err = FlipImage(v, src, FLIP_VERTICAL)
	if err != nil {
		t.Fatal(err)
	}

	err = SaveImage("/tmp/h.jpg", h, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveImage("/tmp/v.jpg", v, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
}