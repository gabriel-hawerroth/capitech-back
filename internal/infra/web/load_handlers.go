package web

import (
	"database/sql"
	"fmt"
	"net/http"

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

	loadAuthHandlers()
	loadCategoryHandlers()
	loadShoppingCartHandlers()
	loadProductHandlers()
}

func loadAuthHandlers() {
	const basePath = "/auth"

	repository := repositories.NewUserRepository(db)
	service := services.NewAuthService(*repository)
	handler := handlers.NewAuthHandler(*service)

	postMapping(basePath+"/login", handler.DoLogin)
	postMapping(basePath+"/createUser", handler.CreateNewUser)
}

func loadCategoryHandlers() {
	const basePath = "/category"

	repository := repositories.NewCategoryRepository(db)
	service := services.NewCategoryService(*repository)
	handler := handlers.NewCategoryHandler(*service)

	getMapping(basePath, handler.GetCategoriesList)
}

func loadShoppingCartHandlers() {
	const basePath = "/shopping-cart"

	repository := repositories.NewShoppingCartRepository(db)
	service := services.NewShoppingCartService(*repository)
	handler := handlers.NewShoppingCartHandler(*service)

	postMapping(basePath, handler.AddProduct)
	getMapping(basePath+"/getUserShoppingCart", handler.GetUserShoppingCart)
}

func loadProductHandlers() {
	const basePath = "/product"

	repository := repositories.NewProductRepository(db)
	searchLogRepository := repositories.NewSearchLogRepository(db)

	searchLogService := services.NewSearchLogService(*searchLogRepository)
	service := services.NewProductService(*repository, *s3Client, *searchLogService)

	handler := handlers.NewProductHandler(*service)

	getMapping(basePath, handler.GetProductsList)
	getMapping(basePath+"/{id}", handler.GetById)
	getMapping(basePath+"/getTrendingProductsList", handler.GetTrendingProducts)
	getMapping(basePath+"/getBestSellingProductsList", handler.GetBestSellingProducts)
	getMapping(basePath+"/getUserSearchHistory", handler.GetUserSearchHistory)

	postMapping(basePath, handler.Create)

	putMapping(basePath+"/{id}", handler.Update)
	patchMapping(basePath+"/editProductPrice/{id}", handler.ChangePrice)
	patchMapping(basePath+"/editProductStockQuantity/{id}", handler.ChangeStockQuantity)
	patchMapping(basePath+"/changeProductImage/{id}", handler.ChangeImage)
	patchMapping(basePath+"/removeProductImage/{id}", handler.RemoveImage)

	deleteMapping(basePath+"/{id}", handler.RemoveImage)
}

func getMapping(path string, handler http.HandlerFunc) {
	server.AddHandler(fmt.Sprintf("GET %s", path), handler)
}

func postMapping(path string, handler http.HandlerFunc) {
	server.AddHandler(fmt.Sprintf("POST %s", path), handler)
}

func putMapping(path string, handler http.HandlerFunc) {
	server.AddHandler(fmt.Sprintf("PUT %s", path), handler)
}

func patchMapping(path string, handler http.HandlerFunc) {
	server.AddHandler(fmt.Sprintf("PATCH %s", path), handler)
}

func deleteMapping(path string, handler http.HandlerFunc) {
	server.AddHandler(fmt.Sprintf("DELETE %s", path), handler)
}
