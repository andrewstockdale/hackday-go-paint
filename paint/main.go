package main

import (
	"log"
	// "fmt"
	"github.com/gotk3/gotk3/cairo"
	// "github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

const (
	KEY_LEFT  uint = 65361
	KEY_UP    uint = 65362
	KEY_RIGHT uint = 65363
	KEY_DOWN  uint = 65364
)

// Surface object for persistance
var surface cairo.Surface

// Clear surface in white colour 
func clearSurface() {
	cr := cairo.Create(&surface)
	cr.SetSourceRGB(255,255,255)
	cr.Paint()
	cr.Close()
}

// Create new surface of same window-size to draw in
func  createSurface(widget *gtk.DrawingArea) bool {
	// if (cairo.Surface{}) == surface  {
	// 	surface.destroy()
	// }

	// gtkWindow, _ := widget.GetWindow()
	windowWidth := widget.GetAllocatedWidth()
	windowHeight := widget.GetAllocatedHeight()
	surface = *surface.CreateSimilar(cairo.CONTENT_COLOR, windowWidth, windowHeight)
	clearSurface()

	return true
}

func main() {
	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	// frame, _ := gtk.FrameNew("Example")
	drawing_area, _ := gtk.DrawingAreaNew()

	window.Add(drawing_area)
	window.SetTitle("KindBiever")
	window.Connect("destroy", gtk.MainQuit)
	window.SetDefaultSize(1000, 1000)
	window.ShowAll()

	// Data !
	// unitSize := 1.0
	// x := 0.0
	// y := 0.0
	// keyMap := map[uint]func(){
	// 	KEY_LEFT:  func() { x-- },
	// 	KEY_UP:    func() { y-- },
	// 	KEY_RIGHT: func() { x++ },
	// 	KEY_DOWN:  func() { y++ },
	// }

	// Event handlers
	// drawing_area.Connect("configure-event", createSurface(drawing_area))
	// drawing_area.Connect("draw", func(drawing_area *gtk.DrawingArea, cr *cairo.Context) {
	// 	cr.SetSourceRGB(0, 0, 0)
	// 	cr.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
	// 	cr.Stroke()
	// })
	// window.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
	// 	keyEvent := &gdk.EventKey{ev}
	// 	if move, found := keyMap[keyEvent.KeyVal()]; found {
	// 		move()
	// 		fmt.Printf("on key %f %f\n", x, y)
	// 		win.QueueDraw()
	// 	}
	// })

	// window.Connect("motion-notify-event", func(win *gtk.Window, ev *gdk.Event) {
	// // win.Connect("button-press-event", func(win *gtk.Window, ev *gdk.Event) {
		
	// 	motion := &gdk.EventMotion{ev}
	// 	xx, yy := motion.MotionVal()
	// 	// fmt.Printf("test")
	// 	fmt.Printf("x: %f y: %f\n", xx, yy)

	// 	// if move, found := keyMap[keyEvent.KeyVal()]; found {
	// 	// gtk.widget_set_focus_on_click()
	// 	// x--
	// 	x = xx
	// 	y = yy
	// 	win.QueueDraw()
		
	// })

	gtk.Main()
}
