package multiplayer

import (
	"cp23kk1/common/hub"
	"fmt"
	"net/http"
	"strconv"
)

func ServeWs(w http.ResponseWriter, r *http.Request, isCreate bool) {
	ws, err := hub.Upgrader.Upgrade(w, r, nil)
	//Get room's id from client...
	queryValues := r.URL.Query()
	roomId := queryValues.Get("roomId")
	roomName := queryValues.Get("roomName")
	numberPlayer, str_err := strconv.Atoi(queryValues.Get("numberPlayer"))

	room := hub.H.Rooms[roomId]
	if !isCreate && room.Name == "" {
		fmt.Println("service20")

		ws.Close()
		return
	}
	fmt.Println("service25")

	if err != nil || str_err != nil {
		c := err
		if c != nil {
			c = str_err
		}
		ws.Close()
		return
	}
	fmt.Println("service25")

	// ws.ReadJSON(&roomId)

	c := &hub.Connection{Send: make(chan []byte, 256), Ws: ws}

	s := hub.Subscription{Conn: c, RoomId: roomId, RoomName: roomName, NumberPlayer: numberPlayer}

	if isCreate {
		hub.H.CreateLobby <- s
	} else {
		hub.H.Register <- s

	}
	go s.WritePump()
	go s.ReadPump()

}

func GetlobbyService() []hub.Lobby {
	return hub.H.GetLobby()
}
