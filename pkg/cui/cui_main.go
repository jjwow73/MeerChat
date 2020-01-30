package cui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/wkd3475/MeerChat/pkg/protocol"
	"github.com/wkd3475/MeerChat/pkg/meerchat_node"
	"log"
	"strings"
)

const (
	cmdMode  = 1
	chatMode = 2
)

var (
	views = []string{"room1", "room2", "room3"}
	idxView = 0
	curView = -1
	activeRoom = 0
	curMode = cmdMode
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

func newView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	name := fmt.Sprintf("v%v", idxView)
	v, err := g.SetView(name, 1, 1, int(0.2*float32(maxX))-1, maxY/8)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		fmt.Fprintln(v, strings.Repeat(name+" ", 30))
	}

	views = append(views, name)
	curView = len(views) - 1
	idxView += 1
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

func Cui() {
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

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}