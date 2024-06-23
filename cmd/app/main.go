package main

import (
	"fmt"

	"github.com/didikz/godisb/api"
	"github.com/didikz/godisb/config"
)

type User struct {
	ID      int64  `db:"id"`
	Email   string `db:"email"`
	Balance int64  `db:"balance"`
}

func main() {
	config := config.Load("./")
	fmt.Println("config", config)
	srv := api.NewHttpServer(*config)
	err := srv.Run()
	if err != nil {
		panic(err)
	}
}
