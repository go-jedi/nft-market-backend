package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nft-market-backend/appl_row"
)

func (h *Handler) exchangeRates(c *gin.Context) {
	res, statusCode, err := h.services.GetAllExchangeRates()
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
			"message": "успешное получение обменных курсов",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение обменных курсов",
			"result":  res,
		})
	}
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
	res, statusCode, err := h.services.CheckAuth(body.TeleId)
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

func (h *Handler) registrationUser(c *gin.Context) {
	type Body struct {
		TeleId      int64  `json:"tele_id"`
		TeleName    string `json:"tele_name"`
		TeleIdAdmin int64  `json:"tele_id_admin"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.RegistrationUser(appl_row.UserCreate(body))
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

func (h *Handler) getUserLanguage(c *gin.Context) {
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
	res, statusCode, err := h.services.GetUserLanguage(body.TeleId)
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
	statusCode, err := h.services.UpdateLanguage(appl_row.UserUpdateLanguage(body))
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

func (h *Handler) getUserCurrency(c *gin.Context) {
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
	res, statusCode, err := h.services.GetUserCurrency(body.TeleId)
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
			"message": "успешное получение выбранной валюты пользователя",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение выбранной валюты пользователя",
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
	statusCode, err := h.services.UpdateCurrency(appl_row.UserUpdateCurrency(body))
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
	res, statusCode, err := h.services.CheckIsTerms(body.TeleId)
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
	statusCode, err := h.services.AgreeTerms(body.TeleId)
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

func (h *Handler) getUserProfile(c *gin.Context) {
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
	res, statusCode, err := h.services.GetUserProfile(body.TeleId)
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
			"message": "успешное получение профиля пользователя",
			"result":  res,
		})
	}
}

func (h *Handler) getUserMinPrice(c *gin.Context) {
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
	res, statusCode, err := h.services.GetUserMinPrice(body.TeleId)
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
			"message": "успешное получение минималки пользователя",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение минималки пользователя",
			"result":  res,
		})
	}
}

func (h *Handler) getAdminByUser(c *gin.Context) {
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
	res, statusCode, err := h.services.GetAdminByUser(body.TeleId)
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
			"message": "успешное получение закрепленного за пользователем администратора",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение закрепленного за пользователем администратора",
			"result":  res,
		})
	}
}

func (h *Handler) getUserBalance(c *gin.Context) {
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
	res, statusCode, err := h.services.GetUserBalance(body.TeleId)
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
			"message": "успешное получение баланса пользователя",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение баланса пользователя",
			"result":  res,
		})
	}
}

func (h *Handler) checkUserToken(c *gin.Context) {
	type Body struct {
		TeleId   int64  `json:"tele_id"`
		TokenUid string `json:"token_uid"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.CheckUserToken(body.TeleId, body.TokenUid)
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
			"message": "у пользователя имеется данный токен",
			"result":  res,
		})
	} else {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "у пользователя нет данного токена",
			"result":  res,
		})
	}
}

func (h *Handler) buyUserToken(c *gin.Context) {
	type Body struct {
		TeleId     int64   `json:"tele_id"`
		TokenUid   string  `json:"token_uid"`
		TokenPrice float64 `json:"token_price"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.BuyUserToken(appl_row.UserBuyToken(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная покупка токена пользователем",
	})
}

func (h *Handler) sellUserToken(c *gin.Context) {
	type Body struct {
		TeleId     int64   `json:"tele_id"`
		TokenUid   string  `json:"token_uid"`
		TokenPrice float64 `json:"token_price"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.SellUserToken(appl_row.UserSellToken(body))
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная продажа токена пользователем",
		"result":  res,
	})
}

func (h *Handler) getUserNft(c *gin.Context) {
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
	res, statusCode, err := h.services.GetUserNft(body.TeleId)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "успешное получение nft пользователя",
		"result":  res,
	})
}

func (h *Handler) getUserPaymentEvent(c *gin.Context) {
	type Body struct {
		EventUid string `json:"event_uid"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.GetUserPaymentEvent(body.EventUid)
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
			"message": "успешное получение события продажи",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение события продажи",
			"result":  res,
		})
	}
}

func (h *Handler) createWithDrawEvent(c *gin.Context) {
	type Body struct {
		TeleId int64   `json:"tele_id"`
		Price  float64 `json:"price"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.CreateWithDrawEvent(appl_row.UserWithDrawEventCreate(body))
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
			"message": "успешное создание события вывода",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное создание события вывода",
			"result":  res,
		})
	}
}

func (h *Handler) getWithDrawEvent(c *gin.Context) {
	type Body struct {
		WithDrawEventUid string `json:"with_draw_event_uid"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.GetWithDrawEvent(body.WithDrawEventUid)
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
			"message": "успешное получение события вывода",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение события вывода",
			"result":  res,
		})
	}
}
