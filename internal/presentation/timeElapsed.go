package presentation

import (
	"fmt"
	"time"
)

var startTime time.Time

func StartTimer() {
	startTime = time.Now()
}

func PrintTimeElapsed() {
	timeElapsed := time.Since(startTime)
	fmt.Println("Time elapsed:", timeElapsed.String())
}
