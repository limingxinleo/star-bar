package repo

type Repo struct {
	FullName        string `json:"full_name"`
	StargazersCount int64  `json:"stargazers_count"`
	ForksCount      int64  `json:"forks_count"`
}