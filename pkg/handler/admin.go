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

func (h *Handler) checkUserReferral(c *gin.Context) {
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
	res, statusCode, err := h.services.CheckUserReferral(body.TeleId)
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
			"message": "успешное получение количества пользователей по рефералки администратора",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение количества пользователей по рефералки администратора",
			"result":  res,
		})
	}
}

func (h *Handler) getUserReferral(c *gin.Context) {
	type Body struct {
		TeleId     int64 `json:"tele_id"`
		TeleIdUser int64 `json:"tele_id_user"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.GetUserReferral(body.TeleId, body.TeleIdUser)
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
			"message": "успешное получение пользователя рефералки администратора",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешное получение пользователя рефералки администратора",
			"result":  res,
		})
	}
}

func (h *Handler) getUsersReferral(c *gin.Context) {
	type Body struct {
		TeleId int64 `json:"tele_id"`
		Limit  int   `json:"limit"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	res, statusCode, err := h.services.GetUsersReferral(body.TeleId, body.Limit)
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

func (h *Handler) checkIsPremium(c *gin.Context) {
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
	res, statusCode, err := h.services.CheckIsPremium(body.TeleId)
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
			"message": "пользователь со статусом премиум",
			"result":  res,
		})
	} else {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "пользователя без статуса премиум",
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

func (h *Handler) checkIsVerification(c *gin.Context) {
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
	res, statusCode, err := h.services.CheckIsVerification(body.TeleId)
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
			"message": "пользователь является верифицированным",
			"result":  res,
		})
	} else {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "пользователя не является верифицированным",
			"result":  res,
		})
	}
}

func (h *Handler) updateVerification(c *gin.Context) {
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
	statusCode, err := h.services.UpdateVerification(body.TeleId)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная изменение верификации пользователя",
	})
}

func (h *Handler) adminUpdateMinimPrice(c *gin.Context) {
	type Body struct {
		TeleId   int64   `json:"tele_id"`
		MinPrice float64 `json:"min_price"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.AdminUpdateMinimPrice(body.TeleId, body.MinPrice)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная изменение минималки у администратора",
	})
}

func (h *Handler) adminAddBalance(c *gin.Context) {
	type Body struct {
		TeleId    int64   `json:"tele_id"`
		NeedPrice float64 `json:"need_price"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.AdminAddBalance(body.TeleId, body.NeedPrice)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное пополнение баланса пользователя",
	})
}

func (h *Handler) adminChangeMinUser(c *gin.Context) {
	type Body struct {
		TeleId   int64   `json:"tele_id"`
		MinPrice float64 `json:"min_price"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.AdminChangeMinUser(body.TeleId, body.MinPrice)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное изменение минималки пользователя",
	})
}

func (h *Handler) adminChangeBalance(c *gin.Context) {
	type Body struct {
		TeleId    int64   `json:"tele_id"`
		NeedPrice float64 `json:"need_price"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.AdminChangeBalance(body.TeleId, body.NeedPrice)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное изменение баланса пользователя",
	})
}

func (h *Handler) checkIsBlockUser(c *gin.Context) {
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
	res, statusCode, err := h.services.CheckIsBlockUser(body.TeleId)
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
			"message": "пользователь является заблокированным",
			"result":  res,
		})
	} else {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "пользователя не является заблокированным",
			"result":  res,
		})
	}
}

func (h *Handler) adminBlockUser(c *gin.Context) {
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
	statusCode, err := h.services.AdminBlockUser(body.TeleId)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешная блокировка пользователя",
	})
}

func (h *Handler) checkIsVisibleName(c *gin.Context) {
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
	res, statusCode, err := h.services.CheckIsVisibleName(body.TeleId)
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
			"message": "имя администратора скрыто",
			"result":  res,
		})
	} else {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": "имя администратора не скрыто",
			"result":  res,
		})
	}
}

func (h *Handler) adminChangeVisibleName(c *gin.Context) {
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
	statusCode, err := h.services.AdminChangeVisibleName(body.TeleId)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное изменение видимости имя",
	})
}

func (h *Handler) adminBuyTokenUser(c *gin.Context) {
	type Body struct {
		TeleId          int64   `json:"tele_id"`
		TokenUid        string  `json:"token_uid"`
		PriceUser       float64 `json:"price_user"`
		UidPaymentEvent string  `json:"uid_payment_event"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	statusCode, err := h.services.AdminBuyTokenUser(body.TeleId, body.TokenUid, body.PriceUser, body.UidPaymentEvent)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, map[string]interface{}{
		"status":  statusCode,
		"message": "успешное покупка nft у пользователя",
	})
}

func (h *Handler) adminWithdrawApprove(c *gin.Context) {
	type Body struct {
		TeleId           int64  `json:"tele_id"`
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
	res, statusCode, err := h.services.AdminWithdrawApprove(body.TeleId, body.WithDrawEventUid)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if res {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешный вывод денег пользователя",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешный вывод денег пользователя",
			"result":  res,
		})
	}
}

func (h *Handler) adminWithdrawRefuse(c *gin.Context) {
	type Body struct {
		TeleId           int64  `json:"tele_id"`
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
	res, statusCode, err := h.services.AdminWithdrawRefuse(body.TeleId, body.WithDrawEventUid)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		})
		return
	}
	if res {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешный отказ в выводе денег пользователю",
			"result":  res,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешный отказ в выводе денег пользователю",
			"result":  res,
		})
	}
}
