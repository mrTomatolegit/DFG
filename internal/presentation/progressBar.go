package presentation

import (
	"fmt"
	"github.com/gosuri/uilive"
	"math"
	"time"
)

const filled = "="
const empty = " "

func CreateUpdateTicker(currentCount *int, objectiveCount *int) (*uilive.Writer, chan struct{}) {
	writer := uilive.New()
	writer.Start()

	ticker := time.NewTicker(500 * time.Millisecond)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				UpdateProgressBar(writer, *currentCount, *objectiveCount)
			case <-quit:
				ticker.Stop()
				writer.Stop()
				return
			}
		}
	}()
	return writer, quit
}

func makeProgressBar(percentage float64) string {
	var quarterP = int(math.Round(float64(percentage / 4)))
	var bar string
	for i := 0; i < quarterP; i++ {
		bar += filled
	}
	for i := quarterP; i < 25; i++ {
		bar += empty
	}
	return bar
}

func MakeProgressLine(currentCount int, objectiveCount int) string {
	var percentage = float64(currentCount) / float64(objectiveCount) * 100
	return fmt.Sprintf("%v/%v [%s] %v%%\n", currentCount, objectiveCount, makeProgressBar(percentage), math.Floor(percentage))
}

func UpdateProgressBar(writer *uilive.Writer, currentCount int, objectiveCount int) {
	fmt.Fprint(writer, MakeProgressLine(currentCount, objectiveCount))
}
