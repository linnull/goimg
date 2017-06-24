package goimg

import (
	"image"
	"os"
	_ "image/jpeg"
	_ "image/png"
	"image/draw"
	"image/color"
	"github.com/pkg/errors"
)

// 打开普通模式的图像文件
func OpenImage(file string) (image.Image, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// 以RGBA模式打开图像
func OpenRGBAImage(file string) (*image.RGBA, error) {
	img, err := OpenImage(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()

	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, img, bounds.Min, draw.Src)
	return imgRGBA, nil
}

// 以灰度模式打开图像
func OpenGRAYImage(file string) (*image.Gray, error) {
	img, err := OpenImage(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()

	imgGRAY := image.NewGray(bounds)
	draw.Draw(imgGRAY, bounds, img, bounds.Min, draw.Src)
	return imgGRAY, nil
}

// 以透明度模式打开图像
func OpenALPHAImage(file string) (*image.Alpha, error) {
	img, err := OpenImage(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()

	imgALPHA := image.NewAlpha(bounds)
	draw.Draw(imgALPHA, bounds, img, bounds.Min, draw.Src)
	return imgALPHA, nil
}

// 以单通道模式打开图像
// 标志有三种
// IMAGE_SINGLE_RED（红通道）
// IMAGE_SINGLE_GREEN（绿通道）
// IMAGE_SINGLE_BLUE（蓝通道）
func OpenSingleImage(file string, flag int) (*image.RGBA, error) {
	img, err := OpenImage(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()

	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, img, bounds.Min, draw.Src)

	minX := imgRGBA.Bounds().Min.X
	minY := imgRGBA.Bounds().Min.Y
	maxX := imgRGBA.Bounds().Max.X
	maxY := imgRGBA.Bounds().Max.Y

	switch flag {
	case IMAGE_SINGLE_RED:
		imgRED := image.NewRGBA(imgRGBA.Bounds())
		rColor := color.RGBA{}
		for x := minX; x < maxX; x++ {
			for y := minY; y < maxY; y++ {
				r, _, _, a := imgRGBA.At(x, y).RGBA()
				rColor.R = uint8(r)
				rColor.A = uint8(a)
				imgRED.Set(x, y, rColor)
			}
		}
		return imgRED, nil
	case IMAGE_SINGLE_GREEN:
		imgGREEN := image.NewRGBA(imgRGBA.Bounds())
		rColor := color.RGBA{}
		for x := minX; x < maxX; x++ {
			for y := minY; y < maxY; y++ {
				_, g, _, a := imgRGBA.At(x, y).RGBA()
				rColor.G = uint8(g)
				rColor.A = uint8(a)
				imgGREEN.Set(x, y, rColor)
			}
		}
		return imgGREEN, nil
	case IMAGE_SINGLE_BLUE:
		imgBLUE := image.NewRGBA(imgRGBA.Bounds())
		rColor := color.RGBA{}
		for x := minX; x < maxX; x++ {
			for y := minY; y < maxY; y++ {
				_, _, b, a := imgRGBA.At(x, y).RGBA()
				rColor.B = uint8(b)
				rColor.A = uint8(a)
				imgBLUE.Set(x, y, rColor)
			}
		}
		return imgBLUE, nil
	default:
		return nil, errors.New("no more single channel")
	}
}
