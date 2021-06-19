package main

import (
	"container/list"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

//Client is a Struct that defines Client
type Client struct {
	ID       string
	Conn     *websocket.Conn
	Pool     *Pool
	Response *list.List
}

func (client *Client) Start() {
	go client.ping()
	go client.listeningMessage()
	go client.responsePoolClear()
}

func (client *Client) ping() {

	ticker := time.NewTicker(15 * time.Second)
	done := make(chan bool)
	fmt.Println("Ping Pong started", client.ID)

	for {
		select {
		case <-done:
			return

		case <-ticker.C:
			if err := client.Conn.WriteMessage(1, []byte("Ping")); err != nil {
				fmt.Println("Error sending  ping message: ", err)
				return
			}
			fmt.Println("Pingou")
		}
	}
}

func (client *Client) listeningMessage() {

	messages := list.New()

	//Starts a go routine that start to listing to the message and add it to a pool called messages
	go func(client *Client, messages *list.List) {
		for {
			_, message, err := client.Conn.ReadMessage()
			if err != nil {
				log.Fatal("Error listing Message on client: ", client.ID, "Error message: ", err)
			}
			if string(message) == "HeartBit Ok" {
				fmt.Println(string(message))

			} else {
				fmt.Println(string(message), "Json")
				client.Response.PushBack(message)
			}
		}
	}(client, messages)
}

func (client *Client) responsePoolClear() {
	client.Response = list.New()

	for {
		for client.Response.Len() > 0 {
			e := client.Response.Front()
			fmt.Println(string(e.Value.([]byte)))
			client.Response.Remove(e)
		}
	}
}
