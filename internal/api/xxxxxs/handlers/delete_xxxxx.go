package handlers

import (
	"github.com/jindasoft/jinda-platform/xres"
	"github.com/labstack/echo/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteXxxxx godoc
// @Summary Delete a xxxxx
// @Description Soft delete a xxxxx by ID
// @Tags xxxxxs
// @Accept json
// @Produce json
// @Param id path string true "Xxxxx ID"
// @Success 204 {object} nil
// @Failure 400 {object} xres.BadRequestResponse "type: bad_request"
// @Failure 422 {object} xres.UnprocessableEntityResponse "type: operation_failed"
// @Router /xxxxxs/{id} [delete]
func (h *handler) DeleteXxxxx(c *echo.Context) error {
	id := c.Param("id")
	xxxxxID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return xres.BadRequestBindData(c, "invalid xxxxx ID")
	}

	ctx := c.Request().Context()
	err = h.service.DeleteXxxxx(ctx, xxxxxID)
	if err != nil {
		return xres.UnprocessableEntity(c, err.Error())
	}

	return xres.Deleted(c)
}
