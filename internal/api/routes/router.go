package routes

import (
	"net/http"
	"time"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	authHandler "github.com/emoss08/trenova/internal/api/handlers/auth"
	"github.com/emoss08/trenova/internal/api/handlers/commodity"
	"github.com/emoss08/trenova/internal/api/handlers/documentqualityconfig"
	"github.com/emoss08/trenova/internal/api/handlers/equipmentmanufacturer"
	"github.com/emoss08/trenova/internal/api/handlers/equipmenttype"
	"github.com/emoss08/trenova/internal/api/handlers/fleetcode"
	"github.com/emoss08/trenova/internal/api/handlers/hazardousmaterial"
	"github.com/emoss08/trenova/internal/api/handlers/location"
	"github.com/emoss08/trenova/internal/api/handlers/locationcategory"
	organizationHandler "github.com/emoss08/trenova/internal/api/handlers/organization"
	"github.com/emoss08/trenova/internal/api/handlers/reporting"
	"github.com/emoss08/trenova/internal/api/handlers/search"
	"github.com/emoss08/trenova/internal/api/handlers/servicetype"
	"github.com/emoss08/trenova/internal/api/handlers/session"
	"github.com/emoss08/trenova/internal/api/handlers/shipmenttype"
	"github.com/emoss08/trenova/internal/api/handlers/tableconfiguration"
	"github.com/emoss08/trenova/internal/api/handlers/user"
	"github.com/emoss08/trenova/internal/api/handlers/usstate"
	"github.com/emoss08/trenova/internal/api/handlers/worker"
	"github.com/emoss08/trenova/internal/api/middleware"
	"github.com/emoss08/trenova/internal/api/server"
	"github.com/emoss08/trenova/internal/core/services/auth"
	"github.com/emoss08/trenova/internal/infrastructure/cache/redis"
	"github.com/emoss08/trenova/internal/pkg/config"
	"github.com/emoss08/trenova/internal/pkg/logger"
	"github.com/emoss08/trenova/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"go.uber.org/fx"
)

type RouterParams struct {
	fx.In

	// Config
	Config *config.Manager

	// Logger
	Logger *logger.Logger

	// Server
	Server *server.Server

	// Redis
	Redis *redis.Client

	// Services
	AuthService *auth.Service

	// Handlers
	OrganizationHandler          *organizationHandler.Handler
	StateHandler                 *usstate.Handler
	ErrorHandler                 *validator.ErrorHandler
	AuthHandler                  *authHandler.Handler
	UserHandler                  *user.Handler
	SessionHandler               *session.Handler
	SearchHandler                *search.Handler
	WorkerHandler                *worker.Handler
	TableConfigurationHandler    *tableconfiguration.Handler
	FleetCodeHandler             *fleetcode.Handler
	DocumentQualityConfigHandler *documentqualityconfig.Handler
	EquipmentTypeHandler         *equipmenttype.Handler
	EquipmentManufacturerHandler *equipmentmanufacturer.Handler
	ShipmentTypeHandler          *shipmenttype.Handler
	ServiceTypeHandler           *servicetype.Handler
	HazardousMaterialHandler     *hazardousmaterial.Handler
	CommodityHandler             *commodity.Handler
	LocationCategoryHandler      *locationcategory.Handler
	ReportingHandler             *reporting.Handler
	LocationHandler              *location.Handler
}

type Router struct {
	p       RouterParams
	app     fiber.Router
	cfg     *config.Manager
	graphQL *gqlHandler.Server
}

func NewRouter(p RouterParams) *Router {
	return &Router{
		p:       p,
		app:     p.Server.Router(),
		cfg:     p.Config,
		graphQL: p.Server.GraphQL(),
	}
}

func (r *Router) Setup() {
	// API Versioning
	v1 := r.app.Group("api/v1")

	// define the rate limit middleware
	rl := middleware.NewRateLimit(middleware.RateLimitParams{
		Logger: r.p.Logger,
		Redis:  r.p.Redis,
	})

	// setup the global middlewares
	r.setupMiddleware()

	// TODO(Wolfred) Register check and metrics endpoints here

	r.p.AuthHandler.RegisterRoutes(v1)
	r.setupProtectedRoutes(v1, rl)

	// GraphQL
	r.setupGraphQLRoutes()
}

