package models

type GithubEventResponse struct {
	Type string `json:"type"`
	Repo struct {
		Url string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits,omitempty"`
		Action string `json:"action,omitempty"`
	} `json:"payload"`
	CreatedAt string `json:"created_at"`
}
