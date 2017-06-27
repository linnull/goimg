package goimg

import (
	"image"
	"os"
	"image/jpeg"
	"image/png"
	"image/draw"
	"image/color"
	"errors"
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
// 通道标志有三种
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

	switch flag {
	case IMAGE_SINGLE_RED:
		singleImage(imgRGBA, IMAGE_SINGLE_RED)
		return imgRGBA, nil
	case IMAGE_SINGLE_GREEN:
		singleImage(imgRGBA, IMAGE_SINGLE_GREEN)
		return imgRGBA, nil
	case IMAGE_SINGLE_BLUE:
		singleImage(imgRGBA, IMAGE_SINGLE_BLUE)
		return imgRGBA, nil
	default:
		return nil, errors.New("no more single channel")
	}
}

// 选择通道
func singleImage(imgRGBA *image.RGBA, flay int) *image.RGBA {
	minX := imgRGBA.Bounds().Min.X
	minY := imgRGBA.Bounds().Min.Y
	maxX := imgRGBA.Bounds().Max.X
	maxY := imgRGBA.Bounds().Max.Y

	rColor := color.RGBA{}
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			r, g, b, a := imgRGBA.At(x, y).RGBA()
			switch flay {
			case IMAGE_SINGLE_RED:
				rColor.R = uint8(r)
			case IMAGE_SINGLE_GREEN:
				rColor.G = uint8(g)
			case IMAGE_SINGLE_BLUE:
				rColor.B = uint8(b)
			}
			rColor.A = uint8(a)
			imgRGBA.Set(x, y, rColor)
		}
	}
	return imgRGBA
}

// 保存图像文件
// file: 图像路径
// img:  图像对象
// flag: 图像格式（SAVE_JPEG,SAVE_PNG）
func SaveImage(file string, img image.Image, flag int) error {
	isExist := IsExist(file)
	if isExist {
		return errors.New("file is exist")
	}

	if img == nil {
		return errors.New("the img is nil")
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	switch flag {
	case SAVE_JPEG:
		err = jpeg.Encode(f, img, nil)
		if err != nil {
			return err
		}
		return nil
	case SAVE_PNG:
		err = png.Encode(f, img)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("wrong save flag")
	}
}

// 旋转图像
// flag:
// ROTATE_LEFT(向左旋转)
// ROTATE_RIGHT(向右旋转)
// ROTATE_FULL(180度旋转)
// 源图像和目标图像的大小要保存相应
func RotateImage(dst image.Image, src image.Image, flag int) error {
	if src == nil || dst == nil {
		return errors.New("image is nil")
	}

	srcb := src.Bounds()
	dstb := dst.Bounds()

	if srcb.Empty() || dstb.Empty() {
		return errors.New("image is empty")
	}

	switch flag {
	case ROTATE_LEFT, ROTATE_RIGHT:
		if !(srcb.Dx() == dstb.Dy() && srcb.Dy() == dstb.Dx()) {
			return errors.New("wrong image size")
		}
	case ROTATE_FULL:
		if !(srcb.Dx() == dstb.Dx() && srcb.Dy() == dstb.Dy()) {
			return errors.New("wrong image size")
		}
	default:
		return errors.New("wrong flag")
	}

	dstRGBA, dstOk := dst.(*image.RGBA)
	srcRGBA, srcOk := src.(*image.RGBA)
	if dstOk && srcOk{
		return transformRGBA(dstRGBA, srcRGBA, flag)
	}

	dstGRAY, dstOk := dst.(*image.Gray)
	srcGRAY, srcOk := src.(*image.Gray)
	if dstOk && srcOk{
		return transformGRAY(dstGRAY, srcGRAY, flag)
	}

	return errors.New("wrong image format")

}

// 翻转图像
// flag:
// FLIP_HORIZONTAL(水平翻转)
// FLIP_VERTICAL(垂直翻转)
// 源图像和目标图像的大小要保存相应
func FlipImage(dst image.Image, src image.Image, flag int) error {
	if src == nil || dst == nil {
		return errors.New("image is nil")
	}

	srcb := src.Bounds()
	dstb := dst.Bounds()

	if srcb.Empty() || dstb.Empty() {
		return errors.New("image is empty")
	}

	switch flag {
	case FLIP_HORIZONTAL, FLIP_VERTICAL:
		if !(srcb.Dx() == dstb.Dx() && srcb.Dy() == dstb.Dy()) {
			return errors.New("wrong image size")
		}
	default:
		return errors.New("wrong flag")
	}

	dstRGBA, dstOk := dst.(*image.RGBA)
	srcRGBA, srcOk := src.(*image.RGBA)
	if dstOk && srcOk{
		return transformRGBA(dstRGBA, srcRGBA, flag)
	}

	dstGRAY, dstOk := dst.(*image.Gray)
	srcGRAY, srcOk := src.(*image.Gray)
	if dstOk && srcOk{
		return transformGRAY(dstGRAY, srcGRAY, flag)
	}

	return errors.New("wrong image format")
}

func transformRGBA(dst *image.RGBA, src *image.RGBA, flag int) error {
	sb := src.Bounds()
	db := dst.Bounds()

	sminx := sb.Min.X
	sminy := sb.Min.Y
	smaxx := sb.Max.X
	smaxy := sb.Max.Y
	dmaxx := db.Max.X
	dmaxy := db.Max.Y
	var c color.Color

	switch flag {
	case ROTATE_LEFT:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, y, dmaxy-1-x) {
					return errors.New("transform error")
				}
				dst.Set(y, dmaxy-1-x, c)
			}
		}
	case ROTATE_RIGHT:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, dmaxx-1-y, x) {
					return errors.New("transform error")
				}
				dst.Set(dmaxx-1-y, x, c)
			}
		}
	case ROTATE_FULL:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, dmaxx-1-x, dmaxy-1-y) {
					return errors.New("transform error")
				}
				dst.Set(dmaxx-1-x, dmaxy-1-y, c)
			}
		}
	case FLIP_HORIZONTAL:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, dmaxx-1-x, y) {
					return errors.New("transform error")
				}
				dst.Set(dmaxx-1-x, y, c)
			}
		}
	case FLIP_VERTICAL:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, x, dmaxy-1-y) {
					return errors.New("transform error")
				}
				dst.Set(x, dmaxy-1-y, c)
			}
		}
	}

	return nil
}

