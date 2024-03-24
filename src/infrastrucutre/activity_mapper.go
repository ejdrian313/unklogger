package infrastrucutre

import (
	"unklogger.com/src/domain/entities"
)

func ToDBActivity(product *entities.UserActivityLog) *ActivityLog {
	var p = &ActivityLog{
		UserName: product.Name,
		LastSeen: product.LastSeen,
		UserID:   uint(product.UserID),
	}

	return p
}

func FromDBProduct(dbProduct *ActivityLog) *entities.UserActivityLog {
	return &entities.UserActivityLog{
		Name:     dbProduct.UserName,
		LastSeen: dbProduct.LastSeen,
		UserID:   dbProduct.UserID,
	}
}
