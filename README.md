# gommand

A simple way to define multiple applications in the same command.

## Usage

```go
package main

import (
    "github.com/cashcowpro/golog"
    "github.com/cashcowpro/gommand"
    "github.com/cashcowpro/example/monitoring"
    "github.com/cashcowpro/example/research"
)

func main() {
    logger := golog.CreateStdout()
    applications := []gommand.Application{
        research.CreateApplication(),
        monitoring.CreateApplication(),
    }

    gommand.StartApplications(logger, applications)
}
```
