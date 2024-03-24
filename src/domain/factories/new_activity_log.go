package factories

import (
	"errors"
	"time"

	"unklogger.com/src/domain/entities"
)

func NewActivityLog(userID uint, name string) (*entities.UserActivityLog, error) {
	if userID == 0 {
		return nil, errors.New("invalid seller details")
	}
	if name == "" {
		return nil, errors.New("invalid name")
	}
	return &entities.UserActivityLog{Name: name, UserID: userID, LastSeen: time.Now()}, nil
}
