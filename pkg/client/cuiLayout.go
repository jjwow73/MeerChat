package client

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("room", 0, 0, int(0.2*float32(maxX)), maxY-11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Room"
		v.Wrap = true
		views[v.Name()] = v
	}

	if v, err := g.SetView("userinfo", 0, maxY-10, int(0.2*float32(maxX)), maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "UserInfo"
		v.Wrap = true
		var addr, roomID, nickname string
		if a.focusedConnection != nil {
			addr, roomID, nickname = a.focusedConnection.GetConnInfo()
		} else {
			addr, roomID, nickname = "NONE", "NONE", "NONE"
		}
		info := fmt.Sprintf("%s\n%s\n:%s", addr, roomID, nickname)
		v.Write([]byte(info))
		views[v.Name()] = v
	}

	if v, err := g.SetView("chat", int(0.2*float32(maxX))+1, 0, maxX-1, maxY-11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Chat"
		v.Wrap = true
		v.Autoscroll = true
		views[v.Name()] = v
	}

	if v, err := g.SetView("chatline", int(0.2*float32(maxX))+1, maxY-10, maxX-1, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Chatline(ctrl+a)"
		v.Wrap = true
		v.Autoscroll = true
		v.Editable = true
		views[v.Name()] = v
	}

	if v, err := g.SetView("cmdline", 0, maxY-5, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Cmdline"
		v.Wrap = true
		v.Autoscroll = true
		v.Editable = true
		views[v.Name()] = v

		if err = newLine(g, v); err != nil {
			return err
		}

		if _, err = setCurrentViewOnTop(g, "cmdline"); err != nil {
			return err
		}
	}

	return nil
}
