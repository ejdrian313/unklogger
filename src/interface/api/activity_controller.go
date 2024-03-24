package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"unklogger.com/src/application/services"
	"unklogger.com/src/domain/entities"
)

type ActivityController struct {
	service *services.ActivityService
}

func NewActivityController(e *echo.Echo, service *services.ActivityService, m ...echo.MiddlewareFunc) {
	controller := &ActivityController{
		service: service,
	}

	e.POST("/activity", controller.CreateActivity, m...)
	e.GET("/activity", controller.GetAllActivity, m...)
}

func (ac *ActivityController) CreateActivity(c echo.Context) error {
	activity := &entities.UserActivityLog{
		UserID:   c.Get("ID").(uint),
		Name:     c.Get("username").(string),
		LastSeen: time.Now(),
	}

	err := ac.service.CreateActivity(activity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create activity",
		})
	}

	return c.JSON(http.StatusCreated, activity)
}

func (ac *ActivityController) GetAllActivity(c echo.Context) error {
	activity, err := ac.service.GetActivities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch activity",
		})
	}

	return c.JSON(http.StatusOK, activity)
}
