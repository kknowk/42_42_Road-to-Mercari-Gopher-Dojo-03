// package time
package time

import "time"

// mock time
// CurrentTime is a function to get current time
var CurrentTime = func() time.Time {
	return time.Now()
}

func IsNewYear(t time.Time) bool {
	month := t.Month()
	day := t.Day()

	return month == time.January && (day == 1 || day == 2 || day == 3)
}