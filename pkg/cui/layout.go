package cui

import (
	"fmt"
	"github.com/jroimartin/gocui"
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

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("room", 0, 0, int(0.2*float32(maxX)), maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "room"
		v.Wrap = true
	}

	if v, err := g.SetView("chat", int(0.2*float32(maxX)), 0, maxX, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "chat"
		v.Wrap = true
		v.Autoscroll = true
	}

	if v, err := g.SetView("chatline", int(0.2*float32(maxX)), maxY-10, maxX, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "chatline"
		v.Wrap = true
		v.Autoscroll = true
		v.Editable = true
	}

	if v, err := g.SetView("cmdline", -1, maxY-5, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "cmdline"
		v.Wrap = true
		v.Autoscroll = true
		v.Editable = true

		if err = newLine(g, v); err != nil {
			return err
		}

		if _, err = setCurrentViewOnTop(g, "cmdline"); err != nil {
			return err
		}
	}



	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func newLine(g *gocui.Gui, v *gocui.View) error {
	if curMode!= cmdMode {
		return nil
	}
	v.Clear()
	_, err := fmt.Fprint(v, ">>")
	v.SetCursor(3, 0)
	return err
}

func sendMessage(g *gocui.Gui, v *gocui.View) error {
	if curMode!= chatMode {
		return nil
	}
	v.Clear()
	return nil
}

func lockLeftKey(g *gocui.Gui, v *gocui.View) error {
	x, y := v.Cursor()

	var err error
	if x==3&&y==0 {
		err = v.SetCursor(3, 0)
		return err
	}

	v.MoveCursor(-1, 0, true)
	return nil
}

func lockUpKey(g *gocui.Gui, v *gocui.View) error {
	x, y := v.Cursor()

	var err error
	if x>=0&&x<=3&&y==1 {
		err = v.SetCursor(3, 0)
		return err
	}

	v.MoveCursor(0, -1, true)
	return nil
}

func lockBackspace2Key(g *gocui.Gui, v *gocui.View) error {
	x, _ := v.Cursor()

	var err error
	if x>=0&&x<=3 {
		err = v.SetCursor(3, 0)
		return err
	}

	v.EditDelete(true)
	return nil
}

func changeMode(g *gocui.Gui, v *gocui.View) error {
	if curMode == cmdMode {
		if _, err := setCurrentViewOnTop(g, "chatline"); err != nil {
			return err
		}
		curMode = chatMode
		return nil
	}

	if curMode == chatMode {
		if _, err := setCurrentViewOnTop(g, "cmdline"); err != nil {
			return err
		}
		curMode = cmdMode
		return nil
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("chatline", gocui.KeyEnter, gocui.ModNone, sendMessage); err != nil {
		return err
	}
	if err := g.SetKeybinding("cmdline", gocui.KeyEnter, gocui.ModNone, newLine); err != nil {
		return err
	}
	if err := g.SetKeybinding("cmdline", gocui.KeyBackspace2, gocui.ModNone, lockBackspace2Key); err != nil {
		return err
	}
	if err := g.SetKeybinding("cmdline", gocui.KeyArrowLeft, gocui.ModNone, lockLeftKey); err != nil {
		return err
	}
	if err := g.SetKeybinding("cmdline", gocui.KeyArrowUp, gocui.ModNone, lockUpKey); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlA, gocui.ModNone, changeMode); err != nil {
		return err
	}

	return nil
}

func Cui() {
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