package cerror

import (
	"google.golang.org/grpc/status"
)

var (
	ErrUserHasExists  = status.Error(601, "用户名已存在")
	ErrUserNotFound   = status.Error(602, "用户不存在")
	ErrEmailHasExists = status.Error(603, "邮箱已存在")
	ErrPhoneHasExists = status.Error(604, "手机号已存在")
	ErrPasswordWrong  = status.Error(605, "密码错误")
	ErrUserDisabled   = status.Error(606, "用户被禁用")
)

var (
	ErrNotLogin         = status.Error(701, "用户未登录")
	ErrPermissionDenied = status.Error(702, "权限不足")
	ErrTokenInvalid     = status.Error(703, "token无效")
	ErrTokenExpired     = status.Error(704, "token过期")
	ErrTokenNotMatch    = status.Error(705, "token不匹配")
)

var (
	ErrRoleHasExists       = status.Error(801, "角色名称已存在")
	ErrRoleNotExists       = status.Error(802, "角色不存在")
	ErrParentRoleNotExists = status.Error(803, "父角色不存在")
	ErrInvalidRole         = status.Error(804, "添加的角色中有无效数据")
)

var (
	ErrHasSameApi         = status.Error(901, "添加的权限中有相同的api")
	ErrCasbinUpdateFailed = status.Error(902, "casbin更新失败")
	ErrApiAlreadyExists   = status.Error(1001, "api已存在")
	ErrApiNotFound        = status.Error(1002, "api不存在")
)

var (
	ErrCategoryNameAlreadyExists = status.Error(1101, "分类名称已存在")
	ErrCategoryNotFound          = status.Error(1102, "分类不存在")
)

var (
	ErrMenuNotFound       = status.Error(1201, "菜单不存在")
	ErrMenuHasExists      = status.Error(1202, "菜单已存在")
	ErrMenuParentNotFound = status.Error(1203, "父菜单不存在")
	ErrMenuInvalid        = status.Error(1204, "添加的菜单中有无效数据")
)

var (
	ErrFileNotFound = status.Error(1301, "文件不存在")
)

var (
	ErrDictionaryHasExists       = status.Error(1401, "字典已存在")
	ErrDictionaryNotFound        = status.Error(1402, "字典不存在")
	ErrDictionaryDetailHasExists = status.Error(1403, "字典详情已存在")
	ErrDictionaryDetailNotFound  = status.Error(1404, "字典详情不存在")
)

var (
	ErrAchieveHasExists = status.Error(1501, "成果已存在，请检查成果编号")
)

var (
	ErrGenerateSnowflakeIdFailed = status.Error(1601, "生成雪花id失败")
)
