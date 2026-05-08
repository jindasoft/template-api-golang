package handlers

import (
	_ "template-api-golang/internal/api/xxxxxs/models"

	"github.com/jindasoft/jinda-platform/xres"
	"github.com/labstack/echo/v5"
)

// GetXxxxxByID godoc
// @Summary Get Xxxxx by ID
// @Description Get a specific xxxxx record by ID
// @Tags xxxxxs
// @Accept json
// @Produce json
// @Param id path string true "Xxxxx ID"
// @Success 200 {object} xres.SuccessResponse[models.GetXxxxxByIDResponse]
// @Failure 422 {object} xres.UnprocessableEntityResponse "type: operation_failed"
// @Router /v1/xxxxxs/{id} [get]
func (h *handler) GetXxxxxByID(c *echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()
	res, err := h.service.ViewXxxxxByID(ctx, id)
	if err != nil {
		return xres.UnprocessableEntity(c, err.Error())
	}

	return xres.Success(c, res)
}
