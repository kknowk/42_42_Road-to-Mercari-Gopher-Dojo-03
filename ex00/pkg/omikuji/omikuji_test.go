// ex00/pkg/omikuji/omikuji.go
package omikuji

import (
	mytime "omikuji/pkg/time"
	"testing"
	"time"
)

func HelperIsNewYear(t *testing.T, year int, month time.Month, day int) func() time.Time {
	t.Helper()
	return func() time.Time {
		return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	}
}

func TestRandomOmikuji(t *testing.T) {

	originalCurrentTime := mytime.CurrentTime
	defer func() { mytime.CurrentTime = originalCurrentTime }()

	tests := []struct {
		name        string
		currentTime func() time.Time
		want        bool
	}{
		{name: "New Year's Day", currentTime: HelperIsNewYear(t, 2021, 1, 1), want: true},
		{name: "New Year's Day", currentTime: HelperIsNewYear(t, 2021, 1, 2), want: true},
		{name: "New Year's Day", currentTime: HelperIsNewYear(t, 2021, 1, 3), want: true},
		{name: "New Year's Day", currentTime: HelperIsNewYear(t, 2021, 1, 4), want: true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mytime.CurrentTime = tt.currentTime
			got := RandomOmikuji()
			current := mytime.CurrentTime()
			if current.Month() == time.January && current.Day() >= 1 && current.Day() <= 3 {
				if got.Fortune != "Dai-kichi" {
					t.Errorf("RandomOmikuji() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
