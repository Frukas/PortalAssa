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
		fmt.Println(mail.GetSubject())
		fmt.Println(mail.GetTaskName())
		message := Message{
			Type: 1,
			Body: mail.GetTaskName(),
		}

		pool.Broadcast <- message

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

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWS(pool, w, r)
	})
	http.ListenAndServe(":9090", nil)
}

func main() {
	go gRPCInit()
	webConnectorInit()
}
