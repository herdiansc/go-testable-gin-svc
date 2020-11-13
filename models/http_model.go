package models

type Response struct {
	Code int `json:"code"`
	Content interface{} `json:"content"`
	Message string `json:"message"`
}

var (
	Http httpModel
)

type httpModel struct {}

func (model httpModel) Response(c int, ct interface{}, msg string) Response {
	return Response{
		Code: c,
		Content: ct,
		Message: msg,
	}
}