package main

import (
	"github.com/iikmaulana/example-code/base/config"
	"github.com/joho/godotenv"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/logger"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg := config.Init()

	serr := cfg.Start()

	err = cfg.Server.Start()
	if serr != nil {
		logger.Panic(serr)
	}
}
