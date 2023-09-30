package ws

import (
	"log"

	"github.com/Ahmad940/dropify/platform/hub"
	"github.com/gofiber/contrib/websocket"
)

func Stream(c *websocket.Conn) {
	// When the function returns, unregister the client and close the connection
	defer func() {
		hub.Unregister <- c
		c.Close()
	}()

	// Register the client
	hub.Register <- c

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}

			return // Calls the deferred function, i.e. closes the connection on error
		}

		if messageType == websocket.TextMessage {
			// Broadcast the received message
			hub.Broadcast <- string(message)
		} else {
			log.Println("websocket message received of type", messageType)
		}
	}
}
