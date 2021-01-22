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

func CreateRepo(accesToken string, request github.CreateRepoResponse) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	// Authorization: token ffabaf2e2b5f9b982b37c787dbfce0ed4ba96e36
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accesToken))

	respone, err := resclient.Post(urlCreateRepo, request, headers)
	fmt.Println(respone)
	fmt.Println(err)
	if err != nil {
		log.Println(fmt.Sprintf("error where trying to create new repo ingithub: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(respone.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid respone body",
		}
	}
	defer respone.Body.Close()
	if respone.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid  json respone body",
			}
		}
		errResponse.StatusCode = respone.StatusCode
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
