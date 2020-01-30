package meerchat_node

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/wkd3475/MeerChat/pkg/room"
	"regexp"
	"strings"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func sendCommand(g *gocui.Gui, v *gocui.View) error {
	if curMode != cmdMode {
		return nil
	}
	command := []byte(v.BufferLines()[len(v.BufferLines())-1])[3:]
	if len(command) == 0 {
		newLine(g, v)
		return nil
	}
	if matched, _ := regexp.MatchString(commandRegex, string(command)); matched {
		if matched, _ := regexp.MatchString(commandGetUserInfoRegex, string(command)); matched {
			v.Clear()
			name, ip, port := MeerNode.User.GetUserInfo()
			info := fmt.Sprintf("%s %s:%s\n>> ", name, ip, port)
			v.Write([]byte(info))
			v.SetCursor(3, 1)
			return nil
		} else if matched, _ := regexp.MatchString(commandModifyNicknameRegex, string(command)); matched {
			v.Clear()
			name := strings.Split(string(command), " ")[2]
			MeerNode.User.ModifyNickname(name)

			views["userinfo"].Clear()
			name, ip, port := MeerNode.User.GetUserInfo()
			userInfo := fmt.Sprintf("%s\n%s\n:%s", name, ip, port)
			views["userinfo"].Write([]byte(userInfo))

			info := fmt.Sprintf("name changed : \"%s\"\n>> ", name)
			v.Write([]byte(info))
			v.SetCursor(3, 1)
			return nil
		} else if matched, _ := regexp.MatchString(commandCreateRoomRegex, string(command)); matched {
			v.Clear()
			room, _ := room.NewRoom("hihi", *MeerNode.User, "meer", "00000000", "163.152.3.135", "25258")
			MeerNode.Room = append(MeerNode.Room, *room)
			roomUpdateChan <- true

			info := fmt.Sprintf("create room : \"%s\"\n>> ", "hihi")
			v.Write([]byte(info))
			v.SetCursor(3, 1)
			return nil
		} else {
			v.Clear()
			errMsg := fmt.Sprintf("err : no such command \"%s\" \n>> ", command)
			v.Write([]byte(errMsg))
			v.SetCursor(3, 1)
			return nil
		}
	} else {
		v.Clear()
		errMsg := fmt.Sprintf("err : no such command \"%s\" \n>> ", command)
		v.Write([]byte(errMsg))
		v.SetCursor(3, 1)
		return nil
	}
	return nil
}

func sendMessage(g *gocui.Gui, v *gocui.View) error {
	if curMode != cmdMode {
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
	if err := g.SetKeybinding("cmdline", gocui.KeyEnter, gocui.ModNone, sendCommand); err != nil {
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