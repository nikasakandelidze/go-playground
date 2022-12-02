package main

import (
	"clitranslator/src/controller"
	"clitranslator/src/presenter"
	"clitranslator/src/storage"
)

func main(){
	presenter := presenter.CliPresenter{} 
	storage := storage.InMemoryStorage{QuestionMap: make(map[string]storage.Message)}
	controller := controller.CliController{Storage: storage, Presenter: presenter}
	controller.Run()
}
/*
	Points to cover:
	1) how pointers work in golang: https://go101.org/article/pointer.html
	2) golang memory model https://go101.org/
	3) go lang data structures (primitives, strings, arrays, interfaces, structures and etc)
	4) go lang new keyword and when to use it
	5) when to use pointers vs values in golang
	5) for loops in golang
	6) how does equality work in go for strings, structs, ints, arrays, pointers
	7) what are types in golang? do they somehow relate to types in typescript ?
	8) golang structures
*/