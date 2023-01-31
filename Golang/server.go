package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PatcharaKL/FeelMe_API/rest/users"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func middlewareHandler(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
func endpointHandler(e *echo.Echo, h *users.Handler) {
	r := e.Group("/users")
	e.POST("/login", h.UserLoginHandler)
	e.POST("/logout", h.UserLogOutHandler)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(users.JwtCustomClaims)
		},
		//! Change setring key to get from env
		SigningKey: []byte("GVebOWpKrqyZ9RwPXzazpNpcmA6njskh"),
	}
	r.Use(echojwt.WithConfig(config))

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
