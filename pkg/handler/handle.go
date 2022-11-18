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
		api.POST("/user/registration", h.registrationUser)   // Регистрация пользователя
		api.POST("/user/checkAuth", h.checkAuth)             // проверка на регистрацию пользователя
		api.POST("/user/updateLanguage", h.updateLanguage)   // Изменение языка пользователя
		api.POST("/user/checkIsLanguage", h.checkIsLanguage) // Проверка языка пользователя
		api.POST("/user/updateCurrency", h.updateCurrency)   // Изменить валюту пользователя
		api.POST("/user/checkIsTerms", h.checkIsTerms)       // проверка на пользовательское соглашение
		api.POST("/user/agreeTerms", h.agreeTerms)           // согласие пользователя с пользовательским соглашением
	}

	return router
}
