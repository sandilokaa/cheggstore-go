package cloth

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	SaveCloth(cloth Cloth) (Cloth, error)
	SaveClothVariation(clothVariation ClothVariation) (ClothVariation, error)
	FindAllCloth(search string) ([]Cloth, error)
	FindClothByID(ID int) (Cloth, error)
	FindClothVariationByID(ID int) (ClothVariation, error)
	UpdateClothByID(cloth Cloth) (Cloth, error)
	UpdateClothVariationByID(clothVariation ClothVariation) (ClothVariation, error)
	// UpdateStockByClothID(clothID int, newStock int) error
	DeleteClothById(ID int) (Cloth, error)
	DeleteClothVariationByClothId(clothID int) (ClothVariation, error)
	CreateClothImage(clothImage ClothImage) (ClothImage, error)
	MarkAllImagesAsNonPrimary(clothID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveCloth(cloth Cloth) (Cloth, error) {
	err := r.db.Create(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (r *repository) SaveClothVariation(clothVariation ClothVariation) (ClothVariation, error) {
	err := r.db.Create(&clothVariation).Error
	if err != nil {
		return clothVariation, err
	}

	return clothVariation, nil
}

func (r *repository) FindAllCloth(search string) ([]Cloth, error) {
	var cloths []Cloth

	query := r.db
	if search != "" {
		query = query.Preload("ClothImages", "ClothImages.is_primary = 1").Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Preload("ClothImages", "ClothImages.is_primary = 1").Find(&cloths).Error
	if err != nil {
		return cloths, err
	}

	return cloths, nil
}

func (r *repository) FindClothByID(ID int) (Cloth, error) {
	var cloth Cloth

	err := r.db.Preload(clause.Associations).Preload("ClothImages").Where("id = ?", ID).Find(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (r *repository) FindClothVariationByID(ID int) (ClothVariation, error) {
	var clothVariation ClothVariation

	err := r.db.Where("id = ?", ID).Find(&clothVariation).Error
	if err != nil {
		return clothVariation, err
	}

	return clothVariation, nil
}

func (r *repository) UpdateClothByID(cloth Cloth) (Cloth, error) {
	err := r.db.Save(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (r *repository) UpdateClothVariationByID(clothVariation ClothVariation) (ClothVariation, error) {
	err := r.db.Save(&clothVariation).Error
	if err != nil {
		return clothVariation, err
	}

	return clothVariation, nil
}

// func (r *repository) UpdateStockByClothID(clothID int, newStock int) error {
// 	var cloth ClothVariation
// 	if err := r.db.First(&cloth, clothID).Error; err != nil {
// 		return err
// 	}

// 	cloth.Stock = newStock
// 	return r.db.Save(&cloth).Error
// }

func (r *repository) DeleteClothById(ID int) (Cloth, error) {
	var cloth Cloth
	err := r.db.Where("id = ?", ID).Delete(&cloth).Error
	if err != nil {
		return cloth, err
	}

	return cloth, nil
}

func (r *repository) DeleteClothVariationByClothId(clothID int) (ClothVariation, error) {
	var clothVariation ClothVariation
	err := r.db.Where("cloth_id = ?", clothID).Delete(&clothVariation).Error
	if err != nil {
		return clothVariation, err
	}

	return clothVariation, nil
}

func (r *repository) CreateClothImage(clothImage ClothImage) (ClothImage, error) {
	err := r.db.Create(&clothImage).Error
	if err != nil {
		return clothImage, err
	}

	return clothImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(clothID int) (bool, error) {
	err := r.db.Model(&ClothImage{}).Where("cloth_id = ?", clothID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
