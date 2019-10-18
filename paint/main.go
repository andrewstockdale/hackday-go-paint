package main

import (
	"fmt"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

const (
	KEY_LEFT  uint = 65361
	KEY_UP    uint = 65362
	KEY_RIGHT uint = 65363
	KEY_DOWN  uint = 65364
)

func main() {
	gtk.Init(nil)

	// gui boilerplate
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Paint")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new grid widget to arrange child widgets
	grid, _ := gtk.GridNew()
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	// Create some widgets to put in the grid.
	btn, _ := gtk.ColorButtonNew()
	da, _ := gtk.DrawingAreaNew()

	// Put widgets in grid
	grid.Add(da)
	grid.Add(btn)
	win.Add(grid)
	win.ShowAll()

	// Data !
	unitSize := 1.0
	x := 0.0
	y := 0.0
	keyMap := map[uint]func(){
		KEY_LEFT:  func() { x-- },
		KEY_UP:    func() { y-- },
		KEY_RIGHT: func() { x++ },
		KEY_DOWN:  func() { y++ },
	}

	// Event handlers
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetSourceRGB(0, 0, 0)
		cr.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
		cr.Stroke()
	})
	win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		if move, found := keyMap[keyEvent.KeyVal()]; found {
			move()
			fmt.Printf("on key %f %f\n", x, y)
			win.QueueDraw()
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
		win.QueueDraw()
		
	})

	gtk.Main()
}
