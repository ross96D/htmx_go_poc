package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// VaryCache middleware adds a `Vary` header to the response.
func VaryCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderVary, "HX-Request")
		return next(c)
	}
}

func DefaultResponseHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderVary, "HX-Request")
		return next(c)
	}
}

func Logger() echo.MiddlewareFunc {
	loggerConf := middleware.LoggerConfig{
		Format: "${time_rfc3339} ip:${remote_ip} method:${method} uri:${uri} ${protocol} scode:${status} e:${error} ms:${latency} in:${bytes_in} out:${bytes_out}\n",
	}
	return middleware.LoggerWithConfig(loggerConf)
}
