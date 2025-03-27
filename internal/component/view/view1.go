package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateView1(title string, header *fyne.Container, myWindow fyne.Window) *fyne.Container {
	text1 := canvas.NewText(title, color.Black)
	button1 := widget.NewButtonWithIcon("Ir para View 2", theme.ContentAddIcon(), func() {
		myWindow.Content().(*fyne.Container).Objects[1].Show()
		myWindow.Content().(*fyne.Container).Objects[0].Hide()
	})

	content1 := container.New(layout.NewVBoxLayout(), text1, layout.NewSpacer(), button1)
	return container.New(layout.NewVBoxLayout(), header, content1)
}
