package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(ID int) (User, error)
	SaveOTPRequest(otpRequest OtpRequest) (OtpRequest, error)
	FindOTPByOTP(OTP string) (OtpRequest, error)
	UpdateIsVerifiedOTP(otpRequest OtpRequest) (OtpRequest, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindById(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) SaveOTPRequest(otpRequest OtpRequest) (OtpRequest, error) {
	err := r.db.Create(&otpRequest).Error
	if err != nil {
		return otpRequest, err
	}

	return otpRequest, nil
}

func (r *repository) FindOTPByOTP(OTP string) (OtpRequest, error) {
	var otpRequest OtpRequest

	err := r.db.Where("otp = ?", OTP).First(&otpRequest).Error
	if err != nil {
		return otpRequest, err
	}

	return otpRequest, nil
}

func (r *repository) UpdateIsVerifiedOTP(otpRequest OtpRequest) (OtpRequest, error) {
	err := r.db.Save(&otpRequest).Error
	if err != nil {
		return otpRequest, err
	}

	return otpRequest, nil
}
