/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runContainer()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	containerCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// mapstructure tag would map the values from the config file to the struct

type DockerConfig struct {
	ImageName  string               `mapstructure:"imageName"`
	Cfg        container.Config     `mapstructure:"cfg"`
	HostConfig container.HostConfig `mapstructure:"hostConfig"`
}

func runContainer() error {
	var cfg DockerConfig
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return fmt.Errorf("failed to parse configuration: %w", err)
	}
	//fmt.Printf("Config: %#v\n", cfg)

	// Initialize Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	//NegotiateAPIVersion would allow us to run the containers even if our docker in system is outdated
	cli.NegotiateAPIVersion(context.Background())

	reader, err := cli.ImagePull(context.Background(), cfg.ImageName, image.PullOptions{})
	if err != nil {
		// If there's an error while pulling the image, panic and exit
		return err
	}
	defer reader.Close() // Clean up resources when they're no longer needed

	io.Copy(os.Stdout, reader)

	fmt.Println("Docker image pulled successfully")

	// Create Docker container
	resp, err := cli.ContainerCreate(context.Background(), &cfg.Cfg, &cfg.HostConfig, nil, nil, "")
	if err != nil {
		return fmt.Errorf("failed to create Docker container: %w", err)
	}

	fmt.Println("Container created successfully")

	// Start Docker container
	if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start Docker container: %w", err)
	}

	// Get the logs of the container
	out, err := cli.ContainerLogs(context.Background(), resp.ID, container.LogsOptions{ShowStdout: true})
	if err != nil {
		// if there's an error while getting logs, panic and exit
		return err
	}

	// copy the container logs to the standard output
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	fmt.Println("Docker container started successfully with ID:", resp.ID)
	return nil

}

/*
func (cli *Client) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *v1.Platform, containerName string) (container.CreateResponse, error)

*/
