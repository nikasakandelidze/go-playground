package presenter

import (
	"bufio"
	"fmt"
	"os"
)

type Message struct {
	Content string
	Metadata string
}

type Presenter interface {
	PresentMessage(Message Message)
	GetMessage() Message
}

type CliPresenter struct {}

func (presenter CliPresenter) PresentMessage(message Message){
	// if message.Metadata != ""{
	// 	fmt.Println(message.Metadata)	
	// }
	fmt.Println(message.Content)
}

func (presenter CliPresenter) GetMessage() Message {
	var input *Message = new(Message)// we can use pointers to some addresses like in c
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	input.Content=text
	return *input // we can use dereferensing just like in c
}