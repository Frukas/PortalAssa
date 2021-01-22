package main

import (
	"github.com/gorilla/websocket"
)

//Client is a Struct that defines Client
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

//Message is a Struct that defines Message
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
}
