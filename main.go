package main

import (
	"github.com/kvn-media/shortner-link-api/database"
	"github.com/kvn-media/shortner-link-api/server"
	"github.com/kvn-media/shortner-link-api/utils"
)

func main() {
	database.StartDatabase()
	server.Start(utils.GetConfig())
}
