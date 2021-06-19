package main

import (
	context "context"
	"fmt"
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
	http.ListenAndServe(":9090", nil)

}

func main() {
	go gRPCInit()
	webConnectorInit()
}
