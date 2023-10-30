package services

import "time"

// Formato Fecha
func FormatTimestamp(timestamp string) (string, error) {
	layout := "2006-01-02T15:04:05Z07:00"
	t, err := time.Parse(layout, timestamp)
	if err != nil {
		return "", err
	}
	return t.Format("02/01/2006 15:04"), nil
}
