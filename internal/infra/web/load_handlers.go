package web

import (
	"database/sql"
	"fmt"

	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/web/handlers"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/web/webserver"
	"github.com/gabriel-hawerroth/capitech-back/internal/services"
)

var server *webserver.WebServer
var db *sql.DB

func LoadHandlers(webServer *webserver.WebServer, dbConn *sql.DB) {
	server = webServer
	db = dbConn

	loadProductHandlers()
}

func loadProductHandlers() {
	const productPath = "/product"

	repository := database.NewProductRepository(db)
	service := services.NewProductService(*repository)
	handler := handlers.NewProductHandler(*service)

	server.AddHandler(getMapping(productPath), handler.GetProductsList)
}

func getMapping(path string) string {
	return fmt.Sprint("GET %s", path)
}

func postMapping(path string) string {
	return fmt.Sprint("POST %s", path)
}

func putMapping(path string) string {
	return fmt.Sprint("PUT %s", path)
}

func deleteMapping(path string) string {
	return fmt.Sprint("DELETE %s", path)
}
