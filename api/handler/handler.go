package handler

import (
	"app/config"
	"app/pkg/logger"
	"app/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg      *config.Config
	logger   logger.LoggerI
	storages storage.StorageI
}

type Response struct {
	Status      int
	Description string
	Data        interface{}
}

func NewHandler(cfg *config.Config, store storage.StorageI, logger logger.LoggerI) *Handler {
	return &Handler{
		cfg:      cfg,
		logger:   logger,
		storages: store,
	}
}

func (h *Handler) handlerResponse(c *gin.Context, path string, code int, message interface{}) {
	response := Response{
		Status:      code,
		Description: path,
		Data:        message,
	}

	switch {
	case code < 300:
		// h.logger.Info(path, logger.Any("info", response))
	case code >= 400:
		h.logger.Error(path, logger.Any("info", response))
	}

	c.JSON(code, response)
}

func (h *Handler) getOffsetQuery(offset string) (int, error) {
	if len(offset) <= 0 {
		return h.cfg.DefaultOffset, nil
	}

	return strconv.Atoi(offset)
}

func (h *Handler) getLimitQuery(limit string) (int, error) {

	if len(limit) <= 0 {
		return h.cfg.DefaultLimit, nil
	}

	return strconv.Atoi(limit)
}
