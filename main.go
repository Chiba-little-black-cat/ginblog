package main

import (
	"ginblog/model"
	"ginblog/router"
)

func main() {
	model.InitDb()
	router.InitRouter()
}