func transformGRAY(dst *image.Gray, src *image.Gray, flag int) error {
	sb := src.Bounds()
	db := dst.Bounds()

	sminx := sb.Min.X
	sminy := sb.Min.Y
	smaxx := sb.Max.X
	smaxy := sb.Max.Y
	dmaxx := db.Max.X
	dmaxy := db.Max.Y
	var c color.Color

	switch flag {
	case ROTATE_LEFT:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, y, dmaxy-1-x) {
					return errors.New("transform error")
				}
				dst.Set(y, dmaxy-1-x, c)
			}
		}
	case ROTATE_RIGHT:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, dmaxx-1-y, x) {
					return errors.New("transform error")
				}
				dst.Set(dmaxx-1-y, x, c)
			}
		}
	case ROTATE_FULL:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, dmaxx-1-x, dmaxy-1-y) {
					return errors.New("transform error")
				}
				dst.Set(dmaxx-1-x, dmaxy-1-y, c)
			}
		}
	case FLIP_HORIZONTAL:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, dmaxx-1-x, y) {
					return errors.New("transform error")
				}
				dst.Set(dmaxx-1-x, y, c)
			}
		}
	case FLIP_VERTICAL:
		for x := sminx; x < smaxx; x++ {
			for y := sminy; y < smaxy; y++ {
				c = src.At(x, y)
				if !inBounds(db, x, dmaxy-1-y) {
					return errors.New("transform error")
				}
				dst.Set(x, dmaxy-1-y, c)
			}
		}
	}

	return nil
}

func inBounds(b image.Rectangle, x, y int) bool {
	if x < b.Min.X || x >= b.Max.X {
		return false
	}
	if y < b.Min.Y || y >= b.Max.Y {
		return false
	}
	return true
}

// 裁剪图像
// src:源图像
// x0,x1,y0,y1:裁剪的定点矩阵参数
// 返回新的裁剪图像
func CutImage(src image.Image, x0, y0, x1, y1 int) image.Image {
	if src == nil {
		return nil
	}

	bounds := src.Bounds()

	if bounds.Empty() {
		return nil
	}

	srcRGBA, srcOk := src.(*image.RGBA)
	if srcOk{
		dst := image.NewRGBA(image.Rect(0, 0, x1-x0, y1-y0))
		draw.Draw(dst, dst.Rect, srcRGBA, image.Point{x0, y0}, draw.Src)
		return dst
	}

	srcGRAY, srcOk := src.(*image.Gray)
	if srcOk{
		dst := image.NewGray(image.Rect(0, 0, x1-x0, y1-y0))
		draw.Draw(dst, dst.Rect, srcGRAY, image.Point{x0, y0}, draw.Src)
		return dst
	}

	return nil
}
