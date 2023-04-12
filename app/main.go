package main

import (
	"github.com/gjustoo/Anigo/controller"
	"github.com/gjustoo/Anigo/model"
)

func main() {
	model.Init()
	model.GetAllPosts()
	controller.Start()
}
