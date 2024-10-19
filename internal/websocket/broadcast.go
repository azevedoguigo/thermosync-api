package websocket

import "log"

var broadcast = make(chan Message)

func HandleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("Error to send message:", err.Error())
				client.Close()
				delete(clients, client)
			}
		}
	}
}
