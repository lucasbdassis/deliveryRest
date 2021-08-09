package main

import (
	"delivery/database"
	"delivery/server"
	"fmt"
)

func main(){
	go database.RunDb()

	serv := server.NewServer()

	x :=server.GetApplication()

	fmt.Println(x.Services.Db.Ports[0])
	serv.Run()
}

