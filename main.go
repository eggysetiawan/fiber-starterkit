package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/eggysetiawan/fiber-starterkit/config"
	"github.com/eggysetiawan/fiber-starterkit/internal/domain"
	"github.com/eggysetiawan/fiber-starterkit/internal/handlers"
	"github.com/eggysetiawan/fiber-starterkit/internal/repository"
	"github.com/eggysetiawan/fiber-starterkit/internal/usecases"
	"github.com/eggysetiawan/fiber-starterkit/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/jmoiron/sqlx"
)

type App struct {
	*fiber.App
	config *config.Config
	args   []string
	db     *sqlx.DB
}

func main() {
	config := config.New()
	logger.NewLogger()

	app := App{
		App:    fiber.New(*config.GetFiberConfig()),
		config: config,
		args:   os.Args,
	}

	// Initialize Database
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println("failed to init connection", err)
		return
	}
	defer db.Close()

	app.db = db

	if len(app.args) > 0 {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		defer fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		app.console()
		return
	}

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			config.GetString("BASIC_AUTH_USER"): config.GetString("BASIC_AUTH_PASS"),
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return domain.NewUnauthorizedResponse(c)
		},
	}))

	api := app.Group("/api")

	// example machines
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

func (app *App) console() {
	fmt.Println(app.args)
	switch app.args[1] {
	case "cmd":
		logger.Info("after fix 2")

	default:
		s := strings.Split(config.AppConfig.GetString("ELASTIC_HOST"), ",")
		fmt.Println(s)
	}

}
