package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Go canvas")

	text1 := canvas.NewText("Editor", color.White)
	text2 := canvas.NewText("Scan", color.White)
	text3 := canvas.NewText("Infer", color.White)
	content := container.New(layout.NewVBoxLayout(), layout.NewSpacer(),text1, text2, text3, layout.NewSpacer())

	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewHBoxLayout(), content, layout.NewSpacer(), centered, layout.NewSpacer()))
	myWindow.ShowAndRun()
}

