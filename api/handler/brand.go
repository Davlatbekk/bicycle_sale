package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Brand godoc
// @ID create_brand
// @Router /brand [POST]
// @Summary Create Brand
// @Description Create Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param brand body models.CreateBrand true "CreateBrandRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateBrand(c *gin.Context) {

	var createBrand models.CreateBrand

	err := c.ShouldBindJSON(&createBrand) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create brand", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Brand().Create(context.Background(), &createBrand)
	if err != nil {
		h.handlerResponse(c, "storage.brand.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Brand().GetByID(context.Background(), &models.BrandPrimaryKey{BrandId: id})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create brand", http.StatusCreated, resp)
}

// Get By ID Brand godoc
// @ID get_by_id_brand
// @Router /brand/{id} [GET]
// @Summary Get By ID Brand
// @Description Get By ID Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdBrand(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Brand().GetByID(context.Background(), &models.BrandPrimaryKey{BrandId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get brand by id", http.StatusCreated, resp)
}

// Get List Brand godoc
// @ID get_list_brand
// @Router /brand [GET]
// @Summary Get List Brand
// @Description Get List Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListBrand(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list brand", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list brand", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Brand().GetList(context.Background(), &models.GetListBrandRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list brand response", http.StatusOK, resp)
}

// Update Brand godoc
// @ID update_brand
// @Router /brand/{id} [PUT]
// @Summary Update Brand
// @Description Update Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param brand body models.UpdateBrand true "UpdateBrandRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateBrand(c *gin.Context) {

	var updateBrand models.UpdateBrand

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateBrand)
	if err != nil {
		h.handlerResponse(c, "update brand", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	updateBrand.BrandId = idInt

	rowsAffected, err := h.storages.Brand().Update(context.Background(), &updateBrand)
	if err != nil {
		h.handlerResponse(c, "storage.brand.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.brand.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Brand().GetByID(context.Background(), &models.BrandPrimaryKey{BrandId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update brand", http.StatusAccepted, resp)
}

// DELETE Brand godoc
// @ID delete_brand
// @Router /brand/{id} [DELETE]
// @Summary Delete Brand
// @Description Delete Brand
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param brand body models.BrandPrimaryKey true "DeleteBrandRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteBrand(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Brand().Delete(context.Background(), &models.BrandPrimaryKey{BrandId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.brand.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.brand.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete brand", http.StatusNoContent, nil)
}
