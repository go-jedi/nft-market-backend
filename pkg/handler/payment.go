package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

func (h *Handler) createPayment(c *gin.Context) {
	type Body struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.CreatePayment(appl_row.PaymentCreate(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное создание типа оплаты",
	})
}

func (h *Handler) getAllPayments(c *gin.Context) {
	res, statusCode, err := h.services.GetAllPayments()
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if len(res) > 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение имеющихся платёжек",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение имеющихся платёжек",
			"result":  []int{},
		})
	}
}
