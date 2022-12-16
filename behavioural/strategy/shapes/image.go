package shapes

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"strings"

	"github.com/menxqk/go-design-patterns/behavioural/strategy"
)

type ImageSquare struct {
	strategy.PrintOutput
}

func (i *ImageSquare) Print() error {
	width := 800
	height := 600

	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}

	origin := image.Point{0, 0}
	quality := &jpeg.Options{Quality: 75}

	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)
	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)

	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)

	if i.Writer == nil {
		return errors.New("no writer stored on ImageSquare")
	}

	if err := jpeg.Encode(i.Writer, bgImage, quality); err != nil {
		return errors.New("error writing image")
	}

	if i.LogWriter != nil {
		io.Copy(i.LogWriter, strings.NewReader("image written in provided writer\n"))
	}

	return nil
}
