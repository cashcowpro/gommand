package gommand

import (
	"time"

	"github.com/cashcowpro/golog"
)

// Application defines an interface to all applications.
type Application interface {
	Start() error
}

// StartApplications defined in the array.
func StartApplications(logger golog.Logger, applications []Application) {
	for _, a := range applications {
		go func(a Application) {
			if err := a.Start(); err != nil {
				logger.FatalError("Could start application", err)
			}
		}(a)
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
