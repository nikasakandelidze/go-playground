package store

import (
	"errors"
	"sync"
)

type Record struct {
	Value string
	Offset int
}

type Log struct {
	records []Record
	mutex sync.Mutex
}

func NewLog() *Log{
	return &Log{}
}


func (log *Log) Read(offset int) (Record, error) {
	log.mutex.Lock()
	defer log.mutex.Unlock()
	if offset >= len(log.records) || offset < 0 {
		return Record{}, errors.New("offset out of range") 
	}
	record := log.records[offset]
	return record, nil
}

func (log *Log) Append(record Record) (int, error) {
	log.mutex.Lock()
	defer log.mutex.Unlock()
	record.Offset = len(log.records)
	log.records = append(log.records, record)
	return record.Offset, nil
}

