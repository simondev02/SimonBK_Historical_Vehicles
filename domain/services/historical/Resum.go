package services

import (
	"database/sql"
	"fmt"
)

type Record struct {
	Imei        string
	Plate       string
	Odometer    string
	MinOdometer string
	MaxOdometer string
	Date        string
}

func GetRecords(db *sql.DB) ([]Record, error) {
	query := `
	WITH records AS (
		SELECT imei, plate,
			CAST(properties AS JSON) ->> 'Total Odometer' AS odometer,
			time_stamp_event AT TIME ZONE 'GMT' AT TIME ZONE 'America/Bogota' AS local_time
		FROM avl_records
		WHERE id_company = 21
		AND DATE(time_stamp_event AT TIME ZONE 'GMT' AT TIME ZONE 'America/Bogota') = CURRENT_DATE - INTERVAL '1 day'
	),
	min_records AS (
		SELECT imei,
			MIN(local_time) AS min_time
		FROM records
		GROUP BY imei
	),
	max_records AS (
		SELECT imei,
			MAX(local_time) AS max_time
		FROM records
		GROUP BY imei
	),
	min_odometer AS (
		SELECT r.imei, r.plate, r.odometer AS min_odometer
		FROM records r
		JOIN min_records mr ON r.imei = mr.imei AND r.local_time = mr.min_time
	),
	max_odometer AS (
		SELECT r.imei, r.plate, r.odometer AS max_odometer
		FROM records r
		JOIN max_records mr ON r.imei = mr.imei AND r.local_time = mr.max_time
	)
	SELECT min_odometer.imei, min_odometer.plate,
		(ROUND((CAST(max_odometer.max_odometer AS INTEGER) - CAST(min_odometer.min_odometer AS INTEGER)) / 1000.0)::TEXT || ' Km/h') AS odometer,
		(ROUND(CAST(min_odometer.min_odometer AS INTEGER) / 1000.0)::TEXT || ' Km/h') AS min_odometer,
		(ROUND(CAST(max_odometer.max_odometer AS INTEGER) / 1000.0)::TEXT || ' Km/h') AS max_odometer,
		(CURRENT_DATE - INTERVAL '1 day')::date AS date
	FROM min_odometer
	JOIN max_odometer ON min_odometer.imei = max_odometer.imei;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var record Record
		err := rows.Scan(&record.Imei, &record.Plate, &record.Odometer, &record.MinOdometer, &record.MaxOdometer, &record.Date)
		if err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar las filas: %w", err)
	}

	return records, nil
}
