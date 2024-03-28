package multiplayer

import (
	"cp23kk1/common/hub"
	"fmt"
	"net/http"
	"strconv"
)

func ServeWs(w http.ResponseWriter, r *http.Request, isCreate bool) {
	fmt.Println("service")
	ws, err := hub.Upgrader.Upgrade(w, r, nil)
	//Get room's id from client...
	queryValues := r.URL.Query()
	roomId := queryValues.Get("roomId")
	roomName := queryValues.Get("roomName")
	numberPlayer, str_err := strconv.Atoi(queryValues.Get("numberPlayer"))
	fmt.Println("service18")

	room := hub.H.Rooms[roomId]
	if !isCreate && room.Name == "" {
		ws.Close()

	}
	if err != nil || str_err != nil {
		c := err
		if c != nil {
			c = str_err
		}
		ws.Close()
	}
	// ws.ReadJSON(&roomId)

	c := &hub.Connection{Send: make(chan []byte, 256), Ws: ws}
	fmt.Println("service35")

	s := hub.Subscription{Conn: c, RoomId: roomId, RoomName: roomName, NumberPlayer: numberPlayer}

	if isCreate {
		hub.H.CreateLobby <- s
	} else {
		hub.H.Register <- s

	}
	go s.WritePump()
	go s.ReadPump()

}
