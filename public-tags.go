package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/heroku/docker-registry-client/registry"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Base Docker image is missing. Try alpine, ubuntu, hashicorp/terraform, codefresh/cli etc")
		os.Exit(1)
	}

	baseImage := os.Args[1]

	dockerHubConnection := connectToDockerHub()

	tagsFound := findAllTags(dockerHubConnection, baseImage)

	tagCount := len(tagsFound)

	if tagCount == 0 {
		fmt.Printf("Could not find tags for image %s. Try alpine, ubuntu, hashicorp/terraform, codefresh/cli etc\n", baseImage)
		os.Exit(1)
	}

	newestTag := findNewestSemanticTag(tagsFound)

	if newestTag == "" {
		fmt.Println("Could not any find tags that follow semantic versioning.")
		os.Exit(1)
	}
	fmt.Println(newestTag)

}

func connectToDockerHub() *registry.Registry {
	url := "https://registry.hub.docker.com"
	username := "" // anonymous
	password := "" // anonymous

	transport := http.DefaultTransport
	dockerHubConnection, err := newFromTransport(url, username, password, transport, registry.Quiet)

	if err != nil {
		fmt.Println(err)
	}

	return dockerHubConnection
}

func findAllTags(dockerHubConnection *registry.Registry, baseImage string) []string {
	if strings.Contains(baseImage, "/") == false {
		baseImage = "library/" + baseImage
	}
	tags, err := dockerHubConnection.Tags(baseImage)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tags

}

func newFromTransport(registryURL, username, password string, transport http.RoundTripper, logf registry.LogfCallback) (*registry.Registry, error) {
	url := strings.TrimSuffix(registryURL, "/")
	transport = registry.WrapTransport(transport, url, username, password)
	registry := &registry.Registry{
		URL: url,
		Client: &http.Client{
			Transport: transport,
		},
		Logf: logf,
	}

	if err := registry.Ping(); err != nil {
		return nil, err
	}

	return registry, nil
}

func findNewestSemanticTag(allTags []string) string {
	var validTags []semver.Version

	for _, tag := range allTags {
		version, err := semver.NewVersion(tag)
		if err == nil {
			validTags = append(validTags, *version)
		}
	}

	if len(validTags) == 0 {
		return ""
	}

	sort.Slice(validTags, func(i, j int) bool {
		return validTags[i].LessThan(validTags[j])
	})

	return validTags[len(validTags)-1].String()

}
