package main

import (
	"github.com/sahasajib/mini_atm/cmd"
	"github.com/sahasajib/mini_atm/database"
)

func main() {
	database.InitDB()
	cmd.Serve()
}
