package exception

// 编码规则：
// 100100010 长度9位
// 前4位： 1001 代表系统模块
// 后5位： 00010 代表功能点
const (
	// NOTE: 基础
	ApiNoPermissionFail = 403
	GetOrmFail          = 100000010

	// NOTE: 登陆
	LoginFail                        = 100200010
	LoginGetUserNameFail             = 100200020
	LoginInvalidVerificationCodeFail = 100200030
	LoginUsernameOrPasswordFail      = 100200040

	// NOTE: 用户
	InsertUserFail         = 100100010
	InsetUserAvatarFail    = 100100011
	UpdateUserFail         = 100100020
	UpdateUserPwdFail      = 100100021
	UpdateUserResetPwdFail = 100100023
	UpdateUserStatusFail   = 100100022
	DeleteUserFail         = 100100030
	GetPageUserFail        = 100100040
	GetUserFail            = 100100050
	GetUserProfileFail     = 100100051
	GetUserInfoFail        = 100100052

	// NOTE: 岗位
	InsertPostFail  = 100300010
	UpdatePostFail  = 100300020
	DeletePostFail  = 100300030
	GetPagePostFail = 100300040
	GetPostFail     = 100300050
	OptionPostFail  = 100300060

	// NOTE: 部门
	InsertDeptFail  = 100400010
	UpdateDeptFail  = 100400020
	DeleteDeptFail  = 100400030
	GetPageDeptFail = 100400040
	GetDeptFail     = 100400050

	// NOTE: 角色
	InsertRoleFail          = 100500010
	UpdateRoleFail          = 100500020
	UpdateRoleStatusFail    = 100500021
	UpdateRoleDataScopeFail = 100500022
	DeleteRoleFail          = 100500030
	GetPageRoleFail         = 100500040
	GetRoleFail             = 100500050
	GetRoleOptionFail       = 100500060

	// NOTE: 菜单
	InsertMenuFail  = 100600010
	UpdateMenuFail  = 100600020
	DeleteMenuFail  = 100600030
	GetPageMenuFail = 100600040
	GetMenuFail     = 100600050
)

var resultCodeText = map[int]string{
	ApiNoPermissionFail:              "对不起，您没有该接口访问权限，请联系管理员",
	LoginFail:                        "登录失败",
	LoginGetUserNameFail:             "登陆用户名获取失败",
	LoginInvalidVerificationCodeFail: "验证码错误",
	LoginUsernameOrPasswordFail:      "账号或密码错误",
	GetOrmFail:                       "获取Orm失败",
	InsertUserFail:                   "添加用户失败",
	InsetUserAvatarFail:              "添加用户头像失败",
	UpdateUserFail:                   "更新用户失败",
	UpdateUserPwdFail:                "更新用户密码失败",
	UpdateUserResetPwdFail:           "重置用户密码失败",
	UpdateUserStatusFail:             "更新用户状态失败",
	DeleteUserFail:                   "删除用户失败",
	GetPageUserFail:                  "分页查询用户失败",
	GetUserFail:                      "查询用户失败",
	GetUserProfileFail:               "获取个人中心信息失败",
	GetUserInfoFail:                  "获取个人信息失败",
	InsertPostFail:                   "添加岗位失败",
	UpdatePostFail:                   "更新岗位失败",
	DeletePostFail:                   "删除岗位失败",
	GetPagePostFail:                  "分页查询岗位失败",
	GetPostFail:                      "查询岗位失败",
	OptionPostFail:                   "获取职位下拉框数据失败",
	InsertDeptFail:                   "添加部门失败",
	UpdateDeptFail:                   "更新部门失败",
	DeleteDeptFail:                   "删除部门失败",
	GetPageDeptFail:                  "分页查询部门失败",
	GetDeptFail:                      "查询部门失败",
	InsertRoleFail:                   "添加角色失败",
	UpdateRoleFail:                   "更新角色失败",
	UpdateRoleStatusFail:             "更新角色状态失败",
	UpdateRoleDataScopeFail:          "更新角色权限失败",
	DeleteRoleFail:                   "删除角色失败",
	GetPageRoleFail:                  "分页查询角色失败",
	GetRoleFail:                      "查询角色失败",
	GetRoleOptionFail:                "获取角色下拉数据失败",
	InsertMenuFail:                   "添加菜单失败",
	UpdateMenuFail:                   "更新菜单失败",
	DeleteMenuFail:                   "删除菜单失败",
	GetPageMenuFail:                  "分页查询菜单失败",
	GetMenuFail:                      "查询菜单失败",
}

func StatusText(code int) (string, bool) {
	message, ok := resultCodeText[code]
	return message, ok
}
