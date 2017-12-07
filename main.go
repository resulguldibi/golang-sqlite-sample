package main

import (
	"rguldibi.com/golang-sqlite-sample/server"
)

func main() {
	server.NewServer().Run(":8080")
}
