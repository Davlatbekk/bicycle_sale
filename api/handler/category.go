package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Category godoc
// @ID create_category
// @Router /category [POST]
// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.CreateCategory true "CreateCategoryRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCategory(c *gin.Context) {

	var createCategory models.CreateCategory

	err := c.ShouldBindJSON(&createCategory) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create category", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Category().Create(context.Background(), &createCategory)
	if err != nil {
		h.handlerResponse(c, "storage.category.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{CategoryId: id})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create category", http.StatusCreated, resp)
}

// Get By ID Category godoc
// @ID get_by_id_category
// @Router /category/{id} [GET]
// @Summary Get By ID Category
// @Description Get By ID Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCategory(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{CategoryId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get category by id", http.StatusCreated, resp)
}

// Get List Category godoc
// @ID get_list_category
// @Router /category [GET]
// @Summary Get List Category
// @Description Get List Category
// @Tags Category
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCategory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list category", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list category", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Category().GetList(context.Background(), &models.GetListCategoryRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.category.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list category response", http.StatusOK, resp)
}

// Update Category godoc
// @ID update_category
// @Router /category/{id} [PUT]
// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param category body models.UpdateCategory true "UpdateCategoryRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCategory(c *gin.Context) {

	var updateCategory models.UpdateCategory

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateCategory)
	if err != nil {
		h.handlerResponse(c, "update category", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	updateCategory.CategoryId = idInt

	rowsAffected, err := h.storages.Category().Update(context.Background(), &updateCategory)
	if err != nil {
		h.handlerResponse(c, "storage.category.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.category.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{CategoryId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update category", http.StatusAccepted, resp)
}

// DELETE Category godoc
// @ID delete_category
// @Router /category/{id} [DELETE]
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param category body models.CategoryPrimaryKey true "DeleteCategoryRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCategory(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.category.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Category().Delete(context.Background(), &models.CategoryPrimaryKey{CategoryId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.category.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.category.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete category", http.StatusNoContent, nil)
}
