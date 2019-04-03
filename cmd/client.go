package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type GitLabClient struct {
	// private access token
	accessToken string
	// host
	host string
}

func newClient(host, accessToken string) GitLabClient {
	return GitLabClient{
		accessToken: accessToken,
		host:        host,
	}
}

func (c *GitLabClient) GetProjects() ([]GitLabProjectInfo, error) {
	reqURL := c.host + "/api/v4/projects.json"
	page := 1
	list := make([]GitLabProjectInfo, 0, perPageSize)
	for {
		page++
		buff, err := c.doHTTPGet(reqURL, page)
		if err != nil {
			fmt.Println("GitLabClient::GetProjects doHTTPGet error", err)
			return []GitLabProjectInfo{}, err
		}
		fmt.Println(string(buff))

		ret := make([]GitLabProjectInfo, 0, perPageSize)
		err = json.Unmarshal(buff, &ret)
		if err != nil {
			fmt.Println("GitLabClient::GetProjects json.Unmarshal error", err)
			return []GitLabProjectInfo{}, err
		}
		list = append(list, ret...)
		if len(ret) < perPageSize {
			break
		}
	}
	return list, nil
}

func (c *GitLabClient) GetGroups() ([]GitLabGroupInfo, error) {
	reqURL := c.host + "/api/v4/groups.json"
	page := 1
	list := make([]GitLabGroupInfo, 0, perPageSize)
	for {
		page++
		buff, err := c.doHTTPGet(reqURL, page)
		if err != nil {
			fmt.Println("GitLabClient::GetGroups doHTTPGet error", err)
			return []GitLabGroupInfo{}, err
		}
		fmt.Println(string(buff))

		ret := make([]GitLabGroupInfo, 0, perPageSize)
		err = json.Unmarshal(buff, &ret)
		if err != nil {
			fmt.Println("GitLabClient::GetGroups json.Unmarshal error", err)
			return []GitLabGroupInfo{}, err
		}
		fmt.Println(ret)
		list = append(list, ret...)
		if len(ret) < perPageSize {
			break
		}
	}
	return list, nil
}

func (c *GitLabClient) doHTTPGet(destUrl string, page int) ([]byte, error) {
	u, _ := url.Parse(destUrl)
	q := u.Query()
	q.Set(privateToken, c.accessToken)
	q.Set("page", strconv.Itoa(page))
	q.Set(perPage, strconv.Itoa(perPageSize))
	q.Set("simple", simple)
	q.Set("owned", owned)
	u.RawQuery = q.Encode()
	fmt.Println(u.String())
	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			fmt.Println("GitLabClient::doHTTPGet close res error", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("request failed")
	}

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("GitLabClient::doHTTPGet ioutil.ReadAll error", err)
		return nil, err
	}
	return buff, nil
}
