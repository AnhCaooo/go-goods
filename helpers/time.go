package helpers

import (
	"fmt"
	"time"
)

const (
	DATE_FORMAT    string = "2006-01-02"          // this is just the layout of YYYY-MM-DD
	UI_DATE_FORMAT string = "2006-01-02 15:04:05" // this is just the layout of YYYY-MM-DD HH:MM:SS
)

// SetTime returns current date based on application's time with specific hour and minute.
// Parameters:
//   - hour: The hour to set (0-23).
//   - minute: The minute to set (0-59).
//
// Returns:
//   - time.Time: The time set to the specified hour and minute in UTC.
//   - error: An error if the Helsinki timezone could not be loaded.
func SetTime(hour int, minute int) (time.Time, error) {
	year, month, day := time.Now().Date()
	settingTime := time.Date(year, month, day, hour, minute, 0, 0, time.Local)
	return settingTime, nil
}

// LoadHelsinkiLocation loads the time.Location for Helsinki, Finland.
// It returns a pointer to the time.Location and an error if the location
// could not be loaded.
func LoadHelsinkiLocation() (*time.Location, error) {
	location, err := time.LoadLocation("Europe/Helsinki")
	if err != nil {
		return nil, fmt.Errorf("failed to get current location: %s", err.Error())
	}
	return location, nil
}

// getTodayDate returns date of today in "YYYY-MM-DD" format
func GetTodayDate() string {
	today := time.Now()
	return today.Format(DATE_FORMAT)
}

// GetTomorrowDate returns date of tomorrow in "YYYY-MM-DD" format
func GetTomorrowDate() string {
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	return tomorrow.Format(DATE_FORMAT)
}

// getYesterdayDate returns date of yesterday in "YYYY-MM-DD" format
func GetYesterdayDate() string {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	return yesterday.Format(DATE_FORMAT)
}
