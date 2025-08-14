package strategyutils

import (
	"errors"
	"strconv"
	"strings"
)

func ValidateStrategyTime(timeStr string) error {
	timePart := strings.Split(timeStr, ":")

	_, err := strconv.Atoi(timePart[0])
	if err != nil {
		return errors.New("error parsing hour")
	}

	_, err = strconv.Atoi(timePart[1])
	if err != nil {
		return errors.New("error parsing minute")
	}
	_, err = strconv.Atoi(timePart[2])
	if err != nil {
		return errors.New("error parsing second")
	}

	return nil
}
