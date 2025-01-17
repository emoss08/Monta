package documentqualityconfig

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trenova-app/transport/internal/api/middleware"
	"github.com/trenova-app/transport/internal/core/ports/repositories"
	"github.com/trenova-app/transport/internal/core/services/documentqualityconfig"
	"github.com/trenova-app/transport/internal/pkg/ctx"
	"github.com/trenova-app/transport/internal/pkg/validator"
	"go.uber.org/fx"
)

type Handler struct {
	ds *documentqualityconfig.Service
	eh *validator.ErrorHandler
}

type HandlerParams struct {
	fx.In

	DocumentQualityConfigService *documentqualityconfig.Service
	ErrorHandler                 *validator.ErrorHandler
}

func NewHandler(p HandlerParams) *Handler {
	return &Handler{ds: p.DocumentQualityConfigService, eh: p.ErrorHandler}
}

func (h Handler) RegisterRoutes(r fiber.Router, rl *middleware.RateLimiter) {
	api := r.Group("/document-quality-configs")

	api.Get("/", rl.WithRateLimit(
		[]fiber.Handler{h.get},
		middleware.PerMinute(60),
	)...)
}

func (h Handler) get(c *fiber.Ctx) error {
	reqCtx, err := ctx.WithRequestContext(c)
	if err != nil {
		return h.eh.HandleError(c, err)
	}

	dqc, err := h.ds.Get(c.UserContext(), &repositories.GetDocumentQualityConfigOptions{
		OrgID:  reqCtx.OrgID,
		BuID:   reqCtx.BuID,
		UserID: reqCtx.UserID,
	})
	if err != nil {
		return h.eh.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(dqc)
}