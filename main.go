package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/eggysetiawan/fiber-starterkit/config"
	"github.com/eggysetiawan/fiber-starterkit/internal/domain"
	"github.com/eggysetiawan/fiber-starterkit/internal/handlers"
	"github.com/eggysetiawan/fiber-starterkit/internal/repository"
	"github.com/eggysetiawan/fiber-starterkit/internal/usecases"
	"github.com/eggysetiawan/fiber-starterkit/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	*fiber.App
}

func main() {
	utils.InitLoadEnv()
	config := config.New()

	app := App{
		App: fiber.New(*config.GetFiberConfig()),
	}

	// logger
	file, logErr := utils.NewLogFile()

	if logErr != nil {
		fmt.Println("failed to init logger", logErr.Message)
		return
	}

	defer file.Close()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			utils.User: utils.Pwd,
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return domain.NewUnauthorizedResponse(c)
		},
	}),
		logger.New(logger.Config{
			Format:       "${pid} ${time} ${method} ${path} ${body} ${status} ${latency}\n",
			TimeFormat:   time.DateTime,
			TimeZone:     "Asia/Jakarta",
			TimeInterval: 500 * time.Millisecond,
			Output:       file,
		}),
	)

	// Initialize Database
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("failed to init connection", err)
		return
	}
	defer db.Close()

	api := app.Group("/api")

	// machines
	machines := api.Group("/machines")
	mh := handlers.NewMachineHandler(usecases.NewMachineUseCase(repository.NewMachineRepositoryDb(db)))
	machines.Post("/findBy", mh.ShowMachine)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		app.exit()
	}()

	// Start the server
	err = app.Listen(config.GetString("SERVER_PORT"))
	if err != nil {
		app.exit()
	}
}

func (app *App) exit() {
	_ = app.Shutdown()
}
