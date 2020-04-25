package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/heroku/docker-registry-client/registry"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Base Docker image is missing. Try alpine, ubuntu, hashicorp/terraform, codefresh/cli etc")
		os.Exit(1)
	}

	dockeHubConnection := connectToDockerHub()

	tagsFound := findAllTags(dockeHubConnection, os.Args[1])

	log.Printf("Found %d tags for the base image\n", len(tagsFound))

}

func connectToDockerHub() *registry.Registry {
	url := "https://registry.hub.docker.com"
	username := "" // anonymous
	password := "" // anonymous
	dockerHubConnection, err := registry.New(url, username, password)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection successful")

	return dockerHubConnection
}

func findAllTags(dockerHubConnection *registry.Registry, baseImage string) []string {
	if strings.Contains(baseImage, "/") == false {
		baseImage = "library/" + baseImage
	}
	tags, err := dockerHubConnection.Tags(baseImage)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("Found Dockerhub tags: ", tags)
	return tags

	// for _, tag := range tags {
	// 	if tag == imageAndTag.Tag {
	// 		return true
	// 	}
	// }

}
