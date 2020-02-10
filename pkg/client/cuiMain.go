package client

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
)

const (
	cmdMode  = 1
	chatMode = 2
	maxX     = 80
	maxY     = 32
)

var (
	activeRoom     = 0
	totalRoom      = 0
	curMode        = cmdMode
	views          = make(map[string]*gocui.View)
	roomUpdateChan = make(chan bool)
)

type Label struct {
	name string
	w, h int
	body string
}

func NewLabel(name string, body string) *Label {
	lines := strings.Split(body, "\n")

	w := 0
	for _, l := range lines {
		if len(l) > w {
			w = len(l)
		}
	}
	h := len(lines) + 1
	w = w + 1

	return &Label{name: name, w: w, h: h, body: body}
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func UpdateRoom(g *gocui.Gui, timeout chan bool) error {
	rooms := a.connections
	time.Sleep(2 * time.Second)
	for {
		select {
		case <-timeout:
			g.Update(func(g *gocui.Gui) error {
				x := 1
				y := 1
				dy := 3
				for idx, room := range rooms {
					roomName := fmt.Sprintf("room[%v]", idx)
					v, err := g.SetView(roomName, x, y, int(0.2*float32(maxX))-1, y+dy)
					if err != nil {
						if err != gocui.ErrUnknownView {
							return err
						}
						v.Wrap = true
						v.Write([]byte(room.toString()))
					}
					y = y + dy + 1
				}
				return nil
			})
		}
	}
	return nil
}

func newLine(g *gocui.Gui, v *gocui.View) error {
	if curMode != cmdMode {
		return nil
	}
	v.Clear()
	_, err := fmt.Fprint(v, ">> ")
	v.SetCursor(3, 0)
	return err
}

func CuiMain() {
	timeout := make(chan bool)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			timeout <- true
		}
	}()

	go UpdateRoom(g, timeout)
	go receiveOutputChannel(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
