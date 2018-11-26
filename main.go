package main

import (
	"fmt"
	"os"

	"github.com/apsdsm/canvas"
	"github.com/apsdsm/canvas/painter"

	"github.com/gdamore/tcell"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		testRaw()
	}

	if args[0] == "basic" {
		testBasic()
	}

	if args[0] == "box" {
		testBox()
	}

	if args[0] == "fill" {
		testFill()
	}

	if args[0] == "lines" {
		testLines()
	}

	fmt.Println("unknown test")
}

func testRaw() {
	screen := initScreen()

	width, height := screen.Size()

	screen.SetCell(0, 0, tcell.StyleDefault, '+')
	screen.SetCell(width-1, 0, tcell.StyleDefault, '+')
	screen.SetCell(width-1, height-1, tcell.StyleDefault, '+')
	screen.SetCell(0, height-1, tcell.StyleDefault, '+')

	screen.Show()

	waitForQuit(screen)
}

func testBasic() {
	screen := initScreen()

	width, height := screen.Size()

	c := canvas.NewCanvas(screen)
	l := canvas.NewLayer(width, height, 0, 0)

	c.AddLayer(&l)

	l.SetRune(0, 0, '+')
	l.SetRune(width-1, 0, '+')
	l.SetRune(width-1, height-1, '+')
	l.SetRune(0, height-1, '+')

	c.Draw()

	waitForQuit(screen)
}

func testBox() {
	screen := initScreen()

	width, height := screen.Size()

	c := canvas.NewCanvas(screen)
	l := canvas.NewLayer(width, height, 0, 0)

	c.AddLayer(&l)

	painter.DrawBox(&l, 0, 0, l.MaxX, l.MaxY, tcell.StyleDefault)
	painter.DrawBox(&l, 10, 10, l.MaxX-10, l.MaxY-10, tcell.StyleDefault)

	c.Draw()

	waitForQuit(screen)
}

func testFill() {
	screen := initScreen()

	width, height := screen.Size()

	c := canvas.NewCanvas(screen)
	l := canvas.NewLayer(width, height, 0, 0)

	c.AddLayer(&l)

	painter.Fill(&l, 'c', tcell.StyleDefault)

	c.Draw()

	waitForQuit(screen)
}

func testLines() {
	screen := initScreen()

	width, height := screen.Size()

	c := canvas.NewCanvas(screen)
	l := canvas.NewLayer(width, height, 0, 0)

	c.AddLayer(&l)

	painter.DrawVLine(&l, 20, 20, 20, tcell.StyleDefault)
	painter.DrawHLine(&l, 10, 30, 20, tcell.StyleDefault)

	c.Draw()

	waitForQuit(screen)
}

func initScreen() tcell.Screen {
	screen, err := tcell.NewScreen()

	if err != nil {
		fmt.Println("error creating screen")
		os.Exit(1)
	}

	err = screen.Init()

	if err != nil {
		fmt.Println("error initializing screen")
	}

	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	screen.Clear()

	return screen
}

func waitForQuit(screen tcell.Screen) {
	done := false

	for !done {
		polledEvent := screen.PollEvent()

		switch e := polledEvent.(type) {
		case *tcell.EventKey:
			if e.Key() == tcell.KeyRune && string(e.Rune()) == "q" {
				done = true
			}
		}
	}

	screen.Fini()
	os.Exit(0)
}
