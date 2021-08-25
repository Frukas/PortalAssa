package main

import (
	"fmt"
)

//Pool is a pool of connection to the server
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan []Message
	Response   chan string
}

//NewPool initializes a new connection pool
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []Message),
		Response:   make(chan string),
	}
}

//Start initialize the pool
func (pool *Pool) Start() {

	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			client.Start()
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))

		case message := <-pool.Broadcast:

			for _, msg := range message {
				fmt.Println(msg)
			}

			for client := range pool.Clients {

				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println("Error sending message: ", err)
					return
				}
			}
		}
	}
}
