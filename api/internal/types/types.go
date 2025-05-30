// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type AddCategoryReq struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Status      int32  `json:"status,optional"`
}

type AddCategoryResp struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Status      int32  `json:"status"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	DeletedAt   int64  `json:"-"`
}

type AddPolicyReq struct {
	Rules [][]string `json:"rules"`
}

type AddPropertyReq struct {
	Name         string `json:"name"`
	Key          string `json:"key"`
	Description  string `json:"description"`
	Status       int32  `json:"status"`
	Type         string `json:"type"`
	IsSearch     bool   `json:"isSearch"`
	IsRequired   bool   `json:"isRequired"`
	ValidateRule string `json:"validateRule"`
	CategoryId   uint64 `json:"categoryId"`
}

type AddPropertyResp struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	Description  string `json:"description"`
	Status       int32  `json:"status"`
	Type         string `json:"type"`
	IsSearch     bool   `json:"isSearch"`
	IsRequired   bool   `json:"isRequired"`
	ValidateRule string `json:"validateRule"`
	CategoryId   uint64 `json:"categoryId"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	DeletedAt    int64  `json:"-"`
}

type AddRoleReq struct {
	RoleName string `json:"roleName"`
	ParentId uint64 `json:"parentId"`
}

type AddRoleResp struct {
	RoleName  string `json:"roleName"`
	ParentId  uint64 `json:"parentId"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Id        uint64 `json:"id"`
}

type Api struct {
	Id          uint64 `json:"id"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	ApiGroup    string `json:"apiGroup"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	DeletedAt   int64  `json:"-"`
}

type Category struct {
	Id          uint64     `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
	Status      int32      `json:"status"`
	Properties  []Property `json:"properties"`
	CreatedAt   int64      `json:"createdAt"`
	UpdatedAt   int64      `json:"updatedAt"`
	DeletedAt   int64      `json:"-"`
}

type ChangePasswordReq struct {
	UserId      uint64 `json:"userId,optional"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ChangePasswordResp struct {
	Id uint64 `json:"id"`
}

type CreateApiReq struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	ApiGroup    string `json:"apiGroup"`
	Description string `json:"description"`
}

type CreateApiResp struct {
	Id          uint64 `json:"id"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	ApiGroup    string `json:"apiGroup"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	DeletedAt   int64  `json:"-"`
}

type CreateDictionaryDetailReq struct {
	DictionaryId   uint64 `json:"dictionaryId"`
	Key            string `json:"key"`
	Value          string `json:"value"`
	Sort           int32  `json:"sort"`
	IsDefaultValue bool   `json:"isDefaultValue"`
}

type CreateDictionaryDetailResp struct {
	Id             uint64 `json:"id"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
	DeletedAt      int64  `json:"-"`
	DictionaryId   uint64 `json:"dictionaryId"`
	Key            string `json:"key"`
	Value          string `json:"value"`
	Sort           int32  `json:"sort"`
	IsDefaultValue bool   `json:"isDefaultValue"`
}

type CreateDictionaryReq struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status bool   `json:"status"`
	Desc   string `json:"desc"`
}

type CreateDictionaryResp struct {
	Id        uint64 `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"-"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Status    bool   `json:"status"`
	Desc      string `json:"desc"`
}

type CreateMenuReq struct {
	ParentId  uint64 `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Meta      Meta   `json:"meta"`
}

type CreateMenuResp struct {
	Id uint64 `json:"id"`
}

type DeleteAchievementReq struct {
	Id uint64 `path:"id"`
}

type DeleteAchievementResp struct {
}

type DeleteApiReq struct {
	Id uint64 `path:"id"`
}

type DeleteApiResp struct {
}

type DeleteCategoryReq struct {
	Id uint64 `path:"id"`
}

type DeleteCategoryResp struct {
}

type DeleteDictionaryDetailReq struct {
	Id uint64 `path:"id"`
}

