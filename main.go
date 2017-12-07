package main

import (
	"resulguldibi/golang-sqlite-sample/server"
)

func main() {
	server.NewServer().Run(":8080")
}
