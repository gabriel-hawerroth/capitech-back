package web

import (
	"database/sql"
	"fmt"

	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/web/handlers"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/web/webserver"
	"github.com/gabriel-hawerroth/capitech-back/internal/services"
	awsclients "github.com/gabriel-hawerroth/capitech-back/third_party/aws"
)

var server *webserver.WebServer
var db *sql.DB
var s3Client *awsclients.S3Client

func LoadHandlers(webServer *webserver.WebServer, dbConn *sql.DB, s3 *awsclients.S3Client) {
	server = webServer
	db = dbConn
	s3Client = s3

	loadCategoryHandlers()
	loadShoppingCartHandlers()
	loadProductHandlers()
}

func loadCategoryHandlers() {
	const basePath = "/category"

	repository := repositories.NewCategoryRepository(db)
	service := services.NewCategoryService(*repository)
	handler := handlers.NewCategoryHandler(*service)

	server.AddHandler(getMapping(basePath), handler.GetCategoriesList)
}

func loadShoppingCartHandlers() {
	const basePath = "/shopping-cart"

	repository := repositories.NewShoppingCartRepository(db)
	service := services.NewShoppingCartService(*repository)
	handler := handlers.NewShoppingCartHandler(*service)

	server.AddHandler(postMapping(basePath), handler.AddProduct)
	server.AddHandler(getMapping(basePath)+"/getUserShoppingCart", handler.GetUserShoppingCart)
}

func loadProductHandlers() {
	const basePath = "/product"

	repository := repositories.NewProductRepository(db)
	service := services.NewProductService(*repository, *s3Client)
	handler := handlers.NewProductHandler(*service)

	server.AddHandler(getMapping(basePath)+"/{id}", handler.GetById)
	server.AddHandler(getMapping(basePath), handler.GetProductsList)
	server.AddHandler(getMapping(basePath)+"/getTrendingProductsList", handler.GetTrendingProducts)
	server.AddHandler(getMapping(basePath)+"/getBestSellingProductsList", handler.GetBestSellingProducts)
	server.AddHandler(getMapping(basePath)+"/getUserSearchHistory", handler.GetUserSearchHistory)
	server.AddHandler(postMapping(basePath), handler.Create)
	server.AddHandler(putMapping(basePath)+"/{id}", handler.Update)
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
	return fmt.Sprintf("PATCH %s", path)
}

func deleteMapping(path string) string {
	return fmt.Sprintf("DELETE %s", path)
}
