package util

import (
	"errors"
	"os"
)

func IsSelfSignedCertAvailable(certPath, keyPath string) error {
	if _, errCert := os.Stat(certPath); os.IsNotExist(errCert) {
		return errors.New("cert not available")
	}

	if _, errKey := os.Stat(keyPath); os.IsNotExist(errKey) {
		return errors.New("key not available")
	}

	return nil
}
