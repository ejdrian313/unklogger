package services

import (
	"unklogger.com/src/domain/entities"
	"unklogger.com/src/domain/repositories"
)

type ActivityService struct {
	repo repositories.ActivityRepository
}

func NewActivityService(repo repositories.ActivityRepository) *ActivityService {
	return &ActivityService{repo: repo}
}

func (s *ActivityService) CreateActivity(activity *entities.UserActivityLog) error {
	return s.repo.CreateActivity(activity)
}

func (s *ActivityService) GetActivities() ([]*entities.UserActivityLog, error) {
	return s.repo.GetActivities()
}
