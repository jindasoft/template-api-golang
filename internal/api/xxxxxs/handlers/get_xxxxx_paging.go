package handlers

import (
	"template-api-examples/internal/api/xxxxxs/models"

	"github.com/jindasoft/jinda-platforms/xres"
	"github.com/labstack/echo/v5"
)

// GetXxxxxPaging godoc
// @Summary Get Xxxxx with pagination
// @Description Get xxxxx records with pagination support
// @Tags xxxxx
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} models.SwaggerGetXxxxxPagingResponse
// @Failure 400 {object} xres.BadRequestResponse "type: bind_data_error, validation_error"
// @Failure 422 {object} xres.UnprocessableEntityResponse "type: operation_failed"
// @Router /v1/xxxxx [get]
func (h *handler) GetXxxxxPaging(c *echo.Context) error {
	var req models.GetXxxxxPagingRequest
	if err := c.Bind(&req); err != nil {
		return xres.BadRequestBindData(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return xres.BadRequestValidation(c, err.Error())
	}

	ctx := c.Request().Context()
	res, err := h.service.ViewXxxxxPaging(ctx, &req)
	if err != nil {
		return xres.UnprocessableEntity(c, err.Error())
	}

	meta := xres.PagingMeta{
		TotalItems: 10,
		Offset:     req.Offset,
		Limit:      req.Limit,
	}

	return xres.Paging(c, res, meta)
}
