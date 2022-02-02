// struct demo
package main

import (
	"fmt"
	"time"
)

// Budget is a budgest for campaign
type Budget struct {
	CampaignID string
	Balance    float64 // USD
	Expires    time.Time
}

func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

func (b *Budget) Update(sum float64) {
	b.Balance += sum
}

func main() {
	b := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b.TimeLeft())

	b.Update(10.5)
	fmt.Println(b.Balance)

}
