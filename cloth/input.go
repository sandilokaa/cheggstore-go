package cloth

import (
	"cheggstore/user"
)

type CreateClothInput struct {
	User        user.User
	MaterialID  int    `json:"material_id" binding:"required"`
	SupplierID  int    `json:"supplier_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Color       string `json:"color" binding:"required"`
	Price       string `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Size        string `json:"size" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
}

type ClothInputDetail struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateClothInput struct {
	User        user.User
	MaterialID  int    `json:"material_id"`
	SupplierID  int    `json:"supplier_id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Size        string `json:"size"`
	Stock       int    `json:"stock"`
}

type CreateClothImageInput struct {
	ClothID   int  `form:"cloth_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
	User      user.User
}
