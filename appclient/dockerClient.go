package appclient

/// strategy //
//  get container name //
//
// pool interval
// get stats memory
// get container info - status
// setup channel and then pass data for display
// keep on updating the ui

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
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
	StatusInfo    ContainerStatus
}

func (dc *DockerClient) NewClient(c chan ContainerStatus) {

	dc.channel = c
	dc.StatusInfo = ContainerStatus{}

	client, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}
	dc.targetClient = client
}

func (dc *DockerClient) GetContainerByName(containerName string) {

	containers := dc.GetContainerInfo()

	for _, container := range containers {

		log.Println("getting out the names", container.Names[0])

		if strings.ToLower(container.Names[0]) == strings.ToLower("/"+containerName) {

			dc.containerId = container.ID

			dc.StatusInfo.Name = container.Names[0]
			dc.StatusInfo.Image = container.Image
			dc.StatusInfo.Status = container.State
			dc.channel <- dc.StatusInfo
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

	buf := new(bytes.Buffer)
	buf.ReadFrom(containerStats.Body)

	dockerStat := new(ContainerStat)
	json.Unmarshal(buf.Bytes(), dockerStat)

	//log.Println(buf.String())
	log.Println(containerStats.OSType)

	if err != nil {
		log.Println("Unable to get container stat.")
	}

	dc.StatusInfo.Stats = dockerStat
	dc.channel <- dc.StatusInfo

	return containerStats
}

func (dc *DockerClient) GetSwarmService() {

	serviceList, _ := dc.targetClient.ServiceList(context.Background(), types.ServiceListOptions{})

	for _, n := range serviceList {
		log.Println("Service List nodes", n.ID)
	}

	taskList, _ := dc.targetClient.TaskList(context.Background(), types.TaskListOptions{})

	for _, n := range taskList {
		log.Println("--------------------------------")
		log.Println("Task List name", n.Name)
		log.Println("Task List id", n.ID)
		log.Println("Task List serviceid", n.ServiceID)
		log.Println("Task List status", n.Status)
		//	log.Println("Task List nodes", n.NetworksAttachments[0].Addresses)

		log.Println("--------------------------------")
	}

}

func (dc *DockerClient) GetDiskUsage() {

	log.Println("getting disk usage data")

	dusage, _ := dc.targetClient.DiskUsage(context.Background())

	var totalVolSize int64
	var totalContainerSize int64
	var totalImageSize int64

	for _, v := range dusage.Volumes {

		totalVolSize += v.UsageData.Size
	}

	for _, c := range dusage.Containers {

		totalContainerSize += c.SizeRw
	}

	for _, i := range dusage.Images {

		totalImageSize += i.Size
	}

	log.Println("printing volume info")

	log.Println(totalVolSize)
	log.Println(totalContainerSize)
	log.Println(totalImageSize)

	dc.StatusInfo.Disk = &TotalDiskUsage{}
	dc.StatusInfo.Disk.Volumes = totalVolSize
	dc.StatusInfo.Disk.Containers = totalContainerSize
	dc.StatusInfo.Disk.Images = totalImageSize

	dc.channel <- dc.StatusInfo
}
