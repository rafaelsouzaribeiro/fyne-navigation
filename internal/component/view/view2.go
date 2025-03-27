package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/rafaelsouzaribeiro/internal/component/button"
)

func CreateView2(title string, header *fyne.Container, myWindow fyne.Window) *fyne.Container {
	text2 := canvas.NewText(title, color.Black)
	button2 := button.CreateButton("Ir para View 1", myWindow, []int{0, 1})

	content2 := container.New(layout.NewVBoxLayout(), text2, layout.NewSpacer(), button2)
	return container.New(layout.NewVBoxLayout(), header, content2)
}
