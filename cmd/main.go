package main

import (

	"fmt"
	"log"
	"os"
	"time"
	"github.com/sepehrmohseni/go-web-boilerplate/routes"
	"github.com/sepehrmohseni/go-web-boilerplate/config"
	"github.com/sepehrmohseni/go-web-boilerplate/database"


	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// get config file
	conf, err := config.GetConfig("./config.yml")
	if err != nil {
		panic(err)
	}

	// registering app with custom json encoder
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Prefork:     true,
	})

	// registering middlewares
	app.Use(cors.New())
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   10 * time.Second,
		CacheControl: true,
	}))
	app.Use(etag.New())
	app.Use(compress.New())
	app.Use(pprof.New())
	app.Use(recover.New())

	// logging requests and writing to file
	logName := fmt.Sprintf("./logs/%v.log", time.Now().Format("02-January-2006"))
	logFile, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()
	app.Use(logger.New(logger.Config{
		Output:     logFile,
		TimeFormat: "02-Jan-2006, 15:04:05",
		TimeZone:   "Iran/Tehran",
		Format:     "\nTime: ${time}\nReferer: ${referer}\nUserAgent: ${ua}\nMethod: ${method}\nPath: ${path}\nStatus: ${status}\nLatency: ${latency}\n\n---------------------\n\n",
	}))

	// connect database
	database.ConnectDB()

	// parent route group
	v1 := app.Group(conf.App.BaseURL)

	// metrics route
	metricsRoute := v1.Group("/metrix")
	metricsRoute.Get("/", monitor.New())

	// test routes for app
	testRoute := v1.Group("/tst")
	routes.TestRoutes(testRoute)

	// starting app
	log.Fatal(app.Listen(conf.App.Port))
}
