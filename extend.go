package goimg

import (
	"image"
	"image/draw"
)

// 垂直增加镜像图像，仅支持RGBA格式的图像
// 处理完的图像是源图像的两倍（y_=2*y）
func LevelReversed(src *image.RGBA) *image.RGBA {
	if src == nil {
		return nil
	}

	bounds := src.Bounds()
	if bounds.Empty() {
		return nil
	}

	dst := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), 2 * bounds.Dy()))
	dstB1 := image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	dstB2 := image.Rect(0, bounds.Dy(), bounds.Dx(), 2 * bounds.Dy())

	src_ := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	FlipImage(src_, src, FLIP_VERTICAL)

	draw.Draw(dst, dstB1, src, src.Bounds().Min, draw.Src)
	draw.Draw(dst, dstB2, src_, src_.Bounds().Min, draw.Src)

	return dst
}

// 水平增加镜像图像，仅支持RGBA格式的图像
// 处理完的图像是源图像的两倍（x_=2*x）
func VerticalReversed(src *image.RGBA) *image.RGBA {
	if src == nil {
		return nil
	}

	bounds := src.Bounds()
	if bounds.Empty() {
		return nil
	}

	dst := image.NewRGBA(image.Rect(0, 0, 2 * bounds.Dx(), bounds.Dy()))
	dstB1 := image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	dstB2 := image.Rect(bounds.Dx(), 0, 2 * bounds.Dx(), bounds.Dy())

	src_ := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	FlipImage(src_, src, FLIP_HORIZONTAL)

	draw.Draw(dst, dstB1, src, src.Bounds().Min, draw.Src)
	draw.Draw(dst, dstB2, src_, src_.Bounds().Min, draw.Src)

	return dst
}

