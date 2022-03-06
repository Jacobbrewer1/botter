package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Jacobbrewer1/botter/config"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func CreateIssue(title, body string) (Issue, User, error) {
	randomUser, err := getRandomRepoCollaborator("botter")
	if err != nil {
		return Issue{ID: nil}, User{ID: nil}, err
	}
	u, err := GetUser(*randomUser.Login)
	if err != nil {
		return Issue{}, User{}, err
	}
	rawJson, err := postBotterIssue(title, body, *u.Login)
	if err != nil {
		return Issue{ID: nil}, User{ID: nil}, err
	}
	i, err := decodeSingleIssue(rawJson)
	if err != nil {
		return Issue{ID: nil}, User{ID: nil}, nil
	}
	return i, u, nil
}

func GetUser(logon string) (User, error) {
	rawJson, err := requestGithub(fmt.Sprintf("users/%v", logon))
	if err != nil {
		return User{ID: nil}, err
	}
	return decodeSingleUser(rawJson)
}

func GetRepoCollaborators(repo string) ([]User, error) {
	rawJson, err := requestGithub(fmt.Sprintf("repos/Jacobbrewer1/%v/collaborators", repo))
	if err != nil {
		return nil, err
	}
	return decodeMultiUser(rawJson)
}

func GetBranches(repo string) ([]Branch, error) {
	rawJson, err := requestGithub(fmt.Sprintf("repos/Jacobbrewer1/%v/branches", repo))
	if err != nil {
		return nil, err
	}
	return decodeMultiBranch(rawJson)
}

func decodeSingleBranch(rawJson json.RawMessage) (Branch, error) {
	var b Branch
	err := json.Unmarshal(rawJson, &b)
	return b, err
}

func decodeMultiBranch(rawJson json.RawMessage) ([]Branch, error) {
	var b []Branch
	err := json.Unmarshal(rawJson, &b)
	return b, err
}

func getRandomRepoCollaborator(repo string) (User, error) {
	rawJson, err := requestGithub(fmt.Sprintf("repos/Jacobbrewer1/%v/collaborators", repo))
	if err != nil {
		return User{ID: nil}, err
	}
	users, err := decodeMultiUser(rawJson)
	return users[rand.Intn(len(users))], err
}

func decodeSingleUser(rawJson json.RawMessage) (User, error) {
	var u User
	err := json.Unmarshal(rawJson, &u)
	return u, err
}

func decodeMultiUser(rawJson json.RawMessage) ([]User, error) {
	var u []User
	err := json.Unmarshal(rawJson, &u)
	return u, err
}

func GetBotterIssues() ([]Issue, error) {
	jsonRaw, err := requestGithub("repos/Jacobbrewer1/botter/issues")
	if err != nil {
		return nil, err
	}
	i, err := decodeMultiIssues(jsonRaw)
	if err != nil {
		return nil, err
	}
	return removePullRequests(i)
}

func removePullRequests(i []Issue) ([]Issue, error) {
	if len(i) == 1 && i[0].IsPullRequest() {
		log.Println("Only issue found is a pr")
		return nil, nil
	}
	for z, x := range i {
		if x.IsPullRequest() {
			//log.Println("pr found in issues array")
			if z >= len(i) {
				break
			}
			i = removeIssueSlice(i, z)
		}
	}
	return i, nil
}

func decodeMultiIssues(jsonRaw json.RawMessage) ([]Issue, error) {
	var issues []Issue
	err := json.Unmarshal(jsonRaw, &issues)
	return issues, err
}

func decodeSingleIssue(jsonRaw json.RawMessage) (Issue, error) {
	var i Issue
	err := json.Unmarshal(jsonRaw, &i)
	return i, err
}

func removeIssueSlice(i []Issue, x int) []Issue {
	if x < 0 {
		return i
	}
	if i[len(i)-1].IsPullRequest() {
		for j := len(i) - 2; j >= x; {
			if !i[j].IsPullRequest() {
				i[x] = i[j]
				return i[:j]
			}
			j--
		}
		return i[:x]
	}
	i[x] = i[len(i)-1]
	return i[:len(i)-1]
}

func GetRepos() ([]Repository, error) {
	jsonRaw, err := requestGithub("user/repos")
	if err != nil {
		return nil, err
	}
	return decodeRepos(jsonRaw)
}

func decodeRepos(jsonRaw json.RawMessage) ([]Repository, error) {
	var repositories []Repository
	err := json.Unmarshal(jsonRaw, &repositories)
	return repositories, err
}

func postBotterIssue(title, bodyString, assignee string) (json.RawMessage, error) {
	issueData := NewIssue{
		Title:    title,
		Body:     bodyString,
		Assignee: assignee,
	}
	body, err := json.Marshal(issueData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.github.com/repos/Jacobbrewer1/botter/issues", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %v", *config.ApiSecrets.GithubApiToken))
	req.Header.Set("accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func requestGithub(endpoint string) (json.RawMessage, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/%v", endpoint), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %v", *config.ApiSecrets.GithubApiToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
