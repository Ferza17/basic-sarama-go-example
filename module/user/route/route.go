package route

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	userLogic "github.com/ferza17/kafka-basic/consumer/module/user/logic"
	"github.com/ferza17/kafka-basic/consumer/restapi"
)

type Handler struct {
	fx.In
	UserLogic userLogic.UserLogic
	EchoRoute *restapi.Router
}

func NewRoutes(h Handler, m ...echo.MiddlewareFunc) Handler {
	h.Route(m...)
	return h
}

func (h *Handler) Route(m ...echo.MiddlewareFunc) {
	r := h.EchoRoute.Group("v1/users", m...)
	r.POST("", h.CreateUser)
}
