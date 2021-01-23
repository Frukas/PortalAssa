package main

import (
	context "context"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskMatchDB struct {
	folderID string `bson:"folderID"`
	TaskName string `bson:"TaskName"`
	Subject  string `bson:"Subject"`
	Body     string `bson:"Body"`
	Sender   string `bson:"Sender"`
	Priority int    `bson:"Priority"`
}

func getTaskName(mailInfo *MailSimple) *MailSimple {

	client := getDBClient()

	filter := bson.M{"FolderID": mailInfo.FolderID}
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "Priority", Value: -1}})

	collection := client.Database("PortalAssa").Collection("TaskMatch")

	res, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		log.Fatalf("erro na para conseguir a lista de regras de Task da pasta: %v erro: %v", mailInfo.GetFolderID(), err)
	}

	task := "Task not found"

	for res.Next(context.Background()) {
		flag := true
		var en taskMatchDB

		if err = res.Decode(&en); err != nil {
			log.Fatalf("Erro de ler elementos: %v", err)
		}

		if en.Subject != "na" {
			if !strings.Contains(mailInfo.GetSubject(), en.Subject+string(" ")) {
				flag = false
			}
		}

		if en.Body != "na" {
			if !strings.Contains(mailInfo.GetBodyPreview(), en.Body+string(" ")) {
				flag = false
			}
		}

		if en.Sender != "na" {
			if !strings.Contains(mailInfo.GetFromAddress(), en.Sender+string(" ")) {
				flag = false
			}
		}

		if flag {
			task = en.TaskName
		}
	}
	mailInfo.TaskName = task

	return mailInfo
}

func getDBClient() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Erro em conectar ao DB para verificação de task: %v/n", err)
	}
	return client
}
