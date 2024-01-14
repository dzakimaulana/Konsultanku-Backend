package functions

import (
	"time"
)

func DateConvert(dateString string) (string, error) {
	dateFormat := "02/01/2006"
	parsedDate, err := time.Parse(dateFormat, dateString)
	if err != nil {
		return "", err
	}
	return parsedDate.Format("2006-01-02"), nil
}

func CurrentTime() string {
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	return timeString
}
