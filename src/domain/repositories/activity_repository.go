package repositories

import "unklogger.com/src/domain/entities"

type ActivityRepository interface {
	GetActivities() ([]*entities.UserActivityLog, error)
	CreateActivity(activity *entities.UserActivityLog) error
}
