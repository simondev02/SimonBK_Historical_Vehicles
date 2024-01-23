package services

import (
	"database/sql"
	"fmt"
	"net/smtp"
	"strings"
)

func SendRecordsByEmail(db *sql.DB) error {
	records, err := GetRecords(db)
	if err != nil {
		return fmt.Errorf("error al obtener los registros: %w", err)
	}

	var body strings.Builder
	for _, record := range records {
		body.WriteString(fmt.Sprintf("Imei: %s, Plate: %s, Odometer: %s, MinOdometer: %s, MaxOdometer: %s, Date: %s\n",
			record.Imei, record.Plate, record.Odometer, record.MinOdometer, record.MaxOdometer, record.Date))
	}

	from := "secamc93@gmail.com"
	password := "sebastian230993"
	to := "secamc93@gmail.com"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := "From: Yo" + from + "MI" +
		"To: " + to + "\n" +
		"Subject: Resultados de los registros\n\n" +
		body.String()

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("error al enviar el correo electrónico: %w", err)
	}

	return nil
}
