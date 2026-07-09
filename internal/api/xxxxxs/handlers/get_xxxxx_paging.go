package handlers

import (
	"template-api-golang/internal/api/xxxxxs/models"

	"github.com/jindasoft/jinda-platform/xres"
	"github.com/labstack/echo/v5"
)

// GetXxxxxPaging godoc
// @Summary Get xxxxx paging
// @Description Get xxxxx paging
// @Tags xxxxxs
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} xres.PagingResponse[[]models.GetXxxxxPagingResponse]
// @Failure 400 {object} xres.BadRequestResponse "type: bad_request"
// @Failure 422 {object} xres.UnprocessableEntityResponse "type: operation_failed"
// @Router /xxxxxs [get]
func (h *handler) GetXxxxxPaging(c *echo.Context) error {
	var req models.GetXxxxxPagingRequest
	if err := c.Bind(&req); err != nil {
		return xres.BadRequestBindData(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return xres.BadRequestValidation(c, err.Error())
	}

	ctx := c.Request().Context()
	res, meta, err := h.service.ViewXxxxxPaging(ctx, &req)
	if err != nil {
		return xres.UnprocessableEntity(c, err.Error())
	}

	return xres.Paging(c, res, meta)
}
