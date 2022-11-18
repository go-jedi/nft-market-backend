package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

func (h *Handler) registrationUser(c *gin.Context) {
	type Body struct {
		TeleId   int64  `json:"tele_id"`
		TeleName string `json:"tele_name"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.RegistrationUser(appl_row.UserCreate(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная регистрация пользователя",
	})
}

func (h *Handler) checkAuth(c *gin.Context) {
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
	res, err, statusCode := h.services.CheckAuth(body.TeleId)
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
			"message": "пользователь уже существует",
			"result":  res,
		})
	} else {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "пользователя не существует",
			"result":  res,
		})
	}
}

func (h *Handler) updateLanguage(c *gin.Context) {
	type Body struct {
		TeleId int64  `json:"tele_id"`
		Lang   string `json:"lang"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.UpdateLanguage(appl_row.UserUpdateLanguage(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная изменение языка пользователя",
	})
}

func (h *Handler) checkIsLanguage(c *gin.Context) {
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
	res, err, statusCode := h.services.CheckIsLanguage(body.TeleId)
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
			"message": "успешное получение выбранного языка пользователя",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение выбранного языка пользователя",
			"result":  res,
		})
	}
}

func (h *Handler) updateCurrency(c *gin.Context) {
	type Body struct {
		TeleId   int64  `json:"tele_id"`
		Currency string `json:"currency"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	err, statusCode := h.services.UpdateCurrency(appl_row.UserUpdateCurrency(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная изменение валюты пользователя",
	})
}

func (h *Handler) checkIsTerms(c *gin.Context) {
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
	res, err, statusCode := h.services.CheckIsTerms(body.TeleId)
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
			"message": "пользователь согласен с пользовательским соглашением",
			"result":  res,
		})
	} else {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "пользователь не согласен с пользовательским соглашением",
			"result":  res,
		})
	}
}

func (h *Handler) agreeTerms(c *gin.Context) {
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
	err, statusCode := h.services.AgreeTerms(body.TeleId)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "пользователь успешно согласился с пользовательским соглашением",
	})
}
