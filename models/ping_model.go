package models

type DivideRequest struct {
	A float64 `json:"a" validate:"required"`
	B float64 `json:"b" validate:"required"`
}

type CalculateRequest struct {
	A float64 `json:"a" validate:"required"`
	Operator string `json:"operator" validate:"required,oneof=+ - * /"`
	B float64 `json:"b" validate:"required_if=Operator /"`
}