type DeleteDictionaryDetailResp struct {
	Id uint64 `json:"id"`
}

type DeleteDictionaryReq struct {
	Id uint64 `path:"id"`
}

type DeleteDictionaryResp struct {
	Id uint64 `json:"id"`
}

type DeleteFileReq struct {
	Id uint64 `path:"id"`
}

type DeleteFileResp struct {
}

type DeleteMenuReq struct {
	Id uint64 `path:"id"`
}

type DeleteMenuResp struct {
}

type DeletePropertyReq struct {
	Id uint64 `path:"id"`
}

type DeletePropertyResp struct {
}

type DeleteRoleReq struct {
	RoleId uint64 `path:"id"`
}

type DeleteRoleResp struct {
	Id uint64 `json:"id"`
}

type DeleteUserReq struct {
	UserId uint64 `path:"id"`
}

type DeleteUserResp struct {
	Id uint64 `json:"id"`
}

type Dictionary struct {
	Id                uint64              `json:"id"`
	CreatedAt         int64               `json:"createdAt"`
	UpdatedAt         int64               `json:"updatedAt"`
	DeletedAt         int64               `json:"-"`
	Name              string              `json:"name"`
	Type              string              `json:"type"`
	Status            bool                `json:"status"`
	Desc              string              `json:"desc"`
	DictionaryDetails []*DictionaryDetail `json:"dictionaryDetails"`
}

type DictionaryDetail struct {
	Id             uint64 `json:"id"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
	DeletedAt      int64  `json:"-"`
	DictionaryId   uint64 `json:"dictionaryId"`
	Key            string `json:"key"`
	Value          string `json:"value"`
	Sort           int32  `json:"sort"`
	IsDefaultValue bool   `json:"isDefaultValue"`
}

type DownloadFileReq struct {
	Id uint64 `path:"id"`
}

type DownloadFileResp struct {
	FileData []byte `json:"fileData"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	FileSize int64  `json:"fileSize"`
}

type File struct {
	Id         uint64 `json:"id"`
	FileKey    string `json:"fileKey"`
	FileUrl    string `json:"fileUrl"`
	FileName   string `json:"fileName"`
	FileSize   int64  `json:"fileSize"`
	FileType   string `json:"fileType"`
	BusinessId uint64 `json:"bussinessId"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
	DeletedAt  int64  `json:"-"`
}

type GetBussinessFilesReq struct {
	BusinessId    uint64 `json:"bussinessId"`
	BussinessName string `json:"bussinessName"`
}

type GetBussinessFilesResp struct {
	Files []*File `json:"files"`
}

type GetFileUrlReq struct {
	FileKey string `json:"fileKey"`
}

type GetFileUrlResp struct {
	Url string `json:"url"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type Menu struct {
	Id        uint64 `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"-"`
	MenuLevel uint   `json:"-"`
	ParentId  uint64 `json:"parentId"`  // 父菜单ID
	Path      string `json:"path"`      // 路由path
	Name      string `json:"name"`      // 路由name
	Hidden    bool   `json:"hidden"`    // 是否在列表隐藏
	Component string `json:"component"` // 对应前端文件路径
	Sort      int    `json:"sort"`      // 排序标记
	Meta      Meta   `json:"meta"`      // 附加属性
	Children  []Menu `json:"children"`
}

type Meta struct {
	ActiveName string `json:"activeName"`
	KeepAlive  bool   `json:"keepAlive"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	CloseTab   bool   `json:"closeTab"`
}

type PolicyInfo struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type Property struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	Description  string `json:"description"`
	Status       int32  `json:"status,optional"`
	Type         string `json:"type"`
	IsSearch     bool   `json:"isSearch"`
	IsRequired   bool   `json:"isRequired"`
	ValidateRule string `json:"validateRule"`
	CategoryId   uint64 `json:"categoryId"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	DeletedAt    int64  `json:"-"`
}

type QueryAllApiReq struct {
}

type QueryAllApiResp struct {
	Apis []Api `json:"apis"`
}

type QueryAllMenuTreeReq struct {
}

type QueryAllMenuTreeResp struct {
	Menus []Menu `json:"menus"`
}

type QueryApiDetailReq struct {
	Id uint64 `path:"id"`
}

type QueryApiDetailResp struct {
	Id          uint64 `json:"id"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	ApiGroup    string `json:"apiGroup"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	DeletedAt   int64  `json:"-"`
}

type QueryCategoryDetailReq struct {
	Id uint64 `path:"id"`
}

type QueryCategoryDetailResp struct {
	Id          uint64     `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
	Status      int32      `json:"status"`
	Properties  []Property `json:"properties"`
	CreatedAt   int64      `json:"createdAt"`
	UpdatedAt   int64      `json:"updatedAt"`
	DeletedAt   int64      `json:"-"`
}

type QueryCategoryReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Name     string `json:"name,optional"`
	Type     string `json:"type,optional"`
	Status   int32  `json:"status,optional"`
	OrderBy  string `json:"orderBy,optional"`
}

type QueryCategoryResp struct {
	Total      int64       `json:"total"`
	Categories []*Category `json:"categories"`
	Page       int64       `json:"page"`
	PageSize   int64       `json:"pageSize"`
}

type QueryDictionaryDetailReq struct {
	Id uint64 `path:"id"`
}

type QueryDictionaryDetailResp struct {
	Id        uint64 `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"-"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Status    bool   `json:"status"`
	Desc      string `json:"desc"`
}

type QueryDictionaryListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Name     string `json:"name,optional"`
	Type     string `json:"type,optional"`
	OrderBy  string `json:"orderBy,optional"`
}

