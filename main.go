package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type DrawableRectangle struct {
	widget.BaseWidget
	rectangles []fyne.CanvasObject
	startPos   fyne.Position
	endPos     fyne.Position
	drawing    bool
}

func NewDrawableRectangle() *DrawableRectangle {
	dr := &DrawableRectangle{}
	dr.ExtendBaseWidget(dr)
	return dr
}

type drawableRectangleRenderer struct {
	drawableRectangle *DrawableRectangle
	background        *canvas.Rectangle
}

func (d *DrawableRectangle) CreateRenderer() fyne.WidgetRenderer {
	background := canvas.NewRectangle(color.White)
	return &drawableRectangleRenderer{
		drawableRectangle: d,
		background:        background,
	}
}

func (r *drawableRectangleRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, len(r.drawableRectangle.rectangles)+1)
	objects = append(objects, r.background)
	objects = append(objects, r.drawableRectangle.rectangles...)
	return objects
}

func (r *drawableRectangleRenderer) Destroy() {}

func (r *drawableRectangleRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
}

func (r *drawableRectangleRenderer) MinSize() fyne.Size {
	return fyne.NewSize(100, 100)
}

func (r *drawableRectangleRenderer) Refresh() {
	if r.drawableRectangle.drawing {
		tempRect := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 128})
		tempRect.Move(fyne.NewPos(min(r.drawableRectangle.startPos.X, r.drawableRectangle.endPos.X), min(r.drawableRectangle.startPos.Y, r.drawableRectangle.endPos.Y)))
		tempRect.Resize(fyne.NewSize(abs(r.drawableRectangle.endPos.X-r.drawableRectangle.startPos.X), abs(r.drawableRectangle.endPos.Y-r.drawableRectangle.startPos.Y)))
		r.drawableRectangle.rectangles = append(r.drawableRectangle.rectangles, tempRect)
	}
	canvas.Refresh(r.drawableRectangle)
}

func (d *DrawableRectangle) MouseDown(ev *desktop.MouseEvent) {
	d.startPos = ev.Position
	d.drawing = true
}

func (d *DrawableRectangle) MouseUp(ev *desktop.MouseEvent) {
	if d.drawing {
		d.endPos = ev.Position
		d.addRectangle()
		d.drawing = false
		d.Refresh()
	}
}

func (d *DrawableRectangle) MouseMoved(ev *desktop.MouseEvent) {
	if d.drawing {
		d.endPos = ev.Position
		d.Refresh()
	}
}

func (d *DrawableRectangle) addRectangle() {
	rect := canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 255, A: 128})
	rect.Move(fyne.NewPos(min(d.startPos.X, d.endPos.X), min(d.startPos.Y, d.endPos.Y)))
	rect.Resize(fyne.NewSize(abs(d.endPos.X-d.startPos.X), abs(d.endPos.Y-d.startPos.Y)))
	d.rectangles = append(d.rectangles, rect)
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func abs(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Interactive Rectangle Drawing")

	drawableRect := NewDrawableRectangle()
	w.SetContent(drawableRect)

	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
