package storage


type Message struct{
	Id string
	Content string
	Metadata string
}

type Storage interface {
	Insert(message Message)
	GetMessageById(id string) (*Message, error)
	GetAllMessages() []Message
}

type InMemoryStorage struct {
	QuestionMap map[string] Message
}

/* 
	Here is the custom error definition in golang for storage
*/
type NoSuchElementErr struct{
	Id string
}

func (err NoSuchElementErr) Error() string {
	return "No element with id:" + err.Id +" found in storage."
}


func (s InMemoryStorage) Insert(message Message){
	s.QuestionMap[message.Id]=message
}

func (s InMemoryStorage) GetMessageById(id string) (*Message, error) {
	if val ,ok := s.QuestionMap[id]; ok {
		return &val, nil
	}else{
		return nil, NoSuchElementErr{Id:id}
	}
}

func (s InMemoryStorage) GetAllMessages() []Message {
	messages := make([]Message, len(s.QuestionMap))
	for _, value := range s.QuestionMap {
		messages = append(messages, value)
	}
	return messages
}