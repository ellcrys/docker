package main

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ellcrys/docker/client"
)

const composeTemplate = `# generated by integration-cli-on-swarm
version: "3"

services:
  worker:
    image: "{{.WorkerImage}}"
    command: ["-worker-image-digest={{.WorkerImageDigest}}", "-dry-run={{.DryRun}}", "-keep-executor={{.KeepExecutor}}"]
    networks:
      - net
    volumes:
# Bind-mount the API socket so that we can invoke "docker run --privileged" within the service containers
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DOCKER_GRAPHDRIVER={{.EnvDockerGraphDriver}}
      - DOCKER_EXPERIMENTAL={{.EnvDockerExperimental}}
    deploy:
      mode: replicated
      replicas: {{.Replicas}}
      restart_policy:
# The restart condition needs to be any for funker function
        condition: any

  master:
    image: "{{.MasterImage}}"
    command: ["-worker-service=worker", "-input=/mnt/input", "-chunks={{.Chunks}}", "-shuffle={{.Shuffle}}", "-rand-seed={{.RandSeed}}"]
    networks:
      - net
    volumes:
      - {{.Volume}}:/mnt
    deploy:
      mode: replicated
      replicas: 1
      restart_policy:
        condition: none
      placement:
# Make sure the master can access the volume
        constraints: [node.id == {{.SelfNodeID}}]

networks:
  net:

volumes:
  {{.Volume}}:
    external: true
`

type composeOptions struct {
	Replicas     int
	Chunks       int
	MasterImage  string
	WorkerImage  string
	Volume       string
	Shuffle      bool
	RandSeed     int64
	DryRun       bool
	KeepExecutor bool
}

type composeTemplateOptions struct {
	composeOptions
	WorkerImageDigest     string
	SelfNodeID            string
	EnvDockerGraphDriver  string
	EnvDockerExperimental string
}

// createCompose creates "dir/docker-compose.yml".
// If dir is empty, TempDir() is used.
func createCompose(dir string, cli *client.Client, opts composeOptions) (string, error) {
	if dir == "" {
		var err error
		dir, err = ioutil.TempDir("", "integration-cli-on-swarm-")
		if err != nil {
			return "", err
		}
	}
	resolved := composeTemplateOptions{}
	resolved.composeOptions = opts
	workerImageInspect, _, err := cli.ImageInspectWithRaw(context.Background(), defaultWorkerImageName)
	if err != nil {
		return "", err
	}
	if len(workerImageInspect.RepoDigests) > 0 {
		resolved.WorkerImageDigest = workerImageInspect.RepoDigests[0]
	} else {
		// fall back for non-pushed image
		resolved.WorkerImageDigest = workerImageInspect.ID
	}
	info, err := cli.Info(context.Background())
	if err != nil {
		return "", err
	}
	resolved.SelfNodeID = info.Swarm.NodeID
	resolved.EnvDockerGraphDriver = os.Getenv("DOCKER_GRAPHDRIVER")
	resolved.EnvDockerExperimental = os.Getenv("DOCKER_EXPERIMENTAL")
	composeFilePath := filepath.Join(dir, "docker-compose.yml")
	tmpl, err := template.New("").Parse(composeTemplate)
	if err != nil {
		return "", err
	}
	f, err := os.Create(composeFilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if err = tmpl.Execute(f, resolved); err != nil {
		return "", err
	}
	return composeFilePath, nil
}
