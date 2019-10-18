package main

import (
	"fmt"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)
	const width = 1000
	const height = 500

	// gui boilerplate
	var matrix [width][height]float64
	var colorMatrix [width][height]*gdk.RGBA

	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	win.SetTitle("Paint")
	win.SetDefaultSize(width, height)
	win.SetResizable(false)
	// On quit
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new grid widget to arrange child widgets
	grid, _ := gtk.GridNew()
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	// Create some widgets to put in the grid.
	btn, _ := gtk.ColorButtonNew()
	defaultBrushColour := gdk.NewRGBA(0,0,0,1)
	btn.ColorChooser.SetRGBA(defaultBrushColour)

	da, _ := gtk.DrawingAreaNew()
	da.SetSizeRequest(width,height)

	scaler, _ := gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, 1, 50, 1)
	scaler.SetSizeRequest(50,50)

	// Put widgets in grid
	grid.Add(da)
	grid.Add(scaler)
	grid.Add(btn)
	win.Add(grid)

	win.ShowAll()


	// Data !
	unitSize := 1.0
	x := 0.0
	y := 0.0


	// Event handlers
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				if matrix[i][j] > 0 {
					cr.Rectangle(float64(i)*unitSize, float64(j)*unitSize, unitSize, unitSize)
					cr.SetLineWidth(matrix[i][j])

					chosenRGBAValue := colorMatrix[i][j].Floats()
					cr.SetSourceRGB(chosenRGBAValue[0],chosenRGBAValue[1],chosenRGBAValue[2])
					cr.Stroke()
				}
			}
		}
	})

	win.Connect("motion-notify-event", func(win *gtk.Window, ev *gdk.Event) {
		// win.Connect("button-press-event", func(win *gtk.Window, ev *gdk.Event) {
		
		motion := &gdk.EventMotion{ev}
		xx, yy := motion.MotionVal()
		// fmt.Printf("test")
		fmt.Printf("x: %f y: %f\n", xx, yy)

		// if move, found := keyMap[keyEvent.KeyVal()]; found {
		// gtk.widget_set_focus_on_click()
		// x--
		x = xx
		y = yy
		matrix[int(x)][int(y)] = scaler.GetValue()
		colorMatrix[int(x)][int(y)] = btn.ColorChooser.GetRGBA()

		win.QueueDraw()
		
	})

	gtk.Main()
}
