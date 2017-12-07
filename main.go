package main

import (
	"rguldibi.com/SQLiteDemo/server"
)

func main() {
	server.NewServer().Run(":8080")
}
