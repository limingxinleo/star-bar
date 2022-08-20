package repo

type Repo struct {
	FullName        string `json:"full_name"`
	StargazersCount uint64 `json:"stargazers_count"`
	ForksCount      uint64 `json:"forks_count"`
	OpenIssuesCount uint64 `json:"open_issues_count"`
}
