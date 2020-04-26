package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitDiscardInvalid(t *testing.T) {

	var allTags = []string{"one", "invalid-1", "prod", "latest"}

	newest := findNewestSemanticTag(allTags)

	assert.Empty(t, newest, "Invalid tags were not discarded")

}

func TestUnitNewest(t *testing.T) {

	var allTags = []string{"2.1.1", "2.4.34", "1.9.12", "latest", "pre-2"}

	newest := findNewestSemanticTag(allTags)

	assert.Equal(t, "2.4.34", newest, "Newest tag was not found")
}

func TestUnitwithZeroVersions(t *testing.T) {

	var allTags = []string{"0.11.1", "0.24", "0.9.12", "latest", "0.1.2"}

	newest := findNewestSemanticTag(allTags)

	assert.Equal(t, "0.11.1", newest, "Newest tag was not found")
}

func TestIntegrationUbuntu(t *testing.T) {

	dockerHubConnection := connectToDockerHub()

	tagsFound := findAllTags(dockerHubConnection, "ubuntu")

	assert.LessOrEqual(t, 4, len(tagsFound), "Ubuntu should have more than 4 tags in Dockerhub")
}
