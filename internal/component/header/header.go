package header

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func CreateHeader(title string) *fyne.Container {
	headerTitle := canvas.NewText(title, color.Black)
	headerTitle.TextStyle = fyne.TextStyle{Bold: true}
	headerTitle.Alignment = fyne.TextAlignLeading

	header := container.New(layout.NewHBoxLayout(),
		headerTitle,
		layout.NewSpacer(),
	)
	return header
}
