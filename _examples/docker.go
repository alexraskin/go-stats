package plugins

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerPlugin struct {
	client *client.Client
}

func NewDockerPlugin() (*DockerPlugin, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &DockerPlugin{client: cli}, nil
}

func (p *DockerPlugin) Name() string {
	return "docker"
}

func (p *DockerPlugin) Metrics() (map[string]interface{}, error) {
	containers, err := p.client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"num_services": len(containers),
	}, nil
}
