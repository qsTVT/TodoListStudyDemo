package serializer

// 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total int         `json:"total"`
}

// 带总数的返回
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: int(total),
		},
		Msg: "success",
	}
}
