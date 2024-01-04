package response

type Responses interface {
	SetCode(int)
	SetMsg(string)
	SetData(interface{})
	Clone() *response // 初始化/重置
}
