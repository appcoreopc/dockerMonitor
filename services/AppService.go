package services

import (
	"log"
	"runtime/debug"
	"time"

	"github.com/appcoreopc/dockerMonitor/appclient"
)

type AppService struct {
}

func (ap *AppService) Start(instanceName string) {

	log.Println("Start monitoring docker container instance.")

	defer ap.AppServiceRecovery()

	ap.KickOffTimer(instanceName)
}

func (ap *AppService) AppServiceRecovery() {

	if err := recover(); err != nil {
		log.Println(err)
		debug.PrintStack()
	}
}

func (ap *AppService) Execute() {

}

func (ap *AppService) KickOffTimer(instanceName string) {

	statusChannel := make(chan appclient.ContainerStatus, 5)
	docker := new(appclient.DockerClient)
	docker.NewClient(statusChannel)

	ticker := time.NewTicker(5 * time.Second)

	quit := make(chan struct{})

	go func() {

		for {
			select {
			case <-ticker.C:
				docker.GetContainerByName(instanceName)
				docker.GetContainerStat()
				docker.GetDiskUsage()
				docker.GetSwarmService()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// block forever //

	for cs := range statusChannel {

		log.Println(cs.Name)
		log.Println(cs.Status)
		log.Println(cs.Image)
		if cs.Stats != nil {
			log.Println("Memory")
			log.Println("Limit", cs.Stats.Memory_stats.Limit)
			log.Println("Usage", cs.Stats.Memory_stats.Usage)
		}

		if cs.Disk != nil {

			log.Println("Total data volumne used", cs.Disk.Volumes)
			log.Println("Total container size", cs.Disk.Containers)
			log.Println("Total image size", cs.Disk.Images)

		}

	}
}
