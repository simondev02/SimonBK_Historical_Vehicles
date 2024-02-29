package tours

import (
	"SimonBK_Historical_Vehicles/domain/models"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

func GetLastRecordDateByPlateOrImei(db *gorm.DB, plate, imei *string) (*time.Time, error) {
	db.Debug()
	var record models.AvlRecord
	var err error

	if plate != nil {
		err = db.Where("plate ilike ?", plate).Order("time_stamp_event desc").First(&record).Error
	} else if imei != nil {
		err = db.Where("imei ilike ?", imei).Order("time_stamp_event desc").First(&record).Error
	} else {
		return nil, fmt.Errorf("placa y imei no proporcionados")
	}

	if err != nil {
		log.Println("[GetLastRecorDateByPlateOrImei]", err)
		return nil, fmt.Errorf("placa o imei no encontrados")
	}

	return record.TimeStampEvent, nil
}
