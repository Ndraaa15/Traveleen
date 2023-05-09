package time

import (
	"fmt"
	"strings"
	"time"
)

func GenerateDate() string {
	d := time.Now()
	formattedTime := d.Format("Monday, 02 January 2006")

	timeParts := strings.Split(formattedTime, " ")

	if timeParts[0] == "Monday," {
		timeParts[0] = "Senin,"
	} else if timeParts[0] == "Tuesday," {
		timeParts[0] = "Selasa,"
	} else if timeParts[0] == "Wednesday," {
		timeParts[0] = "Rabu,"
	} else if timeParts[0] == "Thursday," {
		timeParts[0] = "Kamis,"
	} else if timeParts[0] == "Friday," {
		timeParts[0] = "Jumat,"
	} else if timeParts[0] == "Saturday," {
		timeParts[0] = "Sabtu,"
	} else if timeParts[0] == "Sunday," {
		timeParts[0] = "Minggu,"
	}

	if timeParts[2] == "January" {
		timeParts[2] = "Januari"
	} else if timeParts[2] == "February" {
		timeParts[2] = "Februari"
	} else if timeParts[2] == "March" {
		timeParts[2] = "Maret"
	} else if timeParts[2] == "April" {
		timeParts[2] = "April"
	} else if timeParts[2] == "May" {
		timeParts[2] = "Mei"
	} else if timeParts[2] == "June" {
		timeParts[2] = "Juni"
	} else if timeParts[2] == "July" {
		timeParts[2] = "Juli"
	} else if timeParts[2] == "August" {
		timeParts[2] = "Agustus"
	} else if timeParts[2] == "September" {
		timeParts[2] = "September"
	} else if timeParts[2] == "October" {
		timeParts[2] = "Oktober"
	} else if timeParts[2] == "November" {
		timeParts[2] = "November"
	} else if timeParts[2] == "December" {
		timeParts[2] = "Desember"
	}

	date := fmt.Sprintf("%s %s %s %s", timeParts[0], timeParts[1], timeParts[2], timeParts[3])

	return date
}