type QueryDictionaryListResp struct {
	Total    int64         `json:"total"`
	List     []*Dictionary `json:"list"`
	Page     int64         `json:"page"`
	PageSize int64         `json:"pageSize"`
}

type QueryFileDetailReq struct {
	Id uint64 `path:"id"`
}

type QueryFileDetailResp struct {
	File File `json:"file"`
}

type QueryMenuDetailReq struct {
	Id uint64 `path:"id"`
}

type QueryMenuDetailResp struct {
	Id        uint64 `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	ParentId  uint64 `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Meta      Meta   `json:"meta"`
	Children  []Menu `json:"children"`
}

type QueryPropertyDetailReq struct {
	Id uint64 `path:"id"`
}

type QueryPropertyDetailResp struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	Description  string `json:"description"`
	Status       int32  `json:"status"`
	Type         string `json:"type"`
	IsSearch     bool   `json:"isSearch"`
	IsRequired   bool   `json:"isRequired"`
	ValidateRule string `json:"validateRule"`
	CategoryId   uint64 `json:"categoryId"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	DeletedAt    int64  `json:"-"`
}

type QueryPropertyReq struct {
	Page       int64  `json:"page"`
	PageSize   int64  `json:"pageSize"`
	CategoryId uint64 `json:"categoryId"`
	Name       string `json:"name,optional"`
	Type       string `json:"type,optional"`
	Status     int32  `json:"status,optional"`
	OrderBy    string `json:"orderBy,optional"`
}

type QueryPropertyResp struct {
	Total      int64      `json:"total"`
	Properties []Property `json:"properties"`
	Page       int64      `json:"page"`
	PageSize   int64      `json:"pageSize"`
}

type QueryRoleMenuTreeReq struct {
	RoleId uint64 `json:"roleId"`
}

type QueryRoleMenuTreeResp struct {
	Menus []Menu `json:"menus"`
}

type QueryRolePoliciesReq struct {
	RoleId uint64 `path:"roleId"`
}

type QueryRolePoliciesResp struct {
	Rules []*PolicyInfo `json:"rules"`
}

type QuerySelfInfoReq struct {
	Id uint64 `json:"id,optional"`
}

type QuerySelfInfoResp struct {
	User QueryUserDetailResp `json:"user"`
}

type QueryUserDetailReq struct {
	Id uint64 `path:"id"`
}

type QueryUserDetailResp struct {
	Id        uint64 `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"-"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	RoleId    uint64 `json:"roleId"`
	Status    int32  `json:"status"`
	Nickname  string `json:"nickname"`
	Gender    int32  `json:"gender"`
	Major     string `json:"major"`
	Grade     string `json:"grade"`
	College   string `json:"college"`
	Realname  string `json:"realname"`
	Class     string `json:"class"`
	Password  string `json:"-"`
}

type QueryUserListReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	Username string `json:"username,optional"`
	Email    string `json:"email,optional"`
	Phone    string `json:"phone,optional"`
	Nickname string `json:"nickname,optional"`
	RoleId   uint64 `json:"roleId,optional"`
	Status   int32  `json:"status,optional"`
	Major    string `json:"major,optional"`
	Grade    string `json:"grade,optional"`
	College  string `json:"college,optional"`
	Realname string `json:"realname,optional"`
	Class    string `json:"class,optional"`
	Gender   int32  `json:"gender,optional"`
	OrderBy  string `json:"orderBy,optional"`
}

type QueryUserListResp struct {
	Total    int64                  `json:"total"`
	List     []*QueryUserDetailResp `json:"list"`
	Page     int64                  `json:"page"`
	PageSize int64                  `json:"pageSize"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,optional"`
	Phone    string `json:"phone,optional"`
	Nickname string `json:"nickname,optional"`
	RoleId   uint64 `json:"roleId"`
	Gender   int32  `json:"gender"`
	Major    string `json:"major"`
	Grade    string `json:"grade"`
	College  string `json:"college"`
	Realname string `json:"realname"`
	Class    string `json:"class"`
}

type RegisterResp struct {
	User *QueryUserDetailResp `json:"user"`
}

type ResetPasswordReq struct {
	UserId uint64 `json:"userId"`
}

type ResetPasswordResp struct {
	Id uint64 `json:"id"`
}

type RoleTree struct {
	RoleName  string      `json:"roleName"`
	ParentId  uint64      `json:"parentId"`
	CreatedAt int64       `json:"createdAt"`
	UpdatedAt int64       `json:"updatedAt"`
	Id        uint64      `json:"id"`
	Children  []*RoleTree `json:"children"`
}

type RoleTreeReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type RoleTreeResp struct {
	Roles    []*RoleTree `json:"roles"`
	Total    int64       `json:"total"`
	Page     int64       `json:"page"`
	PageSize int64       `json:"pageSize"`
}

type SetRoleInfoReq struct {
	RoleId   uint64 `json:"roleId"`
	RoleName string `json:"roleName"`
}

type SetRoleInfoResp struct {
	Id uint64 `json:"id"`
}

type SetRolePoliciesReq struct {
	RoleId uint64       `json:"roleId"`
	Rules  []PolicyInfo `json:"rules"`
}

type SetRolePoliciesResp struct {
	Id uint64 `json:"id"`
}

type SetSelfInfoReq struct {
	Id       uint64 `json:"id,optional"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	RoleId   uint64 `json:"roleId"`
	Status   int32  `json:"status"`
	Nickname string `json:"nickname"`
	Gender   int32  `json:"gender"`
	Major    string `json:"major"`
	Grade    string `json:"grade"`
	College  string `json:"college"`
	Realname string `json:"realname"`
	Class    string `json:"class"`
}

type SetSelfInfoResp struct {
	Id uint64 `json:"id"`
}

type SetUserInfoReq struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	RoleId   uint64 `json:"roleId"`
	Status   int32  `json:"status"`
	Nickname string `json:"nickname"`
	Gender   int32  `json:"gender"`
	Major    string `json:"major"`
	Grade    string `json:"grade"`
	College  string `json:"college"`
	Realname string `json:"realname"`
	Class    string `json:"class"`
}

type SetUserInfoResp struct {
	Id uint64 `json:"id"`
}

type SetUserRoleReq struct {
	UserId  uint64   `json:"userId"`
	RoleIds []uint64 `json:"roleIds"`
}

type SetUserRoleResp struct {
	Id uint64 `json:"id"`
}

type UpdateApiReq struct {
	Id          uint64 `json:"id"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	ApiGroup    string `json:"apiGroup"`
	Description string `json:"description"`
}

