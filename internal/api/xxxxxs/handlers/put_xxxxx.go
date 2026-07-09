package handlers

import (
	"template-api-golang/internal/api/xxxxxs/models"

	"github.com/jindasoft/jinda-platform/xres"
	"github.com/labstack/echo/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PutXxxxx godoc
// @Summary Update an existing xxxxx
// @Description Update an existing xxxxx
// @Tags xxxxxs
// @Accept json
// @Produce json
// @Param id path string true "Xxxxx ID"
// @Param request body models.PutXxxxxRequest true "request body"
// @Success 200 {object} xres.SuccessResponse[models.PutXxxxxResponse]
// @Failure 400 {object} xres.BadRequestResponse "type: bad_request"
// @Failure 422 {object} xres.UnprocessableEntityResponse "type: operation_failed"
// @Router /xxxxxs/{id} [put]
func (h *handler) PutXxxxx(c *echo.Context) error {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return xres.BadRequestBindData(c, err.Error())
	}

	var req models.PutXxxxxRequest
	if err := c.Bind(&req); err != nil {
		return xres.BadRequestBindData(c, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return xres.BadRequestValidation(c, err.Error())
	}

	ctx := c.Request().Context()
	res, err := h.service.EditXxxxx(ctx, oid, &req)
	if err != nil {
		return xres.UnprocessableEntity(c, err.Error())
	}

	return xres.Updated(c, res)
}
