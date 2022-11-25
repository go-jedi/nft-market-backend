package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

func (h *Handler) createDepot(c *gin.Context) {
	type Body struct {
		MammothId       int64   `json:"mammoth_id"`
		MammothUsername string  `json:"mammoth_username"`
		WorkerId        int64   `json:"worker_id"`
		WorkerUsername  string  `json:"worker_username"`
		Amount          float64 `json:"amount"`
		IsShowName      bool    `json:"is_show_name"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.CreateDepot(appl_row.DepotCreate(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная создание депа",
	})
}
