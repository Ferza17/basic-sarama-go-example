package restapi

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/ferza17/kafka-basic/consumer/config"
	"github.com/ferza17/kafka-basic/consumer/enum"
	"github.com/ferza17/kafka-basic/consumer/exception"
	"github.com/ferza17/kafka-basic/consumer/pkg/logger"
	"github.com/ferza17/kafka-basic/consumer/util"
)

func setMiddleware(e *echo.Echo) {

	e.Use(middleware.BodyLimit("20M"))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"Authorization",
			"X-Request-RequestID",
			"X-API-KEY",
			"User-Timezone",
			"User-CurrentTime",
		},
	}))

	e.Pre(setEndpointRequestToCtx)
	e.Pre(setXRequestIDIfNotPresent)
	e.Pre(setXRequestIDToContext)
	e.Pre(middleware.BodyDump(middlewareDumpBodyRequest))
	e.Use(middleware.RateLimiterWithConfig(rateLimitConfig()))
}

// setEndpointRequestToCtx ...
func setEndpointRequestToCtx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ctx = context.WithValue(ctx, "request_endpoint", c.Request().RequestURI)
		httpReq := c.Request().WithContext(ctx)
		c.SetRequest(httpReq)
		return next(c)
	}
}

func setXRequestIDIfNotPresent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("X-Request-ID") == "" {
			c.Request().Header.Set("X-Request-ID", util.GetRandomString())
		}
		return next(c)
	}
}

func setXRequestIDToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx   = c.Request().Context()
			key   = "X-Request-RequestID"
			value = c.Request().Header.Get("X-Request-ID")
		)

		ctx = context.WithValue(ctx, key, value)
		httpReq := c.Request().WithContext(ctx)
		c.SetRequest(httpReq)
		return next(c)
	}
}

// middlewareDumpBodyRequest ...
func middlewareDumpBodyRequest(c echo.Context, reqBody, resBody []byte) {
	if (config.Get().Env == enum.CONFIG_ENV_PROD || config.Get().Env == enum.CONFIG_ENV_STAGING) && config.Get().LogRequestBody {
		requestID := c.Request().Header.Get("X-Request-ID")
		clientVersion := c.Request().Header.Get("Client-Version-Code")
		clientPlatformCode := c.Request().Header.Get("Client-Platform-Code")
		logger.WithFields(map[string]interface{}{
			"requestIP":          c.Request().Header.Get("X-Real-IP"),
			"requestID":          requestID,
			"method":             c.Request().Method,
			"requestURL":         c.Request().RequestURI,
			"requestBody":        string(reqBody),
			"responseBody":       string(resBody),
			"clientVersion":      clientVersion,
			"clientPlatformCode": clientPlatformCode,
		}).Debugf(context.TODO(), "middlewareDumpBodyRequest")
	}
}

// rateLimitConfig
func rateLimitConfig() middleware.RateLimiterConfig {
	return middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 5, Burst: 15, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return util.Response(context, util.ResponseRequest{
				Exception: exception.GetException(exception.GENERAL_TOO_MANY_REQUEST),
			})
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return util.Response(context, util.ResponseRequest{
				Exception: exception.GetException(exception.GENERAL_TOO_MANY_REQUEST),
			})
		},
	}
}
