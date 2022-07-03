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