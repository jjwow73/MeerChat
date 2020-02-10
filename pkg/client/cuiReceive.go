package client

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func receiveOutputChannel(g *gocui.Gui) error {
	for cm := range a.outputChannel {
		if cm.c == nil {
			if cm.jsonMessage.Name == "Local" {
				g.Update(func(g *gocui.Gui) error {
					v, err := g.View("chatline")
					if err != nil {
						return err
					}
					v.Clear()
					info := fmt.Sprintf("LOCAL: %s\n", cm.jsonMessage.Content)
					v.Write([]byte(info))
					v.SetCursor(3, 1)
					return nil
				})
			} else {
				g.Update(func(g *gocui.Gui) error {
					v, err := g.View("chatline")
					if err != nil {
						return err
					}
					v.Clear()
					info := fmt.Sprintf("Server: %s\n", cm.jsonMessage.Content)
					v.Write([]byte(info))
					v.SetCursor(3, 1)
					return nil
				})
			}
		} else if cm.c == a.focusedConnection {
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("chat")
				if err != nil {
					return err
				}
				v.Clear()
				info := fmt.Sprintf("%s: %s\n", cm.jsonMessage.Name, cm.jsonMessage.Content)
				v.Write([]byte(info))
				v.SetCursor(3, 1)
				return nil
			})
		} else {
			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("chatline")
				if err != nil {
					return err
				}
				v.Clear()
				info := fmt.Sprintf("WHAT?: %s\n", cm.jsonMessage.Content)
				v.Write([]byte(info))
				v.SetCursor(3, 1)
				return nil
			})
			log.Println(cm.jsonMessage.Name, ":", string(cm.jsonMessage.Content))
		}
	}
	return nil
}
