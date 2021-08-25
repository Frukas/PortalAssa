package main

import (
	context "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	grpc "google.golang.org/grpc"
)

type server struct{}

var pool = NewPool()

func gRPCInit() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterMailsServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
	fmt.Println("gRPC Started")

}

func (s *server) PostMails(ctx context.Context, req *MailsRequest) (*MailsResponse, error) {

	for _, mail := range req.GetMails().GetMails() {

		var messages messageHandler

		saveMailintoDB(mail)
		folderElments := getFolderListElement()

		for _, message := range folderElments {
			messages.addTaskEntryToArray(message)
		}

		pool.Broadcast <- messages.Messages

	}
	res := &MailsResponse{
		Result: "Message received",
	}
	return res, nil
}

func serverWS(pool *Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit")
	conn, err := Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	//client.Read()
}

func webConnectorInit() {

	go pool.Start()

	fmt.Println("Started the server on the port 9090")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWS(pool, w, r)
	})

	//return the initial task list
	http.HandleFunc("/openTaskList", func(w http.ResponseWriter, r *http.Request) {

		folderElments := getFolderListElement()

		var messages messageHandler

		for _, message := range folderElments {
			messages.addTaskEntryToArray(message)
		}

		if err := json.NewEncoder(w).Encode(messages); err != nil {
			log.Println("Error sending the initial task list: ", err)
		}
	})

	http.HandleFunc("/updateTask", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(reqBody))
	})

	http.ListenAndServe(":9090", nil)

}

func main() {
	go gRPCInit()
	webConnectorInit()
}
