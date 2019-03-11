package main

import (
	"time"
	"todolist-mongodb/controllers"
	"todolist-mongodb/models"
)

func Sleep() {
	time.Sleep(time.Second * 1)
}

func main() {
	controller := controllers.Controller{}
	var step1 = models.Todo{ID: 1, Title: "Todo", Description: "Use Go and Mongodb"}
	var step1A = models.Todo{ID: 1, Title: "Todo", Description: "Use Go, Mongodb & Dep"}
	var step2 = models.Todo{ID: 2, Title: "Push to Github", Description: "Push base bare bones code to github"}
	var step3 = models.Todo{ID: 3, Title: "Test controllers", Description: "Make a todolist_test.go file, writes test for each controller."}

	controller.InsertOne(step1)
	Sleep()
	controller.InsertOne(step2)
	Sleep()
	controller.InsertOne(step3)
	Sleep()
	controller.UpdateOne(step1A)
	Sleep()
	controller.DeleteOne(step1)
	Sleep()
	controller.Find(100)
	Sleep()
	controller.DeleteAll()
}
