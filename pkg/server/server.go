package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
)

const (
	meerModeMessage      = "You've got wrong password. Enter to Meerkat mode."
	paramRequiredMessage = "Some parameters are missing. Check if you enter all required params."
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Start(addr string) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverHandler(w, r)
	})
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	id, password, name, ok := getParamsFromUrl(r)
	if !ok {
		conn.WriteMessage(websocket.TextMessage, []byte(paramRequiredMessage))
		return
	}
	room, auth := getRoom(id, password)
	if !auth {
		conn.WriteMessage(websocket.TextMessage, []byte(meerModeMessage))
	}

	connInfo := newConnInfo(conn, auth, name)

	room.register(connInfo)
	defer room.unregister(connInfo)
	go room.sendMessage(connInfo)
	room.receiveMessage(connInfo)
}

func getParamsFromUrl(r *http.Request) (id string, password string, name string, exist bool) {
	id, exist = getParamFromQuery(r.URL.Query(), "id")
	if !exist {
		return
	}
	password, exist = getParamFromQuery(r.URL.Query(), "password")
	if !exist {
		return
	}
	name, exist = getParamFromQuery(r.URL.Query(), "name")
	if !exist {
		return
	}
	return
}

func getParamFromQuery(query url.Values, param string) (matchedParam string, exist bool) {
	params, exist := query[param]
	if !exist || len(params[0]) < 1 {
		log.Println("Url Param", param, "is missing")
		return matchedParam, exist
	}
	return params[0], exist
}
