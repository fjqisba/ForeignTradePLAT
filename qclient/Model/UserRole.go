package Model

type LoginRet int

const(
	//无效登录
	LOGIN_INVALID LoginRet = 0
	//管理员登录
	LOGIN_ADMIN LoginRet = 1
	//普通用户
	LOGIN_USER LoginRet = 2
)