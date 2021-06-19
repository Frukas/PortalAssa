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
			//go pool.ListeningMessageAndStartPing(client)
			client.Start()
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))

		case message := <-pool.Broadcast:

			//fmt.Println("Sending message to all clients in Pool")
			for client := range pool.Clients {

				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println("Error sending message: ", err)
					return
				}
			}
		}
	}
}

// func (pool *Pool) Ping(client *Client) {

// 	ticker := time.NewTicker(35 * time.Second)
// 	done := make(chan bool)
// 	fmt.Println("Chegou aqui", client.ID)

// 	for {
// 		select {
// 		case <-done:
// 			return

// 		case <-ticker.C:
// 			if err := client.Conn.WriteMessage(1, []byte("Ping")); err != nil {
// 				fmt.Println("Error sending  ping message: ", err)
// 				return
// 			}
// 			fmt.Println("Pingou")
// 		}
// 	}
// }

// func (pool *Pool) ListeningMessageAndStartPing(client *Client) {

// 	messages := list.New()

// 	//starts the ping every x seconds
// 	go pool.Ping(client)

// 	//Starts a go routine that start to listing to the message and add it to a pool called messages
// 	go func(client *Client, messages *list.List) {
// 		for {
// 			_, message, err := client.Conn.ReadMessage()
// 			if err != nil {
// 				log.Fatal("Error listing Message on client: ", client.ID, "Error message: ", err)
// 			}
// 			if string(message) == "HeartBit Ok" {
// 				fmt.Println(string(message))

// 			} else {
// 				fmt.Println(string(message), "Json")
// 			}
// 			messages.PushBack(message)
// 			go pool.PopFirstElementPool(messages)
// 		}
// 	}(client, messages)
// }

// func (pool *Pool) PopFirstElementPool(list *list.List) {

// 	if list.Front() != nil {
// 		fmt.Println(list.Front().Value, "Popping out")
// 		fmt.Println("Number of element:", list.Len())
// 		list.Remove(list.Front())
// 	}
// }
