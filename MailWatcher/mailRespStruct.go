package main

import "time"

//FolderDB is the struct for the Folder collection in the DataBase
type FolderDB struct {
	client         string `bson:"Client"`
	FolderName     string `bson:"FolderName"`
	ParentFolderID string `bson:"ParentFolderID"`
	FolderID       string `bson:"FolderID"`
}

//DeltaMessage struture received from the first delta message query
type DeltaMessage struct {
	OdataContext  string  `json:"@odata.context"`
	OdataNextLink string  `json:"@odata.nextLink"`
	Value         []Value `json:"value"`
}

//DeltaNext from a second delta query on
type DeltaNext struct {
	OdataContext   string  `json:"@odata.context"`
	OdataDeltaLink string  `json:"@odata.deltaLink"`
	Value          []Value `json:"value"`
}

//Value is a basic message structure
type Value struct {
	OdataType                  string        `json:"@odata.type"`
	OdataEtag                  string        `json:"@odata.etag"`
	CreatedDateTime            time.Time     `json:"createdDateTime"`
	LastModifiedDateTime       time.Time     `json:"lastModifiedDateTime"`
	ChangeKey                  string        `json:"changeKey"`
	Categories                 []interface{} `json:"categories"`
	ReceivedDateTime           time.Time     `json:"receivedDateTime"`
	SentDateTime               time.Time     `json:"sentDateTime"`
	HasAttachments             bool          `json:"hasAttachments"`
	InternetMessageID          string        `json:"internetMessageId"`
	Subject                    string        `json:"subject"`
	BodyPreview                string        `json:"bodyPreview"`
	Importance                 string        `json:"importance"`
	ParentFolderID             string        `json:"parentFolderId"`
	ConversationID             string        `json:"conversationId"`
	ConversationIndex          string        `json:"conversationIndex"`
	IsDeliveryReceiptRequested interface{}   `json:"isDeliveryReceiptRequested"`
	IsReadReceiptRequested     bool          `json:"isReadReceiptRequested"`
	IsRead                     bool          `json:"isRead"`
	IsDraft                    bool          `json:"isDraft"`
	WebLink                    string        `json:"webLink"`
	InferenceClassification    string        `json:"inferenceClassification"`
	ID                         string        `json:"id"`
	Body                       struct {
		ContentType string `json:"contentType"`
		Content     string `json:"content"`
	} `json:"body"`
	Sender struct {
		EmailAddress struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"emailAddress"`
	} `json:"sender"`
	From struct {
		EmailAddress struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"emailAddress"`
	} `json:"from"`
	ToRecipients []struct {
		EmailAddress struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"emailAddress"`
	} `json:"toRecipients"`
	CcRecipients  []interface{} `json:"ccRecipients"`
	BccRecipients []interface{} `json:"bccRecipients"`
	ReplyTo       []interface{} `json:"replyTo"`
	Flag          struct {
		FlagStatus string `json:"flagStatus"`
	} `json:"flag"`
}
