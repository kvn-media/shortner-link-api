package main

import (
	"github.com/kvn-media/link-shortner-api/database"
	"github.com/kvn-media/link-shortner-api/server"
	"github.com/kvn-media/link-shortner-api/utils"
)

func main() {
	database.StartDatabase()
	server.Start(utils.GetConfig())
}
