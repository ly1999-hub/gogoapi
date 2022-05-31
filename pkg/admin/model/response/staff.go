package responsemodel

import (
	"myapp/internal/util/ptime"
)

// Staff ...
type Staff struct {
	ID        string              `json:"_id"`
	Name      string              `json:"name"`
	Phone     string              `json:"phone"`
	Active    bool                `json:"active"`
	Email     string              `json:"email"`
	Role      *RoleShort          `json:"role,omitempty"`
	Avatar    string              `json:"avatar"`
	IsRoot    bool                `json:"isRoot"`
	CreatedAt *ptime.TimeResponse `json:"createdAt" swaggertype:"string"`
	UpdatedAt *ptime.TimeResponse `json:"updatedAt" swaggertype:"string"`
}

// StaffLogin ...
type StaffLogin struct {
	ID    string `json:"_id"`
	Token string `json:"token"`
}
