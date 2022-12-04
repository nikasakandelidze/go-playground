package main

import (
	"commit-log/src/server"
	"fmt"
)

func main(){
	defer func () {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	server := server.NewServer()
	server.StartServer()
}