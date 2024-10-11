package handler

import (
	"cheggstore/cloth"
	"cheggstore/helper"
	"cheggstore/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type clothHandler struct {
	service cloth.Service
}

func NewClothHandler(service cloth.Service) *clothHandler {
	return &clothHandler{service}
}

func (h *clothHandler) SaveCloth(c *gin.Context) {
	var input cloth.CreateClothInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create cloth", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCloth, err := h.service.SaveCloth(input)
	if err != nil {
		response := helper.APIResponse("Failed to create cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create cloth", http.StatusOK, "success", cloth.FormatCloth(newCloth))
	c.JSON(http.StatusOK, response)
}

func (h *clothHandler) FindAllCloth(c *gin.Context) {

	search := c.Query("search")

	cloths, err := h.service.FindAllCloth(search)
	if err != nil {
		response := helper.APIResponse("Failed to find cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to find cloth", http.StatusOK, "success", cloth.FormatCloths(cloths))
	c.JSON(http.StatusOK, response)
}

func (h *clothHandler) FindClothByID(c *gin.Context) {
	var input cloth.ClothInputDetail

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cloth, err := h.service.FindClothByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to find cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to find cloth", http.StatusOK, "success", cloth.FormatClothDetail(cloth))
	c.JSON(http.StatusOK, response)
}

func (h *clothHandler) UpdateClothByID(c *gin.Context) {
	var inputID cloth.ClothInputDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get supplier", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData cloth.UpdateClothInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create cloth", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	oldCloth, err := h.service.FindClothByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to find cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedCloth, err := h.service.UpdateClothByID(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to updated cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to updated cloth", http.StatusOK, "success", cloth.UpdatedFormatCloth(updatedCloth, oldCloth))
	c.JSON(http.StatusOK, response)

}

func (h *clothHandler) DeleteClothByID(c *gin.Context) {
	var inputID cloth.ClothInputDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get supplier", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedCloth, err := h.service.DeleteClothByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to deleted cloth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to deleted cloth", http.StatusOK, "success", deletedCloth)
	c.JSON(http.StatusOK, response)

}
