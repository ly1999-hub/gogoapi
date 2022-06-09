package apimodel

import (
	mongodb2 "myapp/module/mongodb"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myapp/internal/model"
	"myapp/internal/response"
	"myapp/internal/util/pstring"
)

// StaffCreate ...
type StaffCreate struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

// Validate ...
func (c StaffCreate) Validate() error {
	c.Name = pstring.TrimSpace(c.Name)
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error(response.CommonInvalidName)),
		validation.Field(&c.Phone, validation.Required.Error(response.CommonInvalidPhone)),
		validation.Field(&c.Email, validation.Required.Error(response.CommonInvalidEmail)),
		validation.Field(&c.Role, validation.Required.Error(response.CommonInvalidRole)),
	)
}

// NewStaffRaw ...
func (c StaffCreate) NewStaffRaw(roleID primitive.ObjectID, password string) model.Staff {
	now := time.Now()
	return model.Staff{
		ID:             mongodb2.NewObjectID(),
		Name:           c.Name,
		Email:          c.Email,
		SearchString:   mongodb2.NonAccentVietnamese(c.Name),
		Phone:          c.Phone,
		Active:         false,
		Role:           &roleID,
		Avatar:         c.Avatar,
		HashedPassword: password,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

// StaffUpdate ...
type StaffUpdate struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Avatar string `json:"avatar"`
}

// Validate ...
func (c StaffUpdate) Validate() error {
	c.Name = pstring.TrimSpace(c.Name)
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error(response.CommonInvalidName)),
		validation.Field(&c.Phone, validation.Required.Error(response.CommonInvalidPhone)),
		validation.Field(&c.Email, validation.Required.Error(response.CommonInvalidEmail)),
		validation.Field(&c.Role, validation.Required.Error(response.CommonInvalidRole)),
	)
}

// StaffAll ...
type StaffAll struct {
	Page    int64  `query:"page"`
	Limit   int64  `query:"limit"`
	Active  string `query:"active"`
	Keyword string `query:"keyword"`
}

// Validate ...
func (o StaffAll) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Page, validation.Required.Error(response.CommonInvalidPagination)),
		validation.Field(&o.Limit, validation.Required.Error(response.CommonInvalidPagination)),
	)
}

// StaffChangePassword ...
type StaffChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// Validate ...
func (c StaffChangePassword) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.OldPassword, validation.Required.Error(response.CommonInvalidOldPassword)),
		validation.Field(&c.NewPassword, validation.Required.Error(response.CommonInvalidNewPassword)),
	)
}

// StaffForgotPassword ...
type StaffForgotPassword struct {
	Email string `json:"email"`
}

// Validate ...
func (c StaffForgotPassword) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required.Error(response.CommonInvalidEmail)),
	)
}

// StaffResetPassword ...
type StaffResetPassword struct {
	Password string `json:"password"`
}

// Validate ...
func (c StaffResetPassword) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Password, validation.Required.Error(response.CommonInvalidNewPassword)),
	)
}

// StaffLoginWithEmail ...
type StaffLoginWithEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate ...
func (c StaffLoginWithEmail) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required.Error(response.CommonInvalidEmail)),
		validation.Field(&c.Password, validation.Required.Error(response.CommonInvalidPassword)),
	)
}
