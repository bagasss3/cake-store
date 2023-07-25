package console

import (
	"cake-store/src/config"
	"cake-store/src/controller"
	"cake-store/src/database"
	"cake-store/src/repository"
	"cake-store/src/router"
	"cake-store/src/service"

	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  "Start running the server",
	Run:   server,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func server(cmd *cobra.Command, args []string) {
	// Initiate DB
	db := database.NewDB()
	defer db.Close()

	// Create Echo instance
	httpServer := echo.New()
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(middleware.CORS())

	// Depedency Injection
	cakeRepository := repository.NewCakeRepository(db)
	cakeService := service.NewCakeService(cakeRepository)
	cakeController := controller.NewCakeController(cakeService)

	router.RouteService(httpServer.Group("/api"), cakeController)

	// Graceful Shutdown
	// Catch Signal
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigChan)
		defer close(sigChan)

		<-sigChan
		log.Info("Received termination signal, initiating graceful shutdown...")
		cancel()
	}()

	// Start http server
	go func() {
		log.Info("Starting server...")
		if err := httpServer.Start(":" + config.Port()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting the server: %v", err)
		}
	}()

	// Shutting down any connection and server
	<-ctx.Done()
	log.Info("Shutting down server...")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Info("Server gracefully shut down")
}
