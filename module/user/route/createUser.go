package route

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ferza17/kafka-basic/consumer/exception"
	"github.com/ferza17/kafka-basic/consumer/module/user/dto"
	"github.com/ferza17/kafka-basic/consumer/pkg/logger"
	"github.com/ferza17/kafka-basic/consumer/util"
)

// CreateUser
// @Title       Create User
// @Summary     Create User
// @Description This endpoint allows you to create a new user by providing their details in the request body.
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       Payload       body   dto.CreateUser true "Request Payload"
// @Router      /v1/users [POST]
func (h *Handler) CreateUser(c echo.Context) error {
	var (
		requestID = c.Request().Header.Get("X-Request-ID")
		ctx       = c.Request().Context()
		reqData   = new(dto.CreateUser)
	)
	reqData.RequestID = requestID
	if err := c.Bind(reqData); err != nil {
		logger.WithField("RequestID", reqData.RequestID).Infof(ctx, "error: %v", err)
		return util.Response(c, util.ResponseRequest{Exception: exception.GetException(exception.GENERAL_BAD_REQUEST)})
	}

	if expt := h.UserLogic.CreateUser(ctx, reqData); expt != nil {
		logger.WithField("RequestID", reqData.RequestID).Infof(ctx, "error: %v", expt.Message)
		return util.Response(c, util.ResponseRequest{Exception: expt})
	}

	return util.Response(c, util.ResponseRequest{
		Data:       nil,
		StatusCode: http.StatusCreated,
	})
}
