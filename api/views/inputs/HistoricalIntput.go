package inputs

import "time"

type HistoricalIntput struct {
	FkCompany  *uint
	FkCustomer *uint
	Plate      *string
	FromDate   *time.Time
	ToDate     *time.Time
}
