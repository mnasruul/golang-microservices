package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateReporequest{
		Name:         "golang intruduction",
		Descrtiption: "a golang intruduction",
		Homepage:     "https://github.com",
		Private:      true,
		HasIssues:    true,
		HasProjects:  true,
		HasWiki:      true,
	}
	//Marshal takes an input interface and attemps to creare a valid json string
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))

	var target CreateReporequest
	// Unmarshall takes an input byte array a *pointer* that we're trying to using this json.
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
}
