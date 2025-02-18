package handler

import (
	"net/http"
	"serinitystore/helper"
	"serinitystore/otp"
	"time"

	"github.com/gin-gonic/gin"
)

type otpHandler struct {
	service otp.Service
}

func NewOTPHandler(service otp.Service) *otpHandler {
	return &otpHandler{service}
}

func (h *otpHandler) SaveOTP(c *gin.Context) {
	var input otp.OTPRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Otp request failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.Otp = helper.GenerateOTP()
	input.Expiry = time.Now().Add(3 * time.Minute)

	newOtpRequest, err := h.service.SaveOTP(input)
	if err != nil {
		response := helper.APIResponse("Otp request failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("OTP has been sent to your email", http.StatusOK, "success", newOtpRequest)

	c.JSON(http.StatusOK, response)
}
