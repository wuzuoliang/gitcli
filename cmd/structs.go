package cmd

import "encoding/json"

const (
	privateToken = "private_token"
	perPage      = "per_page"
	perPageSize  = 20
	simple       = "true"
	owned        = "false"
	buckets      = "GitLab"
	boltDB       = "gitlab.bolt.db"
	groupType    = 1
	projectType  = 2
)

type GitLabGroupInfo struct {
	ID   int
	PID  int
	Name string `json:"path"`
	Path string `json:"full_path"`
}

type GitLabGroupList []GitLabGroupInfo

func (g GitLabGroupList) Len() int {
	return len(g)
}
func (g GitLabGroupList) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
func (g GitLabGroupList) Less(i, j int) bool {
	return g[i].Path < g[j].Path
}

type GitLabProjectInfo struct {
	ID   int
	PID  int
	Name string `json:"path"`
	Path string `json:"path_with_namespace"`
	SSH  string `json:"ssh_url_to_repo"`
}

type GitLabProjectList []GitLabProjectInfo

func (g GitLabProjectList) Len() int {
	return len(g)
}
func (g GitLabProjectList) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
func (g GitLabProjectList) Less(i, j int) bool {
	return g[i].Path < g[j].Path
}

type GitLabInfo struct {
	ID   int
	PID  int
	Name string
	Path string
	SSH  string
	Type int //   groupType = 1  projectType = 2

}

func (g GitLabInfo) Byte() []byte {
	buf, _ := json.Marshal(g)
	return buf
}

type GitLabInfoList []GitLabInfo

func (d GitLabInfoList) Len() int {
	return len(d)
}
func (d GitLabInfoList) Less(i, j int) bool {
	if d[i].Path < d[j].Path {
		return true
	}
	return false
}
func (d GitLabInfoList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
