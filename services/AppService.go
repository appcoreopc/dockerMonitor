package services

import (
	"fmt"
	"time"

	"github.com/appcoreopc/dockerMonitor/appclient"
)

type AppService struct {
}

func (ap *AppService) Start(instanceName string) {

	fmt.Println("start my services")

	defer ap.AppServiceRecovery()

	ap.KickOffTimer(instanceName)
}

func (ap *AppService) AppServiceRecovery() {

	if r := recover(); r != nil {
		fmt.Println("recover")
	}
}

func (ap *AppService) Execute() {

}

func (ap *AppService) KickOffTimer(instanceName string) {

	fmt.Println("kicking off my timer")

	statusChannel := make(chan appclient.ContainerStatus, 5)
	docker := new(appclient.DockerClient)
	docker.NewClient(statusChannel)

	// for cs := range statusChannel {
	// 	fmt.Println("Giving the proper state")
	// 	fmt.Println(cs.Status)
	// }

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	go func() {

		for {
			select {
			case <-ticker.C:
				fmt.Println("Timer elapsed! ")
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
		fmt.Println("Giving the proper state")
		fmt.Println(cs.Status)
	}

}
