package infrastrucutre

import (
	"gorm.io/gorm"
	"unklogger.com/src/domain/entities"
	"unklogger.com/src/domain/repositories"
)

type GormActivityRepository struct {
	db *gorm.DB
}

func NewGormActivityRepository(db *gorm.DB) repositories.ActivityRepository {
	return &GormActivityRepository{db: db}
}

func (repo *GormActivityRepository) GetActivities() ([]*entities.UserActivityLog, error) {
	var dbActivities []ActivityLog
	if err := repo.db.Find(&dbActivities).Error; err != nil {
		return nil, err
	}
	activities := make([]*entities.UserActivityLog, len(dbActivities))
	for i, dbProduct := range dbActivities {
		activities[i] = FromDBProduct(&dbProduct)
	}
	return activities, nil
}

func (repo *GormActivityRepository) CreateActivity(activity *entities.UserActivityLog) error {
	dbActivity := ToDBActivity(activity)
	return repo.db.Save(dbActivity).Error
}
