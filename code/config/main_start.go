package config

import (
	"cli/helpers"
	"strings"
)

const mainTest string = `
package main

import (
	"{appName}/middleware"
	"{appName}/router"
	"{appName}/static"
	"{appName}/config"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func initDb() *sql.DB {
	i, err := strconv.Atoi(goDotEnvVariable("DB_PORT"))
	psqlInfo := fmt.Sprintf("password=%s dbname=%s sslmode=disable host=%s port=%d user=%s ",
		goDotEnvVariable("DB_PASS"), goDotEnvVariable("DB_DB"), goDotEnvVariable("DB_HOST"),
		i, goDotEnvVariable("DB_USER"),
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("error", err)
		return nil
	}

	return db
}

func init() {
	config.Dev = goDotEnvVariable("ENVIRONMENT") != "production"
	config.DB = initDb()

	if goDotEnvVariable("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	r := gin.Default()
	middleware.Run(r)
	static.Run(r)
	router.Run(r)
	r.Run()
	fmt.Println("RUN IN -> http://localhost:8080")
}
`

func Main(appName string) string {
	return helpers.Trim(strings.ReplaceAll(mainTest, "{appName}", appName))
}
