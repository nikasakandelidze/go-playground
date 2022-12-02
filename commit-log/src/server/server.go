package server

import (
	"commit-log/src/store"
	"encoding/json"
	"net/http"
)

type HttpHandlerFunction = func(responseWriter http.ResponseWriter, request *http.Request);

type Server struct {
	log *store.Log
}

type Message struct {
	Message string
}

type AppendLogRequest struct {
	Record store.Record
}

type AppendLogResponse struct {
	Offset int
}

type GetLogRequest struct {
	Offset int
}

type GetLogResponse struct {
	Record store.Record
}

func NewServer() *Server {
	return &Server{ log: store.NewLog() }
}

func (s *Server) handleAppendLog() HttpHandlerFunction {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method=="POST" {
			logRequest := &AppendLogRequest{}
			json.NewDecoder(r.Body).Decode(logRequest)
			result, err := s.log.Append(logRequest.Record)
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Message{Message: err.Error()})
				return
			}
			json.NewEncoder(w).Encode(AppendLogResponse{Offset: result})
		}
	}
}

func (s *Server) handleReadLog() HttpHandlerFunction {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method=="GET" {
			getLogRequest := &GetLogRequest{}
			json.NewDecoder(r.Body).Decode(getLogRequest)
			result, err := s.log.Read(getLogRequest.Offset)
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Message{Message: err.Error()})
				return
			}
			json.NewEncoder(w).Encode(GetLogResponse{Record: result})
		}
	}
}

func (s *Server) StartServer(){
	http.HandleFunc("/append", s.handleAppendLog())
	http.HandleFunc("/read", s.handleReadLog())
	http.ListenAndServe(":3000", nil)
}