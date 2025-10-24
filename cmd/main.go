package main

import (
	"golang/TodoList/conf"
	"golang/TodoList/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(":8080")
}
