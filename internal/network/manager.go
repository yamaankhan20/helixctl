package network

import (
	"context"

	"github.com/yamaankhan20/helixctl/internal/domain"
)

// NetworkManager defines the contract for container network configuration.
type NetworkManager interface {
	Configure(ctx context.Context, containerID domain.ContainerID, pid int, netnsPath string) error
	Release(ctx context.Context, containerID domain.ContainerID) error
}
