package services

import (
	"log"
	"time"

	"github.com/appcoreopc/dockerMonitor/appclient"
)

type AppService struct {
}

func (ap *AppService) Start(instanceName string) {

	log.Println("start my services")

	defer ap.AppServiceRecovery()

	ap.KickOffTimer(instanceName)
}

func (ap *AppService) AppServiceRecovery() {

	if r := recover(); r != nil {
		log.Println("recover")
	}
}

func (ap *AppService) Execute() {

}

func (ap *AppService) KickOffTimer(instanceName string) {

	log.Println("kicking off my timer")

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
		if cs.Stats != nil {
			log.Println("Memory")
			log.Println("Limit", cs.Stats.Memory_stats.Limit)
			log.Println("Usage", cs.Stats.Memory_stats.Usage)
		}

	}

}
