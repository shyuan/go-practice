package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func simulate(samples int) float64 {
	var inside float64

	for i := 0; i < samples; i++ {
		if x, y := rand.Float64(), rand.Float64(); (x*x + y*y) <= 1 {
			inside++
		}
	}

	return (inside / float64(samples) * 4)
}

func main() {
	runs := []struct {
		samples int
	}{
		{100},
		{1000},
		{10000},
		{100000},
		{1000000},
		{10000000},
		{100000000},
	}

	for _, run := range runs {
		fmt.Printf("pi=%f runs=%d\n", simulate(run.samples), run.samples)
	}
}
