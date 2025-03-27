package button

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateButton(title string, myWindow fyne.Window, Object []int) *widget.Button {
	button := widget.NewButtonWithIcon(title, theme.ContentAddIcon(), func() {
		myWindow.Content().(*fyne.Container).Objects[Object[0]].Show()
		myWindow.Content().(*fyne.Container).Objects[Object[1]].Hide()
	})

	return button
}
