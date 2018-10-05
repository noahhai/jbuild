package jbuild_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/noahhai/jbuild"
	"github.com/stretchr/testify/assert"
)

func TestAddMap(t *testing.T) {
	// arrange
	j := jbuild.Jmap{"j": "v2"}

	// act
	j.AddMap(jbuild.Jmap{"v2": "v5"}, "j")
	j.AddMap(jbuild.Jmap{"l": "v2"}, "k", "v1")
	j.AddMap(jbuild.Jmap{"m": "v3"}, "k", "v1")

	//assert
	expectedJson := "{\"j\":{\"v2\":\"v5\"},\"k\":{\"v1\":{\"l\":\"v2\",\"m\":\"v3\"}}}"

	m, err := json.Marshal(j)
	assert.Nil(t, err)
	assert.Equal(t, expectedJson, strings.TrimSpace(string(m)))
}

func TestAdd(t *testing.T) {
	// arrange
	j := jbuild.Jmap{"j": "v2"}

	// act
	j.Add("v2", "k", "v1", "l")
	j.Add("v3", "k", "v1", "m")

	//assert
	expectedJson := "{\"j\":\"v2\",\"k\":{\"v1\":{\"l\":\"v2\",\"m\":\"v3\"}}}"

	m, err := json.Marshal(j)
	assert.Nil(t, err)
	assert.Equal(t, expectedJson, strings.TrimSpace(string(m)))
}
