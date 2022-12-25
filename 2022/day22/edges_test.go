package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCorrespondingEdge(t *testing.T) {
	assert.Equal(t, UF, FU.getCorrespondingEdge())
	assert.Equal(t, BR, RB.getCorrespondingEdge())
}
