package handlers

import (
	"template-api-golang/internal/api/xxxxxs/models"

	"github.com/jindasoft/jinda-platform/xres"
	"github.com/labstack/echo/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PutXxxxx godoc
// @Summary Update a xxxxx
// @Description Update a xxxxx's details by its ID
// @Tags xxxxxs
// @Accept json
// @Produce json
// @Param id path string true "Xxxxx ID"
// @Param request body models.PutXxxxxStatusRequest true "Update Xxxxx Status Request"
// @Success 200 {object} models.PutXxxxxResponse
// @Failure 400 {object} xres.BadRequestResponse "type: bad_request"
// @Failure 422 {object} xres.UnprocessableEntityResponse "type: operation_failed"
// @Router /xxxxxs/{id}/status [put]
func (h *handler) PutSetActiveXxxxx(c *echo.Context) error {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return xres.BadRequestBindData(c, err.Error())
	}

	var req models.PutXxxxxStatusRequest
	if err := c.Bind(&req); err != nil {
		return xres.BadRequestBindData(c, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return xres.BadRequestValidation(c, err.Error())
	}

	ctx := c.Request().Context()
	res, err := h.service.EditXxxxxStatus(ctx, oid, &req)
	if err != nil {
		return xres.UnprocessableEntity(c, err.Error())
	}

	return xres.Updated(c, res)
}
