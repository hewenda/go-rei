package main

import (
	"hewenda/go-rei/cmd"
	"hewenda/go-rei/storage"
)

func main() {
	storage.OpenDatabase()

	cmd.Execute()
}
