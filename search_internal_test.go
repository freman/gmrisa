package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	f, err := os.Open("testdata/result.html")
	assert.NoError(err)
	defer f.Close()

	res, err := searchParse(f)
	assert.NoError(err)

	f2, err := os.Open("testdata/result.json")
	assert.NoError(err)
	defer f2.Close()

	var res2 searchResult
	assert.NoError(json.NewDecoder(f2).Decode(&res2))

	assert.Equal(res2, *res)
}
