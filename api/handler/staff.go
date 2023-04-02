package handler

import (
	"app/api/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Staff godoc
// @ID create_staff
// @Router /staff [POST]
// @Summary Create Staff
// @Description Create Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param staff body models.CreateStaff true "CreateStaffRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateStaff(c *gin.Context) {

	var createStaff models.CreateStaff

	err := c.ShouldBindJSON(&createStaff) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create staff", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Staff().Create(context.Background(), &createStaff)
	if err != nil {
		h.handlerResponse(c, "storage.staff.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{StaffId: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create staff", http.StatusCreated, resp)
}

// Get By ID Staff godoc
// @ID get_by_id_staff
// @Router /staff/{id} [GET]
// @Summary Get By ID Staff
// @Description Get By ID Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdStaff(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{StaffId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get staff by id", http.StatusCreated, resp)
}

// Get List Staff godoc
// @ID get_list_staff
// @Router /staff [GET]
// @Summary Get List Staff
// @Description Get List Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListStaff(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list staff", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list staff", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Staff().GetList(context.Background(), &models.GetListStaffRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list staff response", http.StatusOK, resp)
}

// Get List Staff godoc
// @ID get_list_staff_report
// @Router /staffreport [GET]
// @Summary Get List Staff Report
// @Description Get List Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListReportStaff(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list staff", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list staff", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Staff().GetListReport(context.Background(), &models.GetListReportStaffRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		fmt.Println("xato")
		h.handlerResponse(c, "storage.staffreport.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list staff response", http.StatusOK, resp)
}

// Update Staff godoc
// @ID update_staff
// @Router /staff/{id} [PUT]
// @Summary Update Staff
// @Description Update Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param staff body models.UpdateStaff true "UpdateStaffRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateStaff(c *gin.Context) {

	var updateStaff models.UpdateStaff

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateStaff)
	if err != nil {
		h.handlerResponse(c, "update staff", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	updateStaff.StaffId = idInt

	rowsAffected, err := h.storages.Staff().UpdatePut(context.Background(), &updateStaff)
	if err != nil {
		h.handlerResponse(c, "storage.staff.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.staff.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{StaffId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update staff", http.StatusAccepted, resp)
}

// Update Patch Staff godoc
// @ID update_staff
// @Router /staff/{id} [PATCH]
// @Summary Update PATCH Staff
// @Description Update PATCH Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param staff body models.PatchRequest true "UpdatePatchRequest"
// @Success 202 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchStaff(c *gin.Context) {

	var obj models.PatchRequest

	id := c.Param("id")

	err := c.ShouldBindJSON(&obj)
	if err != nil {
		h.handlerResponse(c, "update staff", http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	obj.ID = idInt

	rowsAffected, err := h.storages.Staff().UpdatePatch(context.Background(), &obj)
	if err != nil {
		h.handlerResponse(c, "storage.staff.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.staff.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Staff().GetByID(context.Background(), &models.StaffPrimaryKey{StaffId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update staff", http.StatusAccepted, resp)
}

// DELETE Staff godoc
// @ID delete_staff
// @Router /staff/{id} [DELETE]
// @Summary Delete Staff
// @Description Delete Staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param staff body models.StaffPrimaryKey true "DeleteStaffRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteStaff(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.staff.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Staff().Delete(context.Background(), &models.StaffPrimaryKey{StaffId: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.staff.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.staff.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete staff", http.StatusNoContent, nil)
}

