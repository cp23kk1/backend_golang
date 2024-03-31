package hub

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (

	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2048
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Subscription struct {
	Conn         *Connection
	RoomId       string
	RoomName     string
	NumberPlayer int
}

type Connection struct {
	// The websocket Connection.
	Ws *websocket.Conn
	// Buffered channel of outbound messages.
	Send chan []byte
}

func (s *Subscription) ReadPump() {
	c := s.Conn
	defer func() {
		//Unregister
		H.Unregister <- *s
		c.Ws.Close()
	}()
	c.Ws.SetReadLimit(maxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		//Reading incoming message...
		_, msg, err := c.Ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		m := message{s.RoomId, msg}

		//send message

		H.Broadcast <- m
	}
}
func (s *Subscription) WritePump() {
	c := s.Conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()
	for {
		select {
		//Listerning message when it comes will write it into writer and then send it to the client
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				fmt.Println(err.Error())
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
func (c *Connection) write(mt int, payload []byte) error {

	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Ws.WriteMessage(mt, payload)
}

// Hub maintains the set of active clients and broadcasts messages to the
type Hub struct {
	// put registered clients into the room.
	Rooms map[string]TRoom
	// Inbound messages from the clients.
	Broadcast chan message

	// Register requests from the clients.
	Register chan Subscription

	// Create Lobby and Register requests from the clients.
	CreateLobby chan Subscription

	// Unregister requests from clients.
	Unregister chan Subscription
}

type message struct {
	Room string `json:"room"`
	Data []byte `json:"data"`
}
type Lobby struct {
	RoomId       string `json:"roomId"`
	Name         string `json:"name"`
	NumberPlayer int    `json:"numberPlayer"`
	MaxPlayer    int    `json:"maxPlayer"`
}
type TRoom struct {
	Name        string
	MaxPlayer   int
	Connections map[*Connection]bool
}

var H = &Hub{
	Broadcast:   make(chan message),
	Register:    make(chan Subscription),
	CreateLobby: make(chan Subscription),
	Unregister:  make(chan Subscription),
	Rooms:       make(map[string]TRoom),
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.Register:

			if len(h.Rooms[s.RoomId].Connections) == h.Rooms[s.RoomId].MaxPlayer {

				// delete(Connections, s.Conn)
				close(s.Conn.Send)

			} else {
				h.Rooms[s.RoomId].Connections[s.Conn] = true
				// s.Conn.Send <- []byte{}
			}
		case s := <-h.CreateLobby:
			room := h.Rooms[s.RoomId]

			//create room
			if room.Name == "" {
				room.Connections = make(map[*Connection]bool)
				room.Name = s.RoomName
				room.MaxPlayer = s.NumberPlayer
				h.Rooms[s.RoomId] = room
			}

			h.Rooms[s.RoomId].Connections[s.Conn] = true
			// s.Conn.Send <- []byte{}

		case s := <-h.Unregister:
			room := h.Rooms[s.RoomId]
			if room.Name != "" {
				if _, ok := room.Connections[s.Conn]; ok {
					delete(room.Connections, s.Conn)
					close(s.Conn.Send)
					if len(room.Connections) == 0 {
						delete(h.Rooms, s.RoomId)
					}
				}
			}
		case m := <-h.Broadcast:
			room := h.Rooms[m.Room]
			for c := range room.Connections {
				select {
				case c.Send <- m.Data:
				default:
					close(c.Send)
					delete(room.Connections, c)
					// if len(room.Connections) == 0 {
					// 	delete(h.Rooms, m.Room)
					// }
				}
			}
		}
	}
}

func (h *Hub) GetLobby() []Lobby {
	result := []Lobby{}
	for i, s := range h.Rooms {
		fmt.Print(s)
		if len(s.Connections) < s.MaxPlayer {
			result = append(result, Lobby{RoomId: i, Name: h.Rooms[i].Name, NumberPlayer: len(h.Rooms[i].Connections), MaxPlayer: h.Rooms[i].MaxPlayer})
		}
	}
	return result
}
