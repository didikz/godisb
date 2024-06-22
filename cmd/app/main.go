package main

import (
	"github.com/didikz/godisb/api"
)

func main() {
	srv := api.NewHttpServer(":8080")
	err := srv.Run()
	if err != nil {
		panic(err)
	}
}
