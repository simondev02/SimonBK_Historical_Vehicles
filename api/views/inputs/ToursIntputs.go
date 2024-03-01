package inputs

import (
	"time"

	"gorm.io/gorm"
)

type ToursInputs struct {
	Db         *gorm.DB
	FkCompany  *uint
	FkCustomer *uint
	Plate      *string
	Imei       *string
	FromDate   *time.Time
	ToDate     *time.Time
	Page       *uint
	PageSize   *uint
}
