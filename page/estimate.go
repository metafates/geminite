package page

import (
	"math"
	"strings"
	"time"

	"github.com/metafates/geminite/config"
)

func estimateReadingDuration(text string) time.Duration {
	if len(text) == 0 {
		return 0
	}

	words := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == '\n' || r == '\r' || r == '\t'
	})

	wordsCount := len(words)

	minutes := math.Ceil(float64(wordsCount) / float64(config.Config.WPM))
	return time.Duration(math.Ceil(minutes) * float64(time.Minute))
}
