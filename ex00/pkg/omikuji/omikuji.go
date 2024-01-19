// ex00/pkg/omikuji/omikuji.go
package omikuji

import (
	"math/rand"
	"time"
	mytime "omikuji/pkg/time"
)

// Omikuji is a struct for omikuji
type OmikujiResult struct {
	Fortune   string `json:"fortune"`
	Health    string `json:"health"`
	Residence string `json:"residence"`
	Travel    string `json:"travel"`
	Study     string `json:"study"`
	Love      string `json:"love"`
}

func RandomOmikuji() OmikujiResult {
	// random seed
	rand.Seed(time.Now().UnixNano())

	fortunes := []string{"Dai-kichi", "Kichi", "Chuu-kichi", "Sho-kichi",
		"Sue-kichi", "Kyo", "Dai-kyo"}
	
	Health := []string{"Excellent", "Good", "Average", "Bad", "Worst"}
	Residence := []string{"Excellent", "Good", "Average", "Bad", "Worst"}
	Travel := []string{"Excellent", "Good", "Average", "Bad", "Worst"}
	Study := []string{"Excellent", "Good", "Average", "Bad", "Worst"}
	Love := []string{"Excellent", "Good", "Average", "Bad", "Worst"}

	var fortuneResult string
	if mytime.IsNewYear(mytime.CurrentTime()) {
		fortuneResult = fortunes[0]
	} else {
		fortuneResult = fortunes[rand.Intn(7)]
	}

	// 意図的にバグを仕込む　
	// fortuneResult = fortunes[rand.Intn(7)]

	// return omikuji result
	return OmikujiResult{
		fortuneResult,
		Health[rand.Intn(5)],
		Residence[rand.Intn(5)],
		Travel[rand.Intn(5)],
		Study[rand.Intn(5)],
		Love[rand.Intn(5)],
	}
}
