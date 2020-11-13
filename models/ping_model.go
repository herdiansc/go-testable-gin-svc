package models

type DivideRequest struct {
	A float64 `json:"a" validate:"required"`
	B float64 `json:"b" validate:"required"`
}
