package postgresql

import (
	"errors"
	"fmt"
	errorsext "gitlab.com/ricardo-public/errors/v2/pkg/errors"
	"gorm.io/gorm"
)

func notFoundOrElseError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("postgresql: %s: %w", errorsext.ErrNotFound, err)
	}

	return fmt.Errorf("postgresql: %s: %w", errorsext.ErrTimeout, err)
}
