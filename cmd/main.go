package main

import (
	"flag"
	"log"
	"thaibev_backend/appconfig"
	"thaibev_backend/cmd/routes"
	"thaibev_backend/config"
	"time"
)

func main() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatalf("error loading location 'Asia/Bangkok': %v\n", err)
	}
	time.Local = ict
	log.Printf("Local time zone %v", time.Now().In(ict))

	cfg := config.LoadFileConfig[appconfig.AppConfig](*configPath)

	e := routes.ServerStart(cfg)
	routes.ServerShutdown(e)

}
