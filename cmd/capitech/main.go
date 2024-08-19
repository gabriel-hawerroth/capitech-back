package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gabriel-hawerroth/capitech-back/configs"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/web"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/web/webserver"
	awsclients "github.com/gabriel-hawerroth/capitech-back/third_party/aws"
	_ "github.com/lib/pq"
)

var confs *configs.Conf

func main() {
	confs = loadConfigs()

	db := openDatabaseConnection()
	defer db.Close()

	s3Client := createS3Client()

	startWebServer(db, s3Client)
}

func loadConfigs() *configs.Conf {
	confs, err := configs.LoadConfig(".")
	checkError(err)
	return confs
}

func openDatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", confs.DBUser, confs.DBPassword, confs.DBName, confs.DBHost, confs.DBPort))
	checkError(err)

	err = db.Ping()
	checkError(err)

	database.MigrateSchema(db)

	return db
}

func createS3Client() *awsclients.S3Client {
	awsS3Client, err := awsclients.NewS3Client(confs.AwsIamAccessKey, confs.AwsIamSecretKey)
	checkError(err)

	log.Println("Successfully connected to AWS S3")
	return awsS3Client
}

func startWebServer(db *sql.DB, s3Client *awsclients.S3Client) {
	webServer := webserver.NewWebServer(confs.WebServerPort)
	web.LoadHandlers(webServer, db, s3Client)
	webServer.Start()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
