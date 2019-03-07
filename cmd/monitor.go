// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/appcoreopc/dockerMonitor/appclient"
	"github.com/appcoreopc/dockerMonitor/services"
	"github.com/spf13/cobra"
)

var instanceName string = ""

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		statusChannel := make(chan appclient.ContainerStatus, 5)
		quit := make(chan struct{})

		dockerClient := appclient.DockerClient{Channel: statusChannel}

		as := services.AppService{
			RestClient:    &dockerClient, // & to pass object to a interface type
			StatusChannel: statusChannel,
			Quit:          quit}

		as.Start(instanceName)
	},
}

func init() {

	monitorCmd.Flags().StringVarP(&instanceName, "Instance", "i", "instance", "please specify container instance name.")
	monitorCmd.MarkFlagRequired("Instance")
	rootCmd.AddCommand(monitorCmd)
}
