package model


type UserRole struct {
	UserId uint64
	RoleId uint64
}

func (UserRole) TableName() string {
	return "st_user_role"
}

