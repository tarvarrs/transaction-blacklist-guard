package walletoperation

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(r gin.IRouter) {
	r.POST("/wallet-operation", h.Process)
}
