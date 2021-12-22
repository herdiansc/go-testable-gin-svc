package models

//type DivideRequest struct {
//	A float64 `json:"a" validate:"required"`
//	B float64 `json:"b" validate:"required"`
//}
//
//type DivideResponse struct {
//	Message string `json:"message"`
//	Result  float64 `json:"result"`
//}
//
//type CalculateRequest struct {
//	A float64 `json:"a" validate:"required"`
//	Operator string `json:"operator" validate:"required,oneof=+ - * /"`
//	B float64 `json:"b" validate:"required_if=Operator /"`
//}

type HealthCheckResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func NewHealthCheckResponse(l string, v string) HealthCheckResponse {
	return HealthCheckResponse{
		Label: l,
		Value: v,
	}
}