// setupMiddleware configures the global middleware stack
func (r *Router) setupMiddleware() {
	logConfig := middleware.LogConfig{
		CustomTags: map[string]string{
			"env": r.cfg.App().Environment,
			"app": r.cfg.App().Name,
		},
		SlowRequestThreshold: 200 * time.Millisecond,
		LogHeaders:           []string{"X-Request-ID", "Content-Type", "Authorization"},
		ExcludePaths:         []string{"/health", "/metrics"},
		Skip: func(c *fiber.Ctx) bool {
			return c.Path() == "/api/v1/auth/login"
		},
		LogRequestBody:  true,
		LogResponseBody: true,
	}

	r.app.Use(
		favicon.New(),
		compress.New(),
		helmet.New(),

		middleware.NewLogger(r.p.Logger, logConfig),
		encryptcookie.New(encryptcookie.Config{
			Key: r.cfg.Server().SecretKey,
		}),
		cors.New(cors.Config{
			AllowOrigins:     r.cfg.Cors().AllowedOrigins,
			AllowCredentials: r.cfg.Cors().AllowCredentials,
			AllowHeaders:     r.cfg.Cors().AllowedHeaders,
			AllowMethods:     r.cfg.Cors().AllowedMethods,
		}),
		pprof.New(),
		requestid.New(),
		idempotency.New(),
	)
}

func (r *Router) graphQLHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.graphQL.ServeHTTP(w, req)
	})
}

func (r *Router) graphQLPlaygroundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		playground.Handler("GraphQL playground", "/graphql").ServeHTTP(w, req)
	})
}

func (r *Router) setupGraphQLRoutes() {
	r.app.All("/graphql", adaptor.HTTPHandler(r.graphQLHandler()))

	r.app.Get("/graphql/playground", adaptor.HTTPHandler(r.graphQLPlaygroundHandler()))
}

// setupProtectedRoutes configures the protected routes
func (r *Router) setupProtectedRoutes(router fiber.Router, rl *middleware.RateLimiter) {
	router.Use(middleware.NewAuthMiddleware(middleware.AuthMiddlewareParams{
		Logger: r.p.Logger,
		Config: r.cfg,
		Auth:   r.p.AuthService,
	}).Authenticate())

	// Organization
	r.p.OrganizationHandler.RegisterRoutes(router, rl)

	// US States
	r.p.StateHandler.RegisterRoutes(router, rl)

	// Users
	r.p.UserHandler.RegisterRoutes(router, rl)

	// Sessions
	r.p.SessionHandler.RegisterRoutes(router)

	// Search
	r.p.SearchHandler.RegisterRoutes(router)

	// Workers
	r.p.WorkerHandler.RegisterRoutes(router, rl)

	// Table Configurations
	r.p.TableConfigurationHandler.RegisterRoutes(router, rl)

	// Fleet Codes
	r.p.FleetCodeHandler.RegisterRoutes(router, rl)

	// Document Quality Configs
	r.p.DocumentQualityConfigHandler.RegisterRoutes(router, rl)

	// Equipment Types
	r.p.EquipmentTypeHandler.RegisterRoutes(router, rl)

	// Equipment Manufacturers
	r.p.EquipmentManufacturerHandler.RegisterRoutes(router, rl)

	// Shipment Types
	r.p.ShipmentTypeHandler.RegisterRoutes(router, rl)

	// Service Types
	r.p.ServiceTypeHandler.RegisterRoutes(router, rl)

	// Hazardous Materials
	r.p.HazardousMaterialHandler.RegisterRoutes(router, rl)

	// Commodities
	r.p.CommodityHandler.RegisterRoutes(router, rl)

	// Location Categories
	r.p.LocationCategoryHandler.RegisterRoutes(router, rl)

	// Reporting
	r.p.ReportingHandler.RegisterRoutes(router, rl)

	// Locations
	r.p.LocationHandler.RegisterRoutes(router, rl)
}
