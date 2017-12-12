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
func StartApplications(applications []Application) {
	for _, a := range applications {
		go func(a Application) {
			if error := a.Start(); error != nil {
				golog.Create().Add("application", a).Add("error", error).Fatal("Could start application")
			}
		}(a)
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
