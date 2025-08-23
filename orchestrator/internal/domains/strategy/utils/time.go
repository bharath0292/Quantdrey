package strategyutils

import (
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func SleepUntil(startTime time.Time) {
	now := time.Now()
	if now.Before(startTime) {
		sleepDuration := time.Until(startTime)
		log.Info().Msgf("sleeping for %s", sleepDuration)
		time.Sleep(sleepDuration)
	}
}

func ParseTime(timeStr string) *time.Time {
	timePart := strings.Split(timeStr, ":")

	// No error handling because we validated while creating
	hour, _ := strconv.Atoi(timePart[0])
	minute, _ := strconv.Atoi(timePart[1])
	second, _ := strconv.Atoi(timePart[2])

	t := time.Now()
	created := time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, 0, t.Location())

	return &created
}
