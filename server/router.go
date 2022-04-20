package server

import (
	"net/http"
	"os"

	"go_example_server/api"
	"go_example_server/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(logDirectory *string) *echo.Echo {
	log.AppLog.Info("server started")
	e := echo.New()

	_, err := os.Stat(*logDirectory)
	if os.IsNotExist(err) {
		err = os.Mkdir(*logDirectory, 0666)
		if err != nil {
			panic(err)
		}
	}
	fp, err := os.OpenFile(*logDirectory+"/debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	e.Debug = true
	defaultAllowHeaders := []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccept, echo.HeaderXRequestedWith, echo.HeaderAccessControlRequestMethod, echo.HeaderAccessControlRequestHeaders}
	// Set Bundle MiddleWare
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: fp,
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     defaultAllowHeaders,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}))

	// Routes
	route := e.Group("")
	{
		route.GET("/", func(c echo.Context) error {
			return c.File("index.html")
		})
		route.GET("/healthcheck", func(c echo.Context) error {
			return c.String(http.StatusOK, "ok")
		})
	}

	v1 := e.Group("/api")
	{
		v1.POST("/user/upload", api.UploadCSV)
		v1.GET("/user/download", api.DownloadCSV)
	}

	return e
}
