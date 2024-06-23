package helpers

import "time"

func UnixTimeToFormattedString(unixTime int64) string {
	// Convert Unix timestamp to time.Time object
	t := time.Unix(unixTime, 0)

	// Format the time using the desired layout
	formattedString := t.Format("2006-01-02 15:04:05")

	return formattedString
}
