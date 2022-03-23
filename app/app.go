package app

import (
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/leoncolt/conv/common"
	"github.com/leoncolt/conv/env"
	"github.com/leoncolt/conv/models"
	"github.com/leoncolt/conv/router"
)

type App struct {
	Engine *echo.Echo
	Router *router.Router
}

func New() *App {
	app := new(App)
	app.Engine = bootstrap()
	return app
}

func (app *App) Register() {
	app.Router = new(router.Router)
	app.Router.Register(app.Engine)

	prometheusInstance := prometheus.NewPrometheus("echo", nil)
	prometheusInstance.Use(app.Engine)
}

func (app *App) Start() {
	app.Engine.Logger.Fatal(app.Engine.Start(env.GetAddress()))
}

func bootstrap() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = common.NewValidator()

	rateConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 100, Burst: 50, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusInternalServerError, models.HTTPError{
				Code:    500,
				Message: "Internal Server Error",
			})
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, models.HTTPError{
				Code:    429,
				Message: "too many request.",
			})
		},
	}

	e.Use(middleware.RateLimiterWithConfig(rateConfig))

	return e
}
