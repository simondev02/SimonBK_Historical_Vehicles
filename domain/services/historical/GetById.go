package services

import (
	"SimonBK_Historical_Vehicles/api/views"
	"SimonBK_Historical_Vehicles/domain/models"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

// GetAvlRecordByID obtiene un registro Avl por ID
func GetHistoricalByID(db *gorm.DB, id uint) (*views.HistoricalById, error) {
	var record models.AvlRecord
	if err := db.First(&record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	// Deserializar las propiedades JSON
	var properties map[string]interface{}
	if err := json.Unmarshal([]byte(*record.Properties), &properties); err != nil {
		return nil, err
	}

	// Convertir el registro al formato de respuesta
	responseRecord := &views.HistoricalById{
		ID:         record.ID,
		Properties: properties,
	}

	return responseRecord, nil
}
