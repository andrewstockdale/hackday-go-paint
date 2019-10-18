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
	var matrix[width][height] bool
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetDefaultSize(width,height)
	win.SetResizable(false)
	da, _ := gtk.DrawingAreaNew()

	// da.setDoubleBuffered(false)
	win.Add(da)
	win.SetTitle("Arrow keys")
	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()


	// Data !
	unitSize := 1.0
	x := 0.0
	y := 0.0
	

	// Event handlers
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetSourceRGB(0, 0, 0)
		cr.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
		cr.Stroke()
		for i := 0; i < width; i++ {
        	for j := 0; j < height; j++ {
        		if matrix[i][j]{
        			cr.Rectangle(float64(i)*unitSize, float64(j)*unitSize, unitSize, unitSize)

        		}
        	}
        }
		cr.Stroke()


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
		matrix[int(x)][int(y)] = true
		win.QueueDraw()
		
	})

	gtk.Main()
}
