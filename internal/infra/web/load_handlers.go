package web

import (
	"database/sql"
	"fmt"
	"net/http"

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

	loadCategoryHandlers()
	loadShoppingCartHandlers()
	loadProductHandlers()

	server.AddHandler("/hello-world", helloWorld)
}

func loadCategoryHandlers() {
	const basePath = "/category"

	repository := database.NewCategoryRepository(db)
	service := services.NewCategoryService(*repository)
	handler := handlers.NewCategoryHandler(*service)

	server.AddHandler(getMapping(basePath), handler.GetCategoriesList)
}

func loadShoppingCartHandlers() {
	const basePath = "/shopping-cart"

	repository := database.NewShoppingCartRepository(db)
	service := services.NewShoppingCartService(*repository)
	handler := handlers.NewShoppingCartHandler(*service)

	server.AddHandler(postMapping(basePath), handler.AddProduct)
	server.AddHandler(getMapping(basePath)+"/getUserShoppingCart", handler.GetUserShoppingCart)
}

func loadProductHandlers() {
	const basePath = "/product"

	repository := database.NewProductRepository(db)
	service := services.NewProductService(*repository)
	handler := handlers.NewProductHandler(*service)

	server.AddHandler(getMapping(basePath)+"/{id}", handler.GetById)
	server.AddHandler(getMapping(basePath), handler.GetProductsList)
	server.AddHandler(getMapping(basePath)+"/getTrendingProductsList", handler.GetTrendingProducts)
	server.AddHandler(getMapping(basePath)+"/getBestSellingProductsList", handler.GetBestSellingProducts)
	server.AddHandler(getMapping(basePath)+"/getUserSearchHistory", handler.GetUserSearchHistory)
	server.AddHandler(postMapping(basePath), handler.Save)
	server.AddHandler(putMapping(basePath)+"/{id}", handler.Edit)
	server.AddHandler(patchMapping(basePath)+"/editProductPrice/{id}", handler.ChangePrice)
	server.AddHandler(patchMapping(basePath)+"/editProductStockQuantity/{id}", handler.ChangeStockQuantity)
	server.AddHandler(patchMapping(basePath)+"/changeProductImage/{id}", handler.ChangeImage)
	server.AddHandler(patchMapping(basePath)+"/removeProductImage/{id}", handler.RemoveImage)
	server.AddHandler(deleteMapping(basePath)+"/{id}", handler.RemoveImage)
}

func getMapping(path string) string {
	return fmt.Sprintf("GET %s", path)
}

func postMapping(path string) string {
	return fmt.Sprintf("POST %s", path)
}

func putMapping(path string) string {
	return fmt.Sprintf("PUT %s", path)
}

func patchMapping(path string) string {
	return fmt.Sprintf("PUT %s", path)
}

func deleteMapping(path string) string {
	return fmt.Sprintf("DELETE %s", path)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
