package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	returnTime, _ := time.Parse("1/02/2006 15:04:05", date)
	return returnTime
	//panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	returnTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(returnTime)
	//panic("Please implement the HasPassed function")
	//July 25, 2019 13:45:00
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	returnTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if returnTime.Hour() >= 12 && returnTime.Minute() >= 0 && returnTime.Hour() <= 18 {
		return true
	}
	return false
	//Thursday, July 25, 2019 13:45:00
	//panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
// 7/25/2019 13:45:00
func Description(date string) string {
	returnTime, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", 
		returnTime.Weekday(), 
		returnTime.Month(), 
		returnTime.Day(), 
		returnTime.Year(), 
		returnTime.Hour(), 
		returnTime.Minute())
	//panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(),time.September,15,0,0,0,0,time.UTC)
	//panic("Please implement the AnniversaryDate function")
}
