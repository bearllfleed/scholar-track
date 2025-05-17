package model

type RoleMenu struct {
	RoleId uint64 `gorm:"column:role_id;type:bigint(20);not null;comment:角色ID"`
	MenuId uint64 `gorm:"column:menu_id;type:bigint(20);not null;comment:菜单ID"`
}

func (RoleMenu) TableName() string {
	return "st_role_menu"
}
