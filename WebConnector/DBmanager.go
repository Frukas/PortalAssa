package main

import (
	context "context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//TaskEntry is how the mail will be saved on the monogoDB
type TaskEntry struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	MailInfo       *MailSimple        `bson:"MailInfo"`
	StartTime      time.Time          `bson:"StartTime,omitempty"`
	EndTime        time.Time          `bson:"EndTime,omitempty"`
	Operator       string             `bson:"Operator,omitempty"`
	DoubleChecker  string             `bson:"DoubleChecker,omitempty"`
	ActualTaskTime int                `bson:"ActualTaskTime,omitempty"`
	ReportTaskTime int                `bson:"ReportTaskTime,omitempty"`
	Status         string             `bson:"Status,omitempty"`
	Abreviation    string             `bson:"Abreviation,omitempty"`
}

//FolderStruct is the struct of the collection "Folder" in the DB
type FolderStruct struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	Client         string             `bson:"Client, omitempty"`
	FolderName     string             `bson:"FolderName, omitempty"`
	ParentFolderID string             `bson:"ParentFolderID, omitempty"`
	FolderID       string             `bson:"FolderID, omitempty"`
	Abreviation    string             `bson:"Abreviation, omitempty"`
	Color          string             `bson:"Color, omitempty"`
}

func getDBClient() *mongo.Client {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Erro em conectar ao DB: %v/n", err)
	}

	return client
}

func getTaskListCollection() *mongo.Collection {
	collection := getDBClient().Database("PortalAssa").Collection("TaskList")

	return collection
}

func getTaskArchiveCollection() *mongo.Collection {

	collection := getDBClient().Database("PortalAssa").Collection("TaskArchive")

	return collection
}

func getFoldersCollection() *mongo.Collection {
	collection := getDBClient().Database("PortalAssa").Collection("Folders")

	return collection
}

func saveMailintoDB(mail *MailSimple) {

	entry := bson.D{
		primitive.E{Key: "MailInfo", Value: mail},
		primitive.E{Key: "StartTime", Value: time.Time{}},
		primitive.E{Key: "EndTime", Value: time.Time{}},
		primitive.E{Key: "Operator", Value: "Operator"},
		primitive.E{Key: "DoubleChecker", Value: "DoubleChecker"},
		primitive.E{Key: "ActualTaskTime", Value: 0},
		primitive.E{Key: "ReportTaskTime", Value: 0},
		primitive.E{Key: "Status", Value: "Open"},
		primitive.E{Key: "Abreviation", Value: "na"},
	}

	if _, err := getTaskListCollection().InsertOne(context.Background(), entry); err != nil {
		log.Println("Fail to insert mail registry in the DB: ", err)
	}
}

func getFolderListElement() []TaskEntry {

	filter := bson.M{}

	var list []TaskEntry

	res, err := getTaskListCollection().Find(context.Background(), filter)
	if err != nil {
		log.Println("erro na pesquisa: ", err)
	}

	if err = res.All(context.Background(), &list); err != nil {
		log.Println("Failed to get the task lists: ", err)
	}
	return list
}

func getAbreviationFromFolder(Folder string) []FolderStruct {

	var list []FolderStruct

	filter := bson.D{
		primitive.E{Key: "FolderID", Value: Folder},
	}

	res, err := getFoldersCollection().Find(context.Background(), filter)
	if err != nil {
		log.Println("Error searching for the abreviation: ", err)
	}

	if err = res.All(context.Background(), &list); err != nil {
		log.Println("Failed to get the task lists: ", err)
	}

	return list
}
