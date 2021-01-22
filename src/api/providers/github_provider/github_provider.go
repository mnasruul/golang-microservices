package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/mnasruul/golang-microservices/src/api/clients/resclient"
	"github.com/mnasruul/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	// Authorization: token ffabaf2e2b5f9b982b37c787dbfce0ed4ba96e36
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := resclient.Post(urlCreateRepo, request, headers)
	fmt.Println(response)
	fmt.Println(err)
	if err != nil {
		log.Println(fmt.Sprintf("error where trying to create new repo ingithub: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}
	defer response.Body.Close()
	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid  json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, result); err != nil {
		log.Println(fmt.Sprintf("error where trying to unmarshall create repo successfully: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error where trying to unmarshalling github create repo response",
		}
	}
	return &result, nil
}
