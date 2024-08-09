package main

import (
	"database/sql"
	"fmt"

	"github.com/gabriel-hawerroth/capitech-back/configs"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/web"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/web/webserver"
	_ "github.com/lib/pq"
)

func main() {
	confs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(confs.DBDriver, getDbConnUrl(*confs))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	webServer := webserver.NewWebServer(confs.WebServerPort)
	web.LoadHandlers(webServer, db)
	webServer.Start()
}

func getDbConnUrl(confs configs.Conf) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", confs.DBUser, confs.DBPassword, confs.DBName, confs.DBHost, confs.DBPort)
}
