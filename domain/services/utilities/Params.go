package utilities

import (
	"SimonBK_Historical_Vehicles/api/views/inputs"
	"SimonBK_Historical_Vehicles/infra/db"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetParams(c *gin.Context) (inputs.Params, error) {
	var params inputs.Params

	params.Db = db.DBConn

	if temp, exists := c.Get("FkCompany"); exists {
		val := uint(temp.(int))
		params.FkCompany = &val
	}

	if temp, exists := c.Get("FkCustomer"); exists {
		val := uint(temp.(int))
		params.FkCustomer = &val
	}

	if temp := c.Query("Plate"); temp != "" {
		params.Plate = &temp
	}

	if temp := c.Query("Imei"); temp != "" {
		params.Imei = &temp
	}

	if temp := c.Query("fromDate"); temp != "" {
		val, err := time.Parse(time.RFC3339, temp)
		if err != nil {
			return params, fmt.Errorf("error parsing fromDate: %v", err)
		}
		params.FromDate = &val
	}

	if temp := c.Query("toDate"); temp != "" {
		val, err := time.Parse(time.RFC3339, temp)
		if err != nil {
			return params, fmt.Errorf("error parsing toDate: %v", err)
		}
		params.ToDate = &val
	}

	if temp := c.Query("page"); temp != "" {
		val, err := strconv.Atoi(temp)
		if err != nil {
			return params, fmt.Errorf("error parsing page: %v", err)
		}
		params.Page = val
	}

	if temp := c.Query("pageSize"); temp != "" {
		val, err := strconv.Atoi(temp)
		if err != nil {
			return params, fmt.Errorf("error parsing pageSize: %v", err)
		}
		params.PageSize = val
	}

	return params, nil
}
