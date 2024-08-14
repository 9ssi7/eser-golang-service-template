package app

import (
	"net/http"

	"github.com/eser/go-service/pkg/bliss"
	"github.com/eser/go-service/pkg/bliss/httpfx"
	"github.com/eser/go-service/pkg/bliss/httpfx/middlewares"
	"github.com/eser/go-service/pkg/bliss/httpfx/modules/healthcheck"
	"github.com/eser/go-service/pkg/bliss/httpfx/modules/openapi"
	"go.uber.org/fx"
)

var appConfig = AppConfig{}

// configfx.Load(&appConfig)

var Module = fx.Module( //nolint:gochecknoglobals
	"app",
	fx.Invoke(
		RegisterRoutes,
	),
	healthcheck.Module,
	openapi.Module,
)

func RegisterRoutes(routes *httpfx.Router) {
	routes.Use(middlewares.ErrorHandlerMiddleware())
	routes.Use(middlewares.ResponseTimeMiddleware())
	routes.Use(middlewares.CorrelationIdMiddleware())
	routes.Use(middlewares.CorsMiddleware())

	routes.
		Route("GET /", func(ctx *httpfx.Context) httpfx.Response {
			return ctx.Results.PlainText("Hello, World!")
		}).
		HasSummary("Homepage").
		HasDescription("This is the homepage of the service.").
		HasResponse(http.StatusOK)
}

func New() *fx.App {
	return fx.New(
		// fx.WithLogger(bliss.GetFxLogger),
		bliss.Module,
		Module,
	)
}
