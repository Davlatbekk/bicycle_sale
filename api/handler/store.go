package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Store godoc
// @ID create_store
// @Router /store [POST]
// @Summary Create Store
// @Description Create Store
// @Tags Store
// @Accept json
// @Produce json
// @Param store body models.CreateStore true "CreateStoreRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateStore(c *gin.Context) {

	var createStore models.CreateStore

	err := c.ShouldBindJSON(&createStore) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create store", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Store().Create(context.Background(), &createStore)
	if err != nil {
		h.handlerResponse(c, "storage.store.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{StoreId: id})
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create store", http.StatusCreated, resp)
}

// Get By ID Store godoc
// @ID get_by_id_store
// @Router /store/{id} [GET]
// @Summary Get By ID Store
// @Description Get By ID Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdStore(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{StoreId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get store by id", http.StatusCreated, resp)
}

// Get List Store godoc
// @ID get_list_store
// @Router /store [GET]
// @Summary Get List Store
// @Description Get List Store
// @Tags Store
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListStore(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list store", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list store", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Store().GetList(context.Background(), &models.GetListStoreRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.store.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list store response", http.StatusOK, resp)
}

// Update Store godoc
// @ID update_store
// @Router /store/{id} [PUT]
// @Summary Update Store
// @Description Update Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param store body models.UpdateStore true "UpdateStoreRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateStore(c *gin.Context) {

	var updateStore models.UpdateStore

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateStore)
	if err != nil {
		h.handlerResponse(c, "update store", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	updateStore.StoreId = idInt

	rowsAffected, err := h.storages.Store().UpdatePut(context.Background(), &updateStore)
	if err != nil {
		h.handlerResponse(c, "storage.store.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.store.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{StoreId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update store", http.StatusAccepted, resp)
}

// Update Patch Store godoc
// @ID update_store
// @Router /store/{id} [PATCH]
// @Summary Update PATCH Store
// @Description Update PATCH Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param store body models.PatchRequest true "UpdatePatchRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchStore(c *gin.Context) {

	var obj models.PatchRequest

	id := c.Param("id")

	err := c.ShouldBindJSON(&obj)
	if err != nil {
		h.handlerResponse(c, "update store", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	obj.ID = idInt

	rowsAffected, err := h.storages.Store().UpdatePatch(context.Background(), &obj)
	if err != nil {
		h.handlerResponse(c, "storage.store.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.store.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{StoreId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update store", http.StatusAccepted, resp)
}

// DELETE Store godoc
// @ID delete_store
// @Router /store/{id} [DELETE]
// @Summary Delete Store
// @Description Delete Store
// @Tags Store
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param store body models.StorePrimaryKey true "DeleteStoreRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteStore(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.store.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Store().Delete(context.Background(), &models.StorePrimaryKey{StoreId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.store.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.store.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete store", http.StatusNoContent, nil)
}
