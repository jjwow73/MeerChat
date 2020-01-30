package meerchat_node

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/wkd3475/MeerChat/pkg/meerchat_node"
	"github.com/wkd3475/MeerChat/pkg/protocol"
	"log"
	"strings"
	"time"
)

const (
	cmdMode  = 1
	chatMode = 2
	maxX = 80
	maxY = 32
)

var (
	activeRoom = 0
	totalRoom = 0
	curMode = cmdMode
	views = make(map[string]*gocui.View)
)

var MeerNode = meerchat_node.NewNode()
var cuiChan = make(chan protocol.Command)

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
	time.Sleep(2 * time.Second)
	for {
		select {
		case <-timeout:
			g.Update(func(g *gocui.Gui) error {
				x := 1
				y := 1
				dy := 3
				for idx := range MeerNode.Room {
					roomName := fmt.Sprintf("room%d", idx)
					v, err := g.SetView(roomName, x, y, int(0.2*float32(maxX))-1, y+dy)
					if err != nil {
						if err != gocui.ErrUnknownView {
							return err
						}
						v.Wrap = true
						v.Write([]byte(MeerNode.Room[idx].GetName()))
					}
					y = y+dy+1
				}
				return nil
			})
		}


		time.Sleep(1 * time.Second)
	}
	return nil
}

func newLine(g *gocui.Gui, v *gocui.View) error {
	if curMode!= cmdMode {
		return nil
	}
	v.Clear()
	_, err := fmt.Fprint(v, ">> ")
	v.SetCursor(3, 0)
	return err
}

func CuiMain() {
	timeout := make(chan bool)
	go MeerNode.CommandReceiver(cuiChan)

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
	} ()

	go UpdateRoom(g, timeout)

	if err := g.MainLoop();  err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}