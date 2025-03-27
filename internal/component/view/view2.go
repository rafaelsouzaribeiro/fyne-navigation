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

func CreateView2(title string, header *fyne.Container, myWindow fyne.Window) *fyne.Container {
	text2 := canvas.NewText(title, color.Black)
	button2 := widget.NewButtonWithIcon("Voltar para View 1", theme.ContentAddIcon(), func() {
		myWindow.Content().(*fyne.Container).Objects[0].Show()
		myWindow.Content().(*fyne.Container).Objects[1].Hide()
	})

	content2 := container.New(layout.NewVBoxLayout(), text2, layout.NewSpacer(), button2)
	return container.New(layout.NewVBoxLayout(), header, content2)
}
