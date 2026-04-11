package handlers

import (
	"template-api-examples/internal/api/xxxxxs/models"

	"github.com/jindasoft/jinda-platforms/xres"
	"github.com/labstack/echo/v5"
)

// PostXxxxx godoc
// @Summary Create a new Xxxxx
// @Description Create a new xxxxx record
// @Tags xxxxx
// @Accept json
// @Produce json
// @Param request body models.PostXxxxxRequest true "Create xxxxx request"
// @Success 201 {object} models.SwaggerPostXxxxxResponse
// @Failure 400 {object} xres.BadRequestResponse "type: bind_data_error, validation_error"
// @Failure 422 {object} xres.UnprocessableEntityResponse
// @Router /v1/xxxxx [post]
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
