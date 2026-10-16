package handlers

import (
	"template-api-golang/internal/api/xxxxxs/models"

	"github.com/jindasoft/template-platform-go/xres"
	"github.com/labstack/echo/v5"
)

// PostXxxxx godoc
// @Summary Create a new xxxxx
// @Description Create a new xxxxx
// @Tags xxxxxs
// @Accept json
// @Produce json
// @Param request body models.PostXxxxxRequest true "request body"
// @Success 201 {object} xres.SuccessResponse[models.PostXxxxxResponse]
// @Failure 400 {object} xres.BadRequestResponse "type: bad_request"
// @Failure 422 {object} xres.UnprocessableEntityResponse "type: operation_failed"
// @Router /xxxxxs [post]
func (h *handler) PostXxxxx(c *echo.Context) error {
	var req models.PostXxxxxRequest
	if err := c.Bind(&req); err != nil {
		return xres.BadRequestBindData(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return xres.BadRequestValidation(c, err.Error())
	}

	ctx := c.Request().Context()
	res, err := h.service.AddXxxxx(ctx, &req)
	if err != nil {
		return xres.UnprocessableEntity(c, err.Error())
	}

	return xres.Created(c, res)
}
