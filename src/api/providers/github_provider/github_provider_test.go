package github_provider

import (
	"errors"
	"github.com/mnasruul/golang-microservices/src/api/clients/resclient"
	"github.com/mnasruul/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	resclient.StartMockups()
	resclient.AddMockup(resclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	response, err := CreateRepo("", github.CreateRepoResponse{})

	assert.Nil(t, response)
	assert.NotNil(t, err)

}
