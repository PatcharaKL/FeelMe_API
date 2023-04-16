package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	models "github.com/PatcharaKL/FeelMe_API/rest/Models"
	"github.com/PatcharaKL/FeelMe_API/rest/tokens"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

}
func endpointTokenHandler(e *echo.Echo, h *tokens.Handler) {
	e.POST("/newtoken", h.NewTokenHandler)
}

func endpointUserHandler(e *echo.Echo, h *users.Handler) {
	r := e.Group("/users")
	e.POST("/login", h.UserLoginHandler)
	e.POST("/logout", h.UserLogOutHandler)
	e.GET("/health-check", h.HealthCheckHandler)
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(tokens.JwtCustomClaims)
		},
		SigningKey: []byte(tokens.Signingkey),
	}
	r.Use(echojwt.WithConfig(config))

	r.GET("/employees/", h.GetAllUserHandler)
	r.POST("/employees/:id/happiness-points", h.HappinesspointHandler)
	r.GET("/employees/:id/:period/happiness-points", h.GetHappinessByUserId)
	r.GET("/profile", h.GetUserDetailByUserId)
}
func main() {
	db := models.InitDB()
	defer db.Close()

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	middlewareHandler(e)

	endpointUserHandler(e, users.NewApplication(db))
	endpointTokenHandler(e, tokens.NewApplication(db))

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
