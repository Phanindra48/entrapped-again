package entrapped

import (
	"github.com/kgthegreat/entrapped-again/Godeps/_workspace/src/github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// time allowed to write a message to the peer
	writeWait = time.Second
	// time allowed to read the next pong message from the peer
	pongWait = 10 * time.Second
	// send pings to peer with this period, must be less than pongWait
	pingPeriod = (pongWait * 9) / 10
	// maximum message size allowed from peer
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
    CheckOrigin:     checkOrigin,
}

func checkOrigin(req *http.Request) bool {
	return true
}

type trooper struct {
	nickname string
	trap     *trap
	ws       *websocket.Conn
	data     chan []byte
}

type message struct {
	t       *trooper
	msg     string
	msgType int
}

func (t *trooper) readPump() {
	defer func() {
		t.dead()
		t.ws.Close()
	}()

	t.ws.SetReadLimit(maxMessageSize)
	t.ws.SetReadDeadline(time.Now().Add(pongWait))
	t.ws.SetPongHandler(func(string) error {
		t.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		msgType, msg, err := t.ws.ReadMessage()
		if err != nil {
			break
		}

		ch.message <- &message{t, string(msg), msgType}
	}
}

// writePump pumps messages to connection
func (t *trooper) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		t.ws.Close()
	}()

	for {
		select {
		case message, ok := <-t.data:
			if !ok {
				t.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := t.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := t.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// write payload to the connection
func (t *trooper) write(mt int, payload []byte) error {
	t.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return t.ws.WriteMessage(mt, payload)
}

// dead used to send close signal
func (t *trooper) dead() {
	ch.dead <- t
}
