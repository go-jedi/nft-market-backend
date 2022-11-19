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
		api.POST("/user/checkAuth", h.checkAuth)             // проверка на регистрацию пользователя
		api.POST("/user/registration", h.registrationUser)   // Регистрация пользователя
		api.POST("/user/getUserLanguage", h.getUserLanguage) // Получить выбранный язык пользователя
		api.POST("/user/updateLanguage", h.updateLanguage)   // Изменение языка пользователя
		api.POST("/user/getUserCurrency", h.getUserCurrency) // Получить выбранную валюту пользователя
		api.POST("/user/updateCurrency", h.updateCurrency)   // Изменить валюту пользователя
		api.POST("/user/checkIsTerms", h.checkIsTerms)       // проверка на пользовательское соглашение
		api.POST("/user/agreeTerms", h.agreeTerms)           // согласие пользователя с пользовательским соглашением
		api.POST("/user/getUserProfile", h.getUserProfile)   // получение профиля пользователя
		api.POST("payment/createPayment", h.createPayment)   // создание платёжки
		api.GET("/payment/getAll", h.getAllPayments)         // получение всех платёжек
	}

	return router
}
