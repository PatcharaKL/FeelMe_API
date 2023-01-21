package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PatcharaKL/FeelMe_API/rest/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func middlewareHandler(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
func endpointHandler(e *echo.Echo, h *users.Handler) {
	e.POST("/userlogin", h.UserLoginHandler)

}
func main() {
	db := users.InitDB()
	defer db.Close()

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	middlewareHandler(e)

	endpointHandler(e, users.NewApplication(db))

	go func() {
		if err := e.Start(":5000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down server")
		}
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	e.Logger.Print("Server shuted down")
}
