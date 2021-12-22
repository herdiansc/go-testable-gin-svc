package models

type DivideRequest struct {
	A float64 `json:"a" validate:"required"`
	B float64 `json:"b" validate:"required"`
}

type DivideResponse struct {
	Label  string `json:"label"`
	Result float64 `json:"result"`
}

func NewDivideResponse(l string, r float64) DivideResponse {
	return DivideResponse{
		Label:  l,
		Result: r,
	}
}

type CalculateRequest struct {
	A float64 `json:"a" validate:"required"`
	Operator string `json:"operator" validate:"required,oneof=+ - * /"`
	B float64 `json:"b" validate:"required_if=Operator /"`
}
