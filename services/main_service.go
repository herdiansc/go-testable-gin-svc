package services

import (
	"fmt"
	"github.com/herdiansc/go-testable-gin-svc/models"
	"time"

	//"github.com/herdiansc/go-testable-gin-svc/main"
	"net/http"
)

type iMainService interface {
	Pong() models.Response
	HealthCheck() models.Response
	RegisterAppStartTime() models.Response
}

var ServiceStartTime time.Time

var (
	MainService iMainService = mainService{}
)

type mainService struct {
	StartTime time.Time
}

func (svc mainService) Pong() models.Response {
	fmt.Println("doing complex things...")
	return models.Http.Response(http.StatusOK, "pong", "success")
}

func (svc mainService) HealthCheck() models.Response {
	fmt.Println("getting service uptime...")
	duration := time.Since(ServiceStartTime).String()
	content := models.NewHealthCheckResponse(fmt.Sprintf("service uptime duration %+v", duration), duration)
	return models.Http.Response(http.StatusOK, content, "success")
}

func (svc mainService) RegisterAppStartTime() models.Response {
	now := time.Now()
	fmt.Printf("registering application start time... %+v\n", now)
	ServiceStartTime = now
	return models.Http.Response(http.StatusOK, fmt.Sprintf("start time is: %+v", now), "success")
}