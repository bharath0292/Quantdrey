package util

import (
	"fmt"
	"os"
)

func ValidateEnvVars(requiredEnvVars []string) error {
	missing := []string{}
	for _, key := range requiredEnvVars {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missing)
	}

	return nil
}
