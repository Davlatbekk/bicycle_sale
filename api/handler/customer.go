package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Customer godoc
// @ID create_customer
// @Router /customer [POST]
// @Summary Create Customer
// @Description Create Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body models.CreateCustomer true "CreateCustomerRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCustomer(c *gin.Context) {

	var createCustomer models.CreateCustomer

	err := c.ShouldBindJSON(&createCustomer) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create customer", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Customer().Create(context.Background(), &createCustomer)
	if err != nil {
		h.handlerResponse(c, "storage.customer.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{CustomerId: id})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create customer", http.StatusCreated, resp)
}

// Get By ID Customer godoc
// @ID get_by_id_customer
// @Router /customer/{id} [GET]
// @Summary Get By ID Customer
// @Description Get By ID Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCustomer(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{CustomerId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get customer by id", http.StatusCreated, resp)
}

// Get List Customer godoc
// @ID get_list_customer
// @Router /customer [GET]
// @Summary Get List Customer
// @Description Get List Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCustomer(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list customer", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list customer", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Customer().GetList(context.Background(), &models.GetListCustomerRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list customer response", http.StatusOK, resp)
}

// Update Customer godoc
// @ID update_customer
// @Router /customer/{id} [PUT]
// @Summary Update Customer
// @Description Update Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param customer body models.UpdateCustomer true "UpdateCustomerRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCustomer(c *gin.Context) {

	var updateCustomer models.UpdateCustomer

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateCustomer)
	if err != nil {
		h.handlerResponse(c, "update customer", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	updateCustomer.CustomerId = idInt

	rowsAffected, err := h.storages.Customer().UpdatePut(context.Background(), &updateCustomer)
	if err != nil {
		h.handlerResponse(c, "storage.customer.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.customer.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{CustomerId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update customer", http.StatusAccepted, resp)
}

// Update Patch Customer godoc
// @ID update_customer
// @Router /customer/{id} [PATCH]
// @Summary Update PATCH Customer
// @Description Update PATCH Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param customer body models.PatchRequest true "UpdatePatchRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchCustomer(c *gin.Context) {

	var obj models.PatchRequest

	id := c.Param("id")

	err := c.ShouldBindJSON(&obj)
	if err != nil {
		h.handlerResponse(c, "update customer", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	obj.ID = idInt

	rowsAffected, err := h.storages.Customer().UpdatePatch(context.Background(), &obj)
	if err != nil {
		h.handlerResponse(c, "storage.customer.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.customer.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Customer().GetByID(context.Background(), &models.CustomerPrimaryKey{CustomerId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update customer", http.StatusAccepted, resp)
}

// DELETE Customer godoc
// @ID delete_customer
// @Router /customer/{id} [DELETE]
// @Summary Delete Customer
// @Description Delete Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param customer body models.CustomerPrimaryKey true "DeleteCustomerRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCustomer(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.customer.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Customer().Delete(context.Background(), &models.CustomerPrimaryKey{CustomerId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.customer.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.customer.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete customer", http.StatusNoContent, nil)
}
