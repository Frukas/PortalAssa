package main

import "go.mongodb.org/mongo-driver/bson/primitive"

//Message is a Struct that defines Message
type Message struct {
	ID            primitive.ObjectID `json:"id"`
	Abreviation   string             `json:"Abreviation"`
	Operator      string             `json:"Operator"`
	DoubleChecker string             `json:"DoubleChecker"`
	TaskName      string             `json:"TaskName"`
	Status        string             `json:"Status"`
	Date          string             `json:"Date"`
}

type messageHandler struct {
	Messages []Message
}

func (m *messageHandler) addTaskEntryToArray(task TaskEntry) {
	newMessage := Message{
		ID: task.ID,
		//Abreviation:   task.MailInfo.FolderID,
		Abreviation:   task.Abreviation,
		Operator:      task.Operator,
		DoubleChecker: task.DoubleChecker,
		TaskName:      task.MailInfo.TaskName,
		Status:        task.Status,
		Date:          task.MailInfo.ReceivedDateTime,
	}

	for _, abrev := range getAbreviationFromFolder(task.MailInfo.FolderID) {
		newMessage.Abreviation = abrev.Abreviation
	}

	m.Messages = append(m.Messages, newMessage)
}
