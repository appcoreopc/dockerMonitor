package appclient

/// strategy //
// get container name //
//
// pool interval
// get stats memory
// get container info - status
// setup channel and then pass data for display
// keep on updating the ui

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	Hostname      string
	ContainerName string
	targetClient  *client.Client
	channel       chan ContainerStatus
	containerId   string
}

func (dc *DockerClient) NewClient(c chan ContainerStatus) {

	dc.channel = c

	client, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	dc.targetClient = client

}

func (dc *DockerClient) GetContainerByName(containerName string) {

	containers := dc.GetContainerInfo()

	for _, container := range containers {

		fmt.Println(container.Names[0])

		if strings.ToLower(container.Names[0]) == strings.ToLower("/"+containerName) {

			dc.containerId = container.ID

			dataUpdate := ContainerStatus{}
			dataUpdate.Name = container.Names[0]
			dataUpdate.Status = container.State
			dc.channel <- dataUpdate
		}
	}
}

func (dc *DockerClient) GetContainerInfo() []types.Container {

	containers, err := dc.targetClient.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	return containers
}

func (dc *DockerClient) GetContainerStat() types.ContainerStats {

	containerStats, err := dc.targetClient.ContainerStats(context.Background(), dc.containerId, false)

	if err != nil {
		panic("Unable to get container stat.")
	}

	dataUpdate := ContainerStatus{}
	dataUpdate.Name = "test"
	dataUpdate.Status = "Extended"
	dc.channel <- dataUpdate

	return containerStats

}