type UpdateApiResp struct {
	Id          uint64 `json:"id"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	ApiGroup    string `json:"apiGroup"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	DeletedAt   int64  `json:"-"`
}

type UpdateCategoryReq struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Status      int32  `json:"status"`
}

type UpdateCategoryResp struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Status      int32  `json:"status"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	DeletedAt   int64  `json:"-"`
}

type UpdateDictionaryDetailReq struct {
	Id             uint64 `json:"id"`
	DictionaryId   uint64 `json:"dictionaryId"`
	Key            string `json:"key"`
	Value          string `json:"value"`
	Sort           int32  `json:"sort"`
	IsDefaultValue bool   `json:"isDefaultValue"`
}

type UpdateDictionaryDetailResp struct {
	Id uint64 `json:"id"`
}

type UpdateDictionaryReq struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status bool   `json:"status"`
	Desc   string `json:"desc"`
}

type UpdateDictionaryResp struct {
	Id uint64 `json:"id"`
}

type UpdateMenuReq struct {
	Id        uint64 `json:"id"`
	ParentId  uint64 `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Meta      Meta   `json:"meta"`
}

type UpdateMenuResp struct {
	Id        uint64 `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt int64  `json:"-"`
	ParentId  uint64 `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Meta      Meta   `json:"meta"`
}

type UpdatePropertyReq struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	Description  string `json:"description"`
	Status       int32  `json:"status"`
	Type         string `json:"type"`
	IsSearch     bool   `json:"isSearch"`
	IsRequired   bool   `json:"isRequired"`
	ValidateRule string `json:"validateRule"`
	CategoryId   uint64 `json:"categoryId"`
}

type UpdatePropertyResp struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	Description  string `json:"description"`
	Status       int32  `json:"status"`
	Type         string `json:"type"`
	IsSearch     bool   `json:"isSearch"`
	IsRequired   bool   `json:"isRequired"`
	ValidateRule string `json:"validateRule"`
	CategoryId   uint64 `json:"categoryId"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	DeletedAt    int64  `json:"-"`
}

type UpdateRoleMenuReq struct {
	RoleId  uint64   `json:"roleId"`
	MenuIds []uint64 `json:"menuIds"`
}

type UpdateRoleMenuResp struct {
}

type UploadAchievementReq struct {
	Code        string                 `form:"code"`
	Name        string                 `form:"name"`
	Status      int32                  `form:"status,optional"`
	Level       int32                  `form:"level"`
	Rank        int32                  `form:"rank"`
	AwardTime   int64                  `form:"awardTime"`
	Share       bool                   `form:"share"`
	Star        uint64                 `form:"star,optional"`
	Uploader    uint64                 `form:"uploader"`
	Team        bool                   `form:"team"`
	TeamMembers []uint64               `form:"teamMembers"`
	Description string                 `form:"description"`
	CategoryId  uint64                 `form:"categoryId"`
	OtherInfo   map[string]interface{} `form:"otherInfo"`
}

type UploadAchievementResp struct {
	Id uint64 `json:"id"`
}

type UploadFileReq struct {
	File          []byte `json:"file"`
	FileName      string `json:"fileName"`
	FileType      string `json:"fileType"`
	BusinessId    uint64 `json:"bussinessId"`
	BussinessName string `json:"bussinessName"`
}

type UploadFileResp struct {
	FileId uint64 `json:"fileId"`
}
