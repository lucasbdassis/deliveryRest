package main

import (
	"delivery/database"
	"delivery/server"
)

func main() {
	serv := server.NewServer()
	Started(serv)
}

func Started(serv server.Server) {
	go database.RunDb(server.GetPort())
	serv.Run()

}
