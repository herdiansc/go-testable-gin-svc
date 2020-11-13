package services

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/herdiansc/go-testable-gin-svc/models"
	"net/http"
)

type iPingService interface {
	Pong() (string, error)
	Divide(request []byte) models.Response
	Calculate(request []byte) models.Response
}

var (
	PingService iPingService = pingService{}
)

type pingService struct {
}

func (svc pingService) Pong() (string, error) {
	fmt.Println("doing complex things...")
	return "pong", nil
}

func (svc pingService) Divide(request []byte) models.Response {
	payload := models.DivideRequest{}

	err := json.Unmarshal(request, &payload)
	if err != nil {
		return models.Http.Response(http.StatusBadRequest, nil, err.Error())
	}
	validate := validator.New()
	err = validate.Struct(payload)
	if err != nil {
		return models.Http.Response(http.StatusBadRequest, nil, err.Error())
	}

	c := payload.A / payload.B
	content := fmt.Sprintf("%0.0f divided by %0.0f is %0.2f", payload.A, payload.B, c)
	return models.Http.Response(http.StatusOK, content, "success")
}

func requiredIf(fl validator.FieldLevel) bool {
	if fl.Field().Float() == 0 && fl.Top().FieldByName("Operator").String() == "/" {
		return false
	}
	return true
}

func (svc pingService) Calculate(request []byte) models.Response {
	payload := models.CalculateRequest{}

	err := json.Unmarshal(request, &payload)
	if err != nil {
		return models.Http.Response(http.StatusBadRequest, nil, err.Error())
	}
	validate := validator.New()
	_ = validate.RegisterValidation("required_if", requiredIf)

	err = validate.Struct(payload)
	if err != nil {
		return models.Http.Response(http.StatusBadRequest, nil, err.Error())
	}

	var c float64
	switch payload.Operator {
	case "+":
		c = payload.A + payload.B
	case "-":
		c = payload.A - payload.B
	case "*":
		c = payload.A * payload.B
	case "/":
		c = payload.A / payload.B
	}
	content := fmt.Sprintf("%0.0f %s %0.0f is %0.2f", payload.A, payload.Operator, payload.B, c)
	return models.Http.Response(http.StatusOK, content, "success")
}
