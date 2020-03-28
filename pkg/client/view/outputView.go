package view

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/jjwow73/MeerChat/pkg/chat"
	"log"
)

const (
	maxX = 80
	maxY = 32
)

var (
	g *gocui.Gui
)

func init() {
	gui, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panicln(err)
	}
	g = gui
}

func Start() {
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !gocui.IsQuit(err) {
		log.Panicln(err)
	}
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func layout(g *gocui.Gui) error {
	if v, err := g.SetView("room", 0, 0, int(0.2*float32(maxX)), maxY-11, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Room"
		v.Wrap = true
	}

	if v, err := g.SetView("userinfo", 0, maxY-10, int(0.2*float32(maxX)), maxY-6, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "UserInfo"
		v.Wrap = true
	}

	if v, err := g.SetView("chat", int(0.2*float32(maxX))+1, 0, maxX-1, maxY-6, 0); err != nil {
		if !gocui.IsUnknownView(err) {
			return err
		}
		v.Title = "Chat"
		v.Wrap = true
		v.Autoscroll = true
	}

	return nil
}

func PrintChatMessage(message *chat.MessageProtocol) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("chat")
		if err != nil {
			return err
		}
		msg := fmt.Sprintf("%s : %s\n", message.Name, message.Message)
		v.Write([]byte(msg))
		return nil
	})
}

type Renderer struct{}

func (v Renderer) PrintRoomList(roomList map[string]bool) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("room")
		if err != nil {
			return err
		}
		v.Clear()

		for name, isFocused := range roomList {
			if isFocused {
				name = "*" + name
			}
			name += "\n"
			v.Write([]byte(name))
		}

		return nil
	})
}
