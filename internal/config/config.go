// Package config defines the system configuration structures.
package config

type ControlPlane struct {
	RESTListen string
	GRPCListen string
}

type Agent struct {
	ControlPlaneAddress string
	RPCListen           string
	RPCAdvertise        string
	NodeID              string
	ManagementIP        string
	ContainerCIDR       string
}
