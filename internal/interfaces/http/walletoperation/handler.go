package walletoperation

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	app "github.com/tarvarrs/transaction-blacklist-guard/internal/application/walletoperation"
	domain "github.com/tarvarrs/transaction-blacklist-guard/internal/domain/walletoperation"
)

type Handler struct {
	service app.Service
}

func NewHandler(service app.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Process(c *gin.Context) {
	var req processRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	result, err := h.service.Process(c.Request.Context(), app.ProcessCommand{
		FromID: req.FromID,
		ToID:   req.ToID,
		Amount: req.Amount,
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrEmptyParticipantID):
			c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, errorResponse{Error: "internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, processResponse{
		Status: result.Status,
	})
}
