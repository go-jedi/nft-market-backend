package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

func (h *Handler) checkIsAdmin(c *gin.Context) {
	type Body struct {
		TeleId int64 `json:"tele_id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.CheckIsAdmin(body.TeleId)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if res {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "пользователь является администратором",
			"result":  res,
		})
	} else {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "пользователя не является администратором",
			"result":  res,
		})
	}
}

func (h *Handler) createReferral(c *gin.Context) {
	type Body struct {
		TeleId        int64  `json:"tele_id"`
		TeleName      string `json:"tele_name"`
		AdminReferral int64  `json:"admin_referral"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.CreateReferral(appl_row.ReferralCreate(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная создание рефералки",
	})
}

func (h *Handler) getUsersReferral(c *gin.Context) {
	type Body struct {
		TeleId int64 `json:"tele_id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.GetUsersReferral(body.TeleId)
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
			"message": "успешное получение пользователей рефералки администратора",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение пользователей рефералки администратора",
			"result":  res,
		})
	}
}

func (h *Handler) adminGetUserProfile(c *gin.Context) {
	type Body struct {
		TeleId int64 `json:"tele_id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.AdminGetUserProfile(body.TeleId)
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
			"message": "успешное получение профиля пользователя",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение пользователей профиля пользователя",
			"result":  res,
		})
	}
}

func (h *Handler) updatePremium(c *gin.Context) {
	type Body struct {
		TeleId int64 `json:"tele_id"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.UpdatePremium(body.TeleId)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная изменение премиума пользователя",
	})
}
