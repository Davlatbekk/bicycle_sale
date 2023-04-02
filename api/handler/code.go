package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Code godoc
// @ID create_code
// @Router /code [POST]
// @Summary Create Code
// @Description Create Code
// @Tags Code
// @Accept json
// @Produce json
// @Param store body models.CreateCode true "CreateCodeRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCode(c *gin.Context) {

	var createCode models.CreateCode

	err := c.ShouldBindJSON(&createCode) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create store", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Code().Create(context.Background(), &createCode)
	if err != nil {
		h.handlerResponse(c, "storage.code.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Code().GetByID(context.Background(), &models.CodePrimaryKey{Code_Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.code.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create code", http.StatusCreated, resp)
}

// Get By ID Code godoc
// @ID get_by_id_code
// @Router /code/{id} [GET]
// @Summary Get By ID Code
// @Description Get By ID Code
// @Tags Code
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCode(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.code.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Code().GetByID(context.Background(), &models.CodePrimaryKey{Code_Id: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.code.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get store by id", http.StatusCreated, resp)
}

// Get List Code godoc
// @ID get_list_code
// @Router /code [GET]
// @Summary Get List Code
// @Description Get List Code
// @Tags Code
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCode(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list code", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list code", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Code().GetList(context.Background(), &models.GetListCodeRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.code.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list code response", http.StatusOK, resp)
}

// Update Code godoc
// @ID update_code
// @Router /code/{id} [PUT]
// @Summary Update Code
// @Description Update Code
// @Tags Code
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param store body models.UpdateCode true "UpdateCodeRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCode(c *gin.Context) {

	var updateCode models.UpdateCode

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateCode)
	if err != nil {
		h.handlerResponse(c, "update code", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.code.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	updateCode.Code_Id = idInt

	rowsAffected, err := h.storages.Code().Update(context.Background(), &updateCode)
	if err != nil {
		h.handlerResponse(c, "storage.code.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.code.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Code().GetByID(context.Background(), &models.CodePrimaryKey{Code_Id: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.code.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update code", http.StatusAccepted, resp)
}

// DELETE Code godoc
// @ID delete_code
// @Router /code/{id} [DELETE]
// @Summary Delete Code
// @Description Delete Code
// @Tags Code
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param store body models.CodePrimaryKey true "DeleteCodeRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCode(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.code.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Code().Delete(context.Background(), &models.CodePrimaryKey{Code_Id: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.code.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.code.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete code", http.StatusNoContent, nil)
}