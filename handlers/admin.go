package handlers

import (
	"delivery/constants"
	"delivery/entities"
	"delivery/pkg/http"
	"delivery/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateXozmak(c *gin.Context) {
	var body entities.Xozmak
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, constants.BadRequest)
		return
	}
	body.ID = uuid.NewString()
	body.CreatedBy = entities.NullString("ab89ca99-3c18-4751-9c07-51a2ee85751e")

	err = h.adminController.CreateXozmak(c, body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}

	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) GetXozmak(c *gin.Context) {
	data, err := h.adminController.GetXozmak(c)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, data)
}

func (h *Handler) UpdateXozmak(c *gin.Context) {
	var body entities.Xozmak
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, constants.BadRequest)
		return
	}
	body.ID = c.Param("id")

	if !utils.IsValidUUID(body.ID) {
		h.handleResponse(c, http.BadRequest, "Invalid UUID format")
		return
	}
	err = h.adminController.UpdateXozmak(c, body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) DeleteXozmak(c *gin.Context) {
	id := c.Param("id")

	if !utils.IsValidUUID(id) {
		h.handleResponse(c, http.BadRequest, "Invalid UUID format")
		return
	}

	err := h.adminController.DeleteXozmak(c, id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var req entities.Category
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, constants.BadRequest)
		return
	}
	req.ID = uuid.NewString()

	err = h.adminController.CreateCategory(c, req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) GetCategory(c *gin.Context) {
	data, err := h.adminController.GetCategory(c)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, data)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	var req entities.Category
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, constants.BadRequest)
		return
	}
	req.ID = c.Param("id")
	if !utils.IsValidUUID(req.ID) {
		h.handleResponse(c, http.BadRequest, "Invalid UUID format")
		return
	}
	req.UpdatedAt = time.Now()
	err = h.adminController.UpdateCategory(c, req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	category_id := c.Param("id")

	if !utils.IsValidUUID(category_id) {
		h.handleResponse(c, http.BadRequest, "Invalid UUID format")
		return
	}

	err := h.adminController.DeleteCategory(c, category_id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) CreateSubCategory(c *gin.Context) {
	var req entities.SubCategory
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, constants.BadRequest)
		return
	}
	req.ID = uuid.NewString()

	err = h.adminController.CreateSubCategory(c, req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) GetSubCategory(c *gin.Context) {
	data, err := h.adminController.GetSubCategory(c)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, data)
}

func (h *Handler) UpdateSubCategory(c *gin.Context) {
	var req entities.SubCategory
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, constants.BadRequest)
		return
	}
	req.ID = c.Param("id")
	if !utils.IsValidUUID(req.ID) {
		h.handleResponse(c, http.BadRequest, "Invalid UUID format")
		return
	}
	req.UpdatedAt = time.Now()
	err = h.adminController.UpdateSubCategory(c, req)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
		return
	}
	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) DeleteSubCategory(c *gin.Context) {
	sub_category_id := c.Param("id")

	if !utils.IsValidUUID(sub_category_id) {
		h.handleResponse(c, http.BadRequest, "Invalid UUID format")
		return
	}

	err := h.adminController.DeleteCategory(c, sub_category_id)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, constants.InternelServError)
	}
	h.handleResponse(c, http.OK, constants.Success)
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var body entities.Product
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, gin.H{
			"message": "error in body parse",
			"errors":  handleBodyParseError(err),
		})
		return
	}

	productID := uuid.New()
	body.ProductID = productID
	for i := range body.ProductDetails {
		body.ProductDetails[i].ProductID = productID
	}
	for i := range body.Files {
		body.Files[i].AttachID = &productID
	}

	err = h.adminController.CreateProduct(c, body)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), "error in CreateProduct")
		return
	}

	h.handleResponse(c, http.OK, body)
}
