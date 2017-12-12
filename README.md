# gommand

A simple way to define multiple application in the same command.

## Usage

```go
package main

import (
	"github.com/cashcowpro/gommand"
	"github.com/cashcowpro/example/monitoring"
	"github.com/cashcowpro/example/research"
)

func main() {
	applications := []gommand.Application{
		research.CreateApplication(),
		monitoring.CreateApplication(),
	}

	gommand.StartApplications(applications)
}
```
