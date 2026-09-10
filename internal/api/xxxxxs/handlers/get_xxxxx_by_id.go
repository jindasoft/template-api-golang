package handlers

import (
	_ "template-api-golang/internal/api/xxxxxs/models"

	"github.com/jindasoft/template-platform-go/xres"
	"github.com/labstack/echo/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetXxxxxByID godoc
// @Summary Get a xxxxx by ID
// @Description Retrieve a xxxxx's details by its ID
// @Tags xxxxxs
// @Accept json
// @Produce json
// @Param id path string true "Xxxxx ID"
// @Success 200 {object} models.GetXxxxxByIDResponse
// @Failure 400 {object} xres.BadRequestResponse "type: bad_request"
// @Failure 422 {object} xres.UnprocessableEntityResponse "type: operation_failed"
// @Router /xxxxxs/{id} [get]
func (h *handler) GetXxxxxByID(c *echo.Context) error {
	id := c.Param("id")
	xxxxxID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return xres.BadRequestBindData(c, "invalid xxxxx ID")
	}

	ctx := c.Request().Context()
	res, err := h.service.ViewXxxxxByID(ctx, xxxxxID)
	if err != nil {
		return xres.UnprocessableEntity(c, err.Error())
	}

	return xres.Success(c, res)
}
