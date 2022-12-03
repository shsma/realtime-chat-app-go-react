package websocket

import (
	"log"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Printf("Register - Size of Connection Pool: %v", len(pool.Clients))
			for client, _ := range pool.Clients {
				log.Println(client)
				err := client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
				if err != nil {
					log.Println(err)
				}
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Printf("Unregister - Size of Connection Pool: %v", len(pool.Clients))
			for client, _ := range pool.Clients {
				err := client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
				if err != nil {
					log.Println(err)
				}
			}
			break
		case message := <-pool.Broadcast:
			log.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}