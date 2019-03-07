package appclient

import (
	"github.com/docker/docker/api/types"
)

type ContainerClient interface {
	NewClient()
	GetContainerByName(containerName string)
	GetContainerInfo() []types.Container
	GetContainerStat() types.ContainerStats

	GetSwarmService()
	GetDiskUsage()
}
