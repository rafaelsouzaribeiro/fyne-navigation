package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/rafaelsouzaribeiro/internal/component/button"
)

func CreateView1(title string, header *fyne.Container, myWindow fyne.Window) *fyne.Container {
	text1 := canvas.NewText(title, color.Black)
	button1 := button.CreateButton("Ir para View 2", myWindow, []int{1, 0})

	content1 := container.New(layout.NewVBoxLayout(), text1, layout.NewSpacer(), button1)
	return container.New(layout.NewVBoxLayout(), header, content1)
}
