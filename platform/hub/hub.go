package hub

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

// Add more data to this type if needed
type Client struct {
	isClosing bool
	mu        sync.Mutex
}

var Clients = make(map[*websocket.Conn]*Client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
var Register = make(chan *websocket.Conn)
var Broadcast = make(chan string)
var Unregister = make(chan *websocket.Conn)

func RunHub() {
	for {
		select {
		case connection := <-Register:
			Clients[connection] = &Client{}
			log.Println("connection registered")

		case message := <-Broadcast:
			log.Println("message received:", message)
			// Send the message to all clients
			for connection, c := range Clients {
				go func(connection *websocket.Conn, c *Client) { // send to each client in parallel so we don't block on a slow client
					c.mu.Lock()
					defer c.mu.Unlock()
					if c.isClosing {
						return
					}
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.isClosing = true
						log.Println("write error:", err)

						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						Unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-Unregister:
			// Remove the client from the hub
			delete(Clients, connection)

			log.Println("connection unregistered")
		}
	}
}
