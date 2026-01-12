package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"seaotterms-db/database"
	blogrouter "seaotterms-db/routes/blog"
)

var (
	// management database connect
	dbs = make(map[string]*gorm.DB)
)

func init() {
	// init logrus settings
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	// init env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf(".env file load error: %v", err)
	}
}

func main() {
	// init migration
	for i := range 2 {
		dbName, db := database.InitDsn(i)
		dbs[dbName] = db
		database.Migration(dbName, dbs[dbName])
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		// AllowOrigins:     os.Getenv("CORS_URL"),
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin,Content-Type,Accept",
		// AllowCredentials: true,
	}))
	// api route group
	apiGroup := app.Group("/api") // main api route group

	// database route group
	blogrouter.BlogRouter(apiGroup, dbs)

	logrus.Fatal(app.Listen(fmt.Sprintf("127.0.0.1:%s", os.Getenv("PORT"))))
}
