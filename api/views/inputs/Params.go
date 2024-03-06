package inputs

import (
	"time"

	"gorm.io/gorm"
)

type Params struct {
	Db         *gorm.DB
	FkCompany  *uint
	FkCustomer *uint
	Plate      *string
	Imei       *string
	FromDate   *time.Time
	ToDate     *time.Time
	Page       int
	PageSize   int
}
