package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rob-bender/nft-market-backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api-v1")
	{
		api.POST("/user/checkAuth", h.checkAuth)                                 // проверка на регистрацию пользователя
		api.POST("/user/registration", h.registrationUser)                       // Регистрация пользователя
		api.POST("/user/getUserLanguage", h.getUserLanguage)                     // Получить выбранный язык пользователя
		api.POST("/user/updateLanguage", h.updateLanguage)                       // Изменение языка пользователя
		api.POST("/user/getUserCurrency", h.getUserCurrency)                     // Получить выбранную валюту пользователя
		api.POST("/user/updateCurrency", h.updateCurrency)                       // Изменить валюту пользователя
		api.POST("/user/checkIsTerms", h.checkIsTerms)                           // проверка на пользовательское соглашение
		api.POST("/user/agreeTerms", h.agreeTerms)                               // согласие пользователя с пользовательским соглашением
		api.POST("/user/getUserProfile", h.getUserProfile)                       // получение профиля пользователя
		api.GET("/payment/getAll", h.getAllPayments)                             // получение всех платёжек
		api.POST("/payment/createPayment", h.createPayment)                      // создание платёжки
		api.GET("/collection/getAll", h.getAllCollections)                       // получение всех коллекций
		api.POST("/collection/createCollection", h.createCollection)             // создание коллекции
		api.POST("/collection/createToken", h.createToken)                       // создание токена для коллекции
		api.POST("/collection/getAllTokensCollection", h.getAllTokensCollection) // получение всех токенов коллекции
		api.POST("/collection/getToken", h.getToken)                             // получение токена
		api.POST("/admin/checkIsAdmin", h.checkIsAdmin)                          //  проверка на администратора
		api.POST("/admin/createReferral", h.createReferral)                      // создание рефералки
		api.POST("/admin/getUserReferral", h.getUsersReferral)                   // получение всех пользователей по рефералке администратора
		api.POST("/admin/getUserProfile", h.adminGetUserProfile)                 // получение профиля пользователя
		api.POST("/admin/updatePremium", h.updatePremium)                        // изменить премиум
	}

	return router
}
