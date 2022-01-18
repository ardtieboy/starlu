package imageprocessing

import (
	"errors"
	"math"
	"strings"

	"github.com/nfnt/resize"

	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

// SubImager is an interface which contains a method SubImage
type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

// Drawable is an interface needed in order to draw on an image
type Drawable interface {
	Set(x, y int, c color.Color)
}

// func guessImageFormat(r os.Reader) (format string, err error) {
// 	_, format, err = image.DecodeConfig(r)
// 	return
// }

// Crop blabla
func Crop(originalImagePath string) (string, error) {
	// Read source image

	splittedFilename := strings.Split(originalImagePath, ".")
	basicFilename := splittedFilename[0]
	format := splittedFilename[len(splittedFilename)-1]

	sourceFile, err := os.Open(originalImagePath)
	if err != nil {
		return "error", errors.New("Could not open image at " + originalImagePath)
	}
	defer sourceFile.Close()

	var sourceImage image.Image

	if format == "jpg" || format == "jpeg" {
		sourceImage, err = jpeg.Decode(sourceFile)
		if err != nil {
			return "error", errors.New("Could not decode the image " + originalImagePath)
		}
	} else if format == "png" {
		sourceImage, err = png.Decode(sourceFile)
		if err != nil {
			return "error", errors.New("Could not decode the image " + originalImagePath)
		}
	} else {
		return "error", errors.New("Only jpeg and png images are supported " + originalImagePath)
	}

	maxX := sourceImage.Bounds().Max.X
	maxY := sourceImage.Bounds().Max.Y

	if maxX < 1200 || maxY < 675 {
		return "error", errors.New("Image too small: " + originalImagePath)
	}

	xRatio := float64(maxX) / 1200
	yRatio := float64(maxY) / 675.0

	smallestRatio := math.Min(xRatio, yRatio)

	rectX := int(1200 * smallestRatio)
	rectY := int(675 * smallestRatio)

	cropableImage, ok := sourceImage.(SubImager)

	if ok != true {
		return "error", errors.New("Not a SubImage interface: " + originalImagePath)
	}

	marginX := (maxX - rectX) / 2.0
	marginY := (maxY - rectY) / 2.0

	croppedFullSizeImage := cropableImage.SubImage(image.Rect(marginX, marginY, marginX+int(rectX), marginY+int(rectY)))
	croppedImage := resize.Thumbnail(1200, 675, croppedFullSizeImage, resize.Lanczos3)

	borderwidth := int(23)

	croppedImageBoundsInOriginalAxis := croppedImage.Bounds()
	drawableCropped := image.NewRGBA(croppedImageBoundsInOriginalAxis)
	// Note that there is a offset needed because the cropped image is still in the original image coordinate system
	// We need to create a new image which is drawable in order to draw borders on them
	draw.Draw(drawableCropped, croppedImageBoundsInOriginalAxis, croppedImage, croppedImageBoundsInOriginalAxis.Min, draw.Src)

	// Shift the rectangles in the case the croppedImageBoundsInOriginalAxis.Min doesn't match (0, 0)
	topBorder := image.Rect(0, 0, 1200, borderwidth)
	lowerBorder := image.Rect(0, 675-borderwidth, 1200, 675)
	leftBorder := image.Rect(0, 0, borderwidth, 675)
	rightBorder := image.Rect(1200-borderwidth, 0, 1200, 675)
	colorRed := color.RGBA{192, 51, 29, 255}

	draw.Draw(drawableCropped, topBorder.Add(croppedImageBoundsInOriginalAxis.Min), &image.Uniform{colorRed}, croppedImageBoundsInOriginalAxis.Min, draw.Src)
	draw.Draw(drawableCropped, lowerBorder.Add(croppedImageBoundsInOriginalAxis.Min), &image.Uniform{colorRed}, croppedImageBoundsInOriginalAxis.Min, draw.Src)
	draw.Draw(drawableCropped, leftBorder.Add(croppedImageBoundsInOriginalAxis.Min), &image.Uniform{colorRed}, croppedImageBoundsInOriginalAxis.Min, draw.Src)
	draw.Draw(drawableCropped, rightBorder.Add(croppedImageBoundsInOriginalAxis.Min), &image.Uniform{colorRed}, croppedImageBoundsInOriginalAxis.Min, draw.Src)

	// Save cropped image
	outputImageName := basicFilename + "_pragafied." + format
	croppedFile, _ := os.Create(outputImageName)
	defer croppedFile.Close()

	if format == "jpg" || format == "jpeg" {
		jpeg.Encode(croppedFile, drawableCropped, &jpeg.Options{100})
	} else if format == "png" {
		png.Encode(croppedFile, drawableCropped)
	} else {
		return "error", errors.New("Only jpeg and png images are supported")
	}
	return outputImageName, nil
}
