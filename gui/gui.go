//go:generate fyne bundle -o bundled.go ../assets
package gui

import (
	"fmt"
	"jf/Ratatouille/staging"
	"jf/Ratatouille/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func InitGui() {
	a := app.New()
	a.Settings().SetTheme(newTheme())
	w := a.NewWindow("Onlyfans Hack")
	w.Resize(fyne.NewSize(500, 400))

	w.SetContent(makeGui())
	w.ShowAndRun()
}

func makeGui() fyne.CanvasObject {
	left := widget.NewLabel("Left")
	right := widget.NewLabel("right")
	userNameInput := widget.NewEntry()
	// userNameInput.
	userNameInput.SetPlaceHolder("Please Enter you Onlyfans Username")
	userPassword := widget.NewPasswordEntry()
	userPassword.SetPlaceHolder("Please Enter you Onlyfans Password")
	targetCreatorName := widget.NewEntry()
	targetCreatorName.SetPlaceHolder("Please Enter your Target Creator")
	// mainContainer:=container.NewBorder()
	content := container.NewGridWithRows(3, targetCreatorName, userNameInput, userPassword)
	return container.NewBorder(makeBanner(), nil, left, right, content)
}
func makeBanner() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() { staging.Stage(utils.GetOs()) }))
	logo := canvas.NewImageFromResource(resourceOnlyfanslogoPng)
	logo.FillMode = canvas.ImageFillContain

	return container.NewStack(toolbar, logo)
}

func toolBarHome() {
	fmt.Println("I am home")
}
