package server

import (
	"net/http"

	"htmx/cmd/web"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/js/*", echo.WrapHandler(fileServer))

	// e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	e.GET("/web", echo.WrapHandler(http.HandlerFunc(web.WebHandler)))
	e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))
	e.POST("/increment", echo.WrapHandler(http.HandlerFunc(web.IncrementWebHandler)))

	e.GET("/", s.HelloWorldHandler)

	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
