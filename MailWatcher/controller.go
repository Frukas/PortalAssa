package main

import (
	context "context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	grpc "google.golang.org/grpc"
)

var token = TokenHandler{
	Token:     "",
	NextToken: "",
	code:      "",
}

var folderInfos []FolderInfo

func folderInfosFromDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Erro em conectar ao DB: %v/n", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("PortalAssa").Collection("Folders")

	filter := bson.D{{}}

	res, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("erro na pesquisa: %v", err)
	}

	for res.Next(ctx) {
		var folder FolderDB
		if err = res.Decode(&folder); err != nil {
			log.Fatalf("Erro de ler elementos: %v", err)
		}

		folderInfos = append(folderInfos, FolderInfo{
			Name:             folder.FolderName,
			FolderID:         folder.FolderID,
			FolderAddress:    fmt.Sprint("https://graph.microsoft.com/v1.0/me/mailFolders/", folder.FolderID, "/messages/delta"),
			NextDeltaAddress: "",
			Token:            &token,
		})
	}
}

//TODO : separade the the folderAddress from the folderID
// var folderInfos = []FolderInfo{
// 	FolderInfo{
// 		Name:             "Teste",
// 		FolderID:         "AQMkADAwATY0MDABLTgwNTMtOTBiMi0wMAItMDAKAC4AAAMGPtl3uaKoToiqB4aFOd24AQCh0MdHQYn1SrJTIf58X13eAANF1YVIAAAA",
// 		FolderAddress:    "https://graph.microsoft.com/v1.0/me/mailFolders/AQMkADAwATY0MDABLTgwNTMtOTBiMi0wMAItMDAKAC4AAAMGPtl3uaKoToiqB4aFOd24AQCh0MdHQYn1SrJTIf58X13eAANF1YVIAAAA/messages/delta",
// 		NextDeltaAddress: "",
// 		Token:            &token,
// 	},
// }

func getNewMailsServiceClient() MailsServiceClient {
	fmt.Println("Hello I´m a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	//defer conn.Close()

	return NewMailsServiceClient(conn)
}

//ControllerStart starts the chain of events tokenHandler and mailWatcher
func ControllerStart(code string) {

	token.code = code

	c := getNewMailsServiceClient()

	token.GetTokenFromCode()

	folderInfosFromDB()

	for _, folder := range folderInfos {
		fmt.Println(folder.Name)
		fmt.Println(folder.FolderAddress)
		go folder.InitialDeltaMailrequest(folder.FolderAddress, c)
	}
}
