package model

import "github.com/bearllfleed/scholar-track/rpc/system/internal/global"

type Menu struct {
	global.StModel
	MenuLevel uint   `json:"-"`
	ParentId  uint64 `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path      string `json:"path" gorm:"comment:路由path"`        // 路由path
	Name      string `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden    bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component string `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort      int    `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	Meta      Meta   `json:"meta" gorm:"embedded;comment:附加属性"` // 附加属性
	Roles     []Role `json:"roles" gorm:"many2many:st_role_menu;"`
	Children  []Menu `json:"children" gorm:"-"`
}
type Meta struct {
	ActiveName string `json:"activeName" gorm:"comment:高亮菜单"`
	KeepAlive  bool   `json:"keepAlive" gorm:"comment:是否缓存"`   // 是否缓存
	Title      string `json:"title" gorm:"comment:菜单名"`        // 菜单名
	Icon       string `json:"icon" gorm:"comment:菜单图标"`        // 菜单图标
	CloseTab   bool   `json:"closeTab" gorm:"comment:自动关闭tab"` // 自动关闭tab
}

func (Menu) TableName() string {
	return "st_menu"
}
