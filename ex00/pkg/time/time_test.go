// package time
package time

import (
	"testing"
	"time"
)

func HelperIsNewYear(t *testing.T, year int, month time.Month, day int) func() time.Time {
	t.Helper()
	return func() time.Time {
		return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	}
}

func TestIsNewYear(t *testing.T) {

	// 初期化されていないとCurrentTime()がゼロ値を返す
	
	originalCurrentTime := CurrentTime
	result := CurrentTime()
	if result.IsZero() {
		t.Fatal("CurrentTime() returned zero value")
	}
	defer func() { CurrentTime = originalCurrentTime }()

	tests := [] struct {
		name string
		currentTime func() time.Time
		want bool
	} {
		{name: "New Year's Day", currentTime: HelperIsNewYear(t, 2021, 1, 1), want: true},
		{name: "New Year's Day", currentTime: HelperIsNewYear(t, 2021, 1, 2), want: true},
		{name: "New Year's Day", currentTime: HelperIsNewYear(t, 2021, 1, 3), want: true},
		{name: "Not New Year's Day", currentTime: HelperIsNewYear(t, 2021, 1, 4), want: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsNewYear(tt.currentTime()); got != tt.want {
				t.Errorf("IsNewYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
