package Model

type CustomerDataRsp struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data []CustomerInformation `json:"data"`
	Count int	`json:"count"`
}

type AssignTaskReq struct {
	Target string	`json:"target"`
	Domain []string	`json:"domain"`
}

type UserLevel int

const(
	//无效登录
	USER_INVALID UserLevel = 0
	//管理员登录
	USER_ADMIN UserLevel = 1
	//普通用户
	USER_USER UserLevel = 2
)