package gommand

import (
	"sync"

	"github.com/cashcowpro/golog"
)

// Application defines an interface to all applications.
type Application interface {
	Start() error
}

// StartApplications defined in the array.
func StartApplications(logger golog.Logger, applications []Application) {
	wg := &sync.WaitGroup{}

	for _, a := range applications {
		wg.Add(1)

		go func(a Application, wg *sync.WaitGroup) {
			defer wg.Done()

			if err := a.Start(); err != nil {
				logger.Error("Could start application", err)
			}
		}(a, wg)
	}

	wg.Wait()
}
