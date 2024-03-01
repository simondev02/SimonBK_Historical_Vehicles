package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/domain/models"
	"errors"

	"gorm.io/gorm"
)

// GetAllHistoricalTourByID obtiene un registro historico de ruta por ID
func GetHistoricalToursById(db *gorm.DB, id uint) (*views.ToursById, error) {
	var record models.AvlRecord
	if err := db.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	responseRecordPoint := &views.ToursById{
		ID:             record.ID,
		TimeStampEvent: record.TimeStampEvent,
		Location:       record.Location,
		Event:          record.Event,
	}

	return responseRecordPoint, nil
}
