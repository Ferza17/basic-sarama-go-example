package restapi

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/ferza17/kafka-basic/consumer/enum"

	_ "github.com/ferza17/kafka-basic/consumer/docs"

	"github.com/ferza17/kafka-basic/consumer/config"
)

type Router struct {
	*echo.Echo
}

func NewRouter() *Router {
	e := echo.New()

	e.IPExtractor = echo.ExtractIPDirect()

	// service health check
	e.GET("/info", func(c echo.Context) error {

		type Info struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
				Build   string `json:"build"`
				Commit  string `json:"commit"`
				UpSince string `json:"upsince"`
				UpTime  string `json:"uptime"`
				Version string `json:"version"`
			} `json:"data"`
			App string `json:"app"`
		}

		appInfo := Info{
			Code:    200,
			Message: "",
			App:     config.AppName,
			Data: struct {
				Build   string "json:\"build\""
				Commit  string "json:\"commit\""
				UpSince string "json:\"upsince\""
				UpTime  string "json:\"uptime\""
				Version string "json:\"version\""
			}{
				Build:   config.Build,
				Commit:  config.Commit,
				UpSince: config.Now.Format(time.RFC3339),
				UpTime:  time.Since(config.Now).String(),
				Version: config.Version,
			},
		}

		return c.JSON(http.StatusOK, appInfo)
	})

	// swagger route docs api
	if config.Get().Env == enum.CONFIG_ENV_STAGING ||
		config.Get().Env == enum.CONFIG_ENV_LOCAL {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	// handdle middleware
	setMiddleware(e)

	return &Router{e}
}
