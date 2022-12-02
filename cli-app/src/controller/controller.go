package controller

import (
	"clitranslator/src/presenter"
	"clitranslator/src/storage"

	"strings"

	"github.com/google/uuid"
)


type Controller interface {
	Insert(message presenter.Message)
	GetMessageById(id string) presenter.Message
	GetAllMessages() []presenter.Message
	Run()
}

type CliController struct{
	Presenter presenter.Presenter
	Storage storage.Storage
}

func (c CliController) Insert(message presenter.Message) string {
	id := uuid.NewString()
	toInsertValue := storage.Message{Id: id, Content: message.Content, Metadata: message.Metadata}
	c.Storage.Insert(toInsertValue)
	return id
}

func (c CliController) GetMessageById(id string) (*presenter.Message, error) {
	val, err := c.Storage.GetMessageById(id)
	if err != nil {
		return nil, err
	}
 	return &presenter.Message{Content: val.Content, Metadata: val.Metadata}, nil
}

func (c CliController) GetAllMessages() []presenter.Message {
	storageMessages := c.Storage.GetAllMessages()
	result := make([]presenter.Message, len(storageMessages))
	for _, v := range storageMessages {
		result=append(result, presenter.Message{Content: v.Content, Metadata: v.Metadata})
	}
	return result
}


func (c CliController) Run(){
	c.Presenter.PresentMessage(presenter.Message{Content: "Welcome to cli V0.0.1"})
	c.Presenter.PresentMessage(presenter.Message{Content: "Please enter what tech news you'd like to see"})
	for {
		message := c.Presenter.GetMessage()
		words := strings.Fields(message.Content)
		if len(words) < 2 && len(words) > 0 && words[0] != "selectall" {
			c.Presenter.PresentMessage(presenter.Message{Content: "Invalid input command, please try again."})
		}else{
			message.Metadata = words[0]
			if len(words) > 1{
				message.Content = words[1]
			}
			if strings.Contains(message.Metadata, "insert") {
				id := c.Insert(message)
				c.Presenter.PresentMessage(presenter.Message{Content: "New element with id: "+ id + " was inserted"})
			} else if strings.Contains(message.Metadata, "selectall") {
				messages := c.GetAllMessages()
				for _, message := range messages {
					c.Presenter.PresentMessage(message)
				}
			} else if strings.Contains(message.Metadata, "select") {
				message, err := c.GetMessageById(message.Content)
				if err != nil {
					c.Presenter.PresentMessage(presenter.Message{Content: err.Error()})
				} else if message.Content != "" {
					c.Presenter.PresentMessage(presenter.Message{Content: message.Content})
				}
			} 
		}
	}	
}