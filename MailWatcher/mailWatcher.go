package main

import (
	context "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//FolderInfo is the basic information of each folder
type FolderInfo struct {
	Name             string
	FolderID         string
	FolderAddress    string
	NextDeltaAddress string
	Token            *TokenHandler
}

//InitialDeltaMailrequest is responsible to start the chain
func (f *FolderInfo) InitialDeltaMailrequest(urlAddress string, c MailsServiceClient) {

	var deltaMessage DeltaMessage
	var deltaNextMessage DeltaNext

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlAddress, nil)
	if err != nil {
		log.Fatalf("First query: Error in the new request: %d", err)
	}

	finalToken := "Bearer " + f.Token.Token

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", finalToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("First query: error executing the request %d", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("First query: Error when reading the body")
	}

	if err := json.Unmarshal([]byte(body), &deltaMessage); err != nil {
		log.Fatalf("First query: error converting to JSON: %d", err)
	}

	if deltaMessage.OdataNextLink != "" {
		f.InitialDeltaMailrequest(deltaMessage.OdataNextLink, c)
	} else {
		if err := json.Unmarshal([]byte(body), &deltaNextMessage); err != nil {
			log.Fatalf("First query: error converting to JSON: %d", err)
		}
		//fmt.Println("next Delta Link: ", deltaNextMessage.OdataDeltaLink)
		f.DeltaMessageWatcher(deltaNextMessage.OdataDeltaLink, c)
	}
}

//DeltaMessageWatcher verifies the messagem from the second time on
func (f *FolderInfo) DeltaMessageWatcher(urlMessage string, c MailsServiceClient) {
	var deltaNextMessage DeltaNext
	mails := Mails{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlMessage, nil)
	if err != nil {
		log.Fatalf("Next query: Error in the new request: %d", err)
	}

	finalToken := "Bearer " + f.Token.Token
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", finalToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(fmt.Sprintf("Next query: error executing the request %d", err))
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Next query: Error when reading the body")
	}

	if err := json.Unmarshal([]byte(body), &deltaNextMessage); err != nil {
		log.Fatalf("Next query: error converting to JSON: %d", err)
	}

	time.Sleep(5 * time.Second)

	//fmt.Println("next address: ", deltaNextMessage.OdataDeltaLink)
	for _, value := range deltaNextMessage.Value {
		mails.Mails = append(mails.Mails, f.convertToProto3(value))
		fmt.Println(mails.Mails[0].TaskName)
		fmt.Println(value.Subject)
	}

	mreq := &MailsRequest{
		Mails: &mails,
	}
	if mails.Mails != nil {

		resp, err := c.PostMails(context.Background(), mreq)
		if err != nil {
			log.Fatal("Erro de envio gRPC:", err)
		}
		fmt.Println("Result of the calling: ", resp.GetResult())

	}

	f.DeltaMessageWatcher(deltaNextMessage.OdataDeltaLink, c)

}

func (f *FolderInfo) convertToProto3(msg Value) *MailSimple {
	sm := &MailSimple{
		ExchangeID:       msg.ConversationID,
		ReceivedDateTime: msg.ReceivedDateTime.String(),
		Subject:          msg.Subject,
		BodyPreview:      msg.BodyPreview,
		WebLink:          msg.WebLink,
		BodyContentType:  msg.Body.ContentType,
		BodyContent:      msg.Body.Content,
		FromAddress:      msg.From.EmailAddress.Address,
		ParentFolderId:   msg.ParentFolderID,
		TaskName:         "",
		FolderID:         f.FolderID,
	}

	return getTaskName(sm)
}
