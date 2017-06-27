package goimg

import (
	"testing"
)

func TestVerticalReversed(t *testing.T) {
	src, err := OpenRGBAImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	dst := LevelReversed(src)

	err = SaveImage("/tmp/lr.jpg", dst, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLevelReversed(t *testing.T) {
	src, err := OpenRGBAImage(jpegfile)
	if err != nil {
		t.Fatal(err)
	}

	dst := VerticalReversed(src)

	err = SaveImage("/tmp/vr.jpg", dst, SAVE_JPEG)
	if err != nil {
		t.Fatal(err)
	}
}