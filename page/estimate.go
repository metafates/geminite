package page

import (
	"math"
	"strings"
	"time"
)

const wordsPerMinute = 200

func estimateReadingDuration(text string) time.Duration {
	if len(text) == 0 {
		return 0
	}

	words := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == '\n' || r == '\r' || r == '\t'
	})

	wordsCount := len(words)

	minutes := math.Ceil(float64(wordsCount) / float64(wordsPerMinute))
	return time.Duration(math.Ceil(minutes) * float64(time.Minute))
}
