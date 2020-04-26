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

// func TestReadMediumDataset(t *testing.T) {

// 	stepsFound := readJSON("testdata/input-medium.json")

// 	assert.Equal(t, 14, len(stepsFound), "Incorrect number of steps found")
// }

// func TestReadFullDataset(t *testing.T) {

// 	stepsFound := readJSON("testdata/input-full.json")

// 	assert.Equal(t, 25, len(stepsFound), "Incorrect number of steps found")
// }
//
