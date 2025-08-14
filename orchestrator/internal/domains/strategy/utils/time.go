package strategyutils

import (
	"errors"
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

func ParseTime(timeStr string) (*time.Time, error) {
	timePart := strings.Split(timeStr, ":")

	hour, err := strconv.Atoi(timePart[0])
	if err != nil {
		return nil, errors.New("error parsing hour")
	}

	minute, err := strconv.Atoi(timePart[1])
	if err != nil {
		return nil, errors.New("error parsing minute")
	}
	second, err := strconv.Atoi(timePart[2])
	if err != nil {
		return nil, errors.New("error parsing second")
	}

	t := time.Now()
	created := time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, 0, t.Location())

	return &created, nil
}
