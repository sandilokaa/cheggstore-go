package routes

import (
	"serinitystore/auth"
	"serinitystore/handler"
	"serinitystore/user"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.RouterGroup, userService user.Service, authService auth.Service) {
	userHandler := handler.NewHandler(userService, authService)

	router.POST("/register", userHandler.RegisterUser)
	router.POST("/sessions", userHandler.LoginUser)
	router.GET("/sessions/oauth", userHandler.GetLoginGoogleURL)
	router.GET("/sessions/oauth/callback", userHandler.CallbackHandler)
	router.POST("/forgot-passwords", userHandler.OtpRequest)
	router.PUT("/forgot-passwords", userHandler.UpdateIsVerifiedOTP)
}
