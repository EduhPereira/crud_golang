package main

import (
	"github.com/EduhPereira/crud_golang/database"
	"github.com/EduhPereira/crud_golang/server"
)

func main(){
	database.StartDB()

	server := server.NewServer()

	server.Run()
}