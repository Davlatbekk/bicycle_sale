package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Order godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.CreateOrder true "CreateOrderRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateOrder(c *gin.Context) {

	var createOrder models.CreateOrder

	err := c.ShouldBindJSON(&createOrder) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create order", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Order().Create(context.Background(), &createOrder)
	if err != nil {
		h.handlerResponse(c, "storage.order.create", http.StatusInternalServerError, err.Error())
		return
	}

	// resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{OrderId: id})
	// if err != nil {
	// 	h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
	// 	return
	// }

	h.handlerResponse(c, "create order", http.StatusCreated, id)
}

// Get By ID Order godoc
// @ID get_by_id_order
// @Router /order/{id} [GET]
// @Summary Get By ID Order
// @Description Get By ID Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdOrder(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{OrderId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get order by id", http.StatusCreated, resp)
}

// Get List Order godoc
// @ID get_list_order
// @Router /order [GET]
// @Summary Get List Order
// @Description Get List Order
// @Tags Order
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListOrder(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list order", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list order", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Order().GetList(context.Background(), &models.GetListOrderRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.order.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list order response", http.StatusOK, resp)
}

// Update Order godoc
// @ID update_order
// @Router /order/{id} [PUT]
// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param order body models.UpdateOrder true "UpdateOrderRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateOrder(c *gin.Context) {

	var updateOrder models.UpdateOrder

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateOrder)
	if err != nil {
		h.handlerResponse(c, "update order", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	updateOrder.OrderId = idInt

	rowsAffected, err := h.storages.Order().Update(context.Background(), &updateOrder)
	if err != nil {
		h.handlerResponse(c, "storage.order.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.order.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{OrderId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update order", http.StatusAccepted, resp)
}

// Update Patch Order godoc
// @ID update_order
// @Router /order/{id} [PATCH]
// @Summary Update PATCH Order
// @Description Update PATCH Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param order body models.PatchRequest true "UpdatePatchRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchOrder(c *gin.Context) {

	var obj models.PatchRequest

	id := c.Param("id")

	err := c.ShouldBindJSON(&obj)
	if err != nil {
		h.handlerResponse(c, "update order", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	obj.ID = idInt

	rowsAffected, err := h.storages.Order().UpdatePatch(context.Background(), &obj)
	if err != nil {
		h.handlerResponse(c, "storage.order.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.order.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{OrderId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update order", http.StatusAccepted, resp)
}

// DELETE Order godoc
// @ID delete_order
// @Router /order/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param order body models.OrderPrimaryKey true "DeleteOrderRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteOrder(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Order().Delete(context.Background(), &models.OrderPrimaryKey{OrderId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.order.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.order.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete order", http.StatusNoContent, nil)
}

// -------------------------------------------------------------------------------------------
// Create Order Item godoc
// @ID create_order_item
// @Router /order_item [POST]
// @Summary Create Order Item
// @Description Create Order Item
// @Tags Order
// @Accept json
// @Produce json
// @Param order_item body models.CreateOrderItem true "CreateOrderItemRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateOrderItem(c *gin.Context) {

	var createOrderItem models.CreateOrderItem

	err := c.ShouldBindJSON(&createOrderItem) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create order", http.StatusBadRequest, err.Error())
		return
	}

	err = h.storages.Order().AddOrderItem(context.Background(), &createOrderItem)
	if err != nil {
		h.handlerResponse(c, "storage.order.create", http.StatusInternalServerError, err.Error())
		return
	}

	// resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{OrderId: id})
	// if err != nil {
	// 	h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
	// 	return
	// }

	h.handlerResponse(c, "create order", http.StatusCreated, "Order Item Added")
}

// DELETE Order Item godoc
// @ID delete_order_item
// @Router /order_item/{id} [DELETE]
// @Summary Delete Order Item
// @Description Delete Order Item
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param item_id query string true "item_id"
// @Param orderItem body models.OrderItemPrimaryKey true "DeleteOrderItemRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteOrderItem(c *gin.Context) {

	id := c.Param("id")
	itemId := c.Query("item_id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	idItemInt, err := strconv.Atoi(itemId)
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	err = h.storages.Order().RemoveOrderItem(context.Background(), &models.OrderItemPrimaryKey{OrderId: idInt, ItemId: idItemInt})
	if err != nil {
		h.handlerResponse(c, "storage.order.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete order", http.StatusNoContent, "Deleted succesfully")
}

// Total Sum Order godoc
// @ID total_sum_order
// @Router /order/total_sum [GET]
// @Summary Total Sum Order
// @Description Total Sum Order
// @Tags Order
// @Accept json
// @Produce json
// @Param order_id query string true "order_id"
// @Param promocode_name query string false "promocode_name"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) OrderTotalSum(c *gin.Context) {
	var orderSum models.OrderTotalSum

	orderId, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		h.handlerResponse(c, "Atoi error order total sum", http.StatusBadRequest, err.Error())
		return
	}

	orderSum.OrderId = orderId
	orderSum.PromocodeName = c.Query("promocode_name")

	totalSum, err := h.storages.Order().OrderTotalSum(context.Background(), &orderSum)
	if err != nil {
		h.handlerResponse(c, "Storage order total sum", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Order total sum", http.StatusOK, totalSum)
}
