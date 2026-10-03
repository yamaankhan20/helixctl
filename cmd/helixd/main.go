package main

import (
	"fmt"

	"github.com/yamaankhan20/helixctl/internal/config"
	"github.com/yamaankhan20/helixctl/internal/observability"
)

func main() {
	logger := observability.NewLogger("helixd")
	logger.Info("starting helixd", "config", config.ControlPlane{})
	fmt.Println("helixd daemon")
}
