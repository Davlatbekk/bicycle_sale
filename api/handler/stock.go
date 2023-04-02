package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Stock godoc
// @ID create_stock
// @Router /stock [POST]
// @Summary Create Stock
// @Description Create Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param stock body models.CreateStock true "CreateStockRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateStock(c *gin.Context) {

	var createStock models.CreateStock

	err := c.ShouldBindJSON(&createStock) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create product", http.StatusBadRequest, err.Error())
		return
	}

	storeId, _, err := h.storages.Stock().Create(context.Background(), &createStock)
	if err != nil {
		h.handlerResponse(c, "storage.stock.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Stock().GetByID(context.Background(), &models.StockPrimaryKey{StoreId: storeId})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create stock", http.StatusCreated, resp)
}

// Get By ID Stock godoc
// @ID get_by_id_stock
// @Router /stock/{id} [GET]
// @Summary Get By ID Stock
// @Description Get By ID Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdStock(c *gin.Context) {

	var err error

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Stock().GetByID(context.Background(), &models.StockPrimaryKey{StoreId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get stock by id", http.StatusCreated, resp)
}

// Get List Stock godoc
// @ID get_list_stock
// @Router /stock [GET]
// @Summary Get List Stock
// @Description Get List Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListStock(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list stock", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list stock", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Stock().GetList(context.Background(), &models.GetListStockRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list stock response", http.StatusOK, resp)
}

// Update Stock godoc
// @ID update_stock
// @Router /stock/{id} [PUT]
// @Summary Update Stock
// @Description Update Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param stock body models.UpdateStock true "UpdateStockRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateStock(c *gin.Context) {

	var updateStock models.UpdateStock

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateStock)
	if err != nil {
		h.handlerResponse(c, "update stock", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	updateStock.StoreId = idInt

	rowsAffected, err := h.storages.Stock().Update(context.Background(), &updateStock)
	if err != nil {
		h.handlerResponse(c, "storage.stock.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.stock.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Stock().GetByID(context.Background(), &models.StockPrimaryKey{StoreId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update stock", http.StatusAccepted, resp)
}

// DELETE Stock godoc
// @ID delete_stock
// @Router /stock/{id} [DELETE]
// @Summary Delete Stock
// @Description Delete Stock
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param stock body models.StockPrimaryKey true "DeleteStockRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteStock(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Stock().Delete(context.Background(), &models.StockPrimaryKey{StoreId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.stock.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.stock.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete stock", http.StatusNoContent, nil)
}

// Send Product godoc
// @ID send_product
// @Router /stock/send_product [PUT]
// @Summary Send Product
// @Description Send Product To Another Store
// @Tags Stock
// @Accept json
// @Produce json
// @Param report body models.SendProduct true "SendProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) SendProductToStore(c *gin.Context) {
	var sendProduct models.SendProduct

	err := c.ShouldBindJSON(&sendProduct)
	if err != nil {
		h.handlerResponse(c, "Bind Json error send product to store", http.StatusBadRequest, err.Error())
		return
	}

	err = h.storages.Stock().SendProduct(context.Background(), &sendProduct)
	if err != nil {
		h.handlerResponse(c, "Storage report  send product", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Get store by id", http.StatusOK, "Success")
}
