package ntp

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func GetCurrentTime() (time.Time, error) {
	exTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get NTP time: %w", err)
	}

	return exTime, nil
}
