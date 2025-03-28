package main

import (
	"github.com/rafaelsouzaribeiro/internal/component/header"
	"github.com/rafaelsouzaribeiro/internal/component/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Application")

	header := header.CreateHeader("Header")

	view1 := view.CreateView1("View 1", header, myWindow)
	view2 := view.CreateView2("View 2", header, myWindow)

	stack := container.NewStack(view1, view2)

	view2.Hide()
	view1.Show()

	myWindow.SetContent(stack)
	myWindow.Resize(fyne.NewSize(800, 480))
	myWindow.ShowAndRun()
}
