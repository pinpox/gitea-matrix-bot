package main

import (
	"time"
)

type GiteaPostData struct {
	Secret     string `json:"secret"`
	Action     string `json:"action"`
	Number     int    `json:"number"`
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	CompareURL string `json:"compare_url"`
	Issue      struct {
		ID     int    `json:"id"`
		URL    string `json:"url"`
		Number int    `json:"number"`
		User   struct {
			ID        int    `json:"id"`
			Login     string `json:"login"`
			FullName  string `json:"full_name"`
			Email     string `json:"email"`
			AvatarURL string `json:"avatar_url"`
			Language  string `json:"language"`
			Username  string `json:"username"`
		} `json:"user"`
		Title       string        `json:"title"`
		Body        string        `json:"body"`
		Labels      []interface{} `json:"labels"`
		Milestone   interface{}   `json:"milestone"`
		Assignee    interface{}   `json:"assignee"`
		Assignees   interface{}   `json:"assignees"`
		State       string        `json:"state"`
		Comments    int           `json:"comments"`
		CreatedAt   time.Time     `json:"created_at"`
		UpdatedAt   time.Time     `json:"updated_at"`
		ClosedAt    interface{}   `json:"closed_at"`
		DueDate     interface{}   `json:"due_date"`
		PullRequest interface{}   `json:"pull_request"`
	} `json:"issue"`
	Commits []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
		URL     string `json:"url"`
		Author  struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Verification interface{} `json:"verification"`
		Timestamp    time.Time   `json:"timestamp"`
	} `json:"commits"`
	Comment struct {
		ID             int    `json:"id"`
		HTMLURL        string `json:"html_url"`
		PullRequestURL string `json:"pull_request_url"`
		IssueURL       string `json:"issue_url"`
		User           struct {
			ID        int    `json:"id"`
			Login     string `json:"login"`
			FullName  string `json:"full_name"`
			Email     string `json:"email"`
			AvatarURL string `json:"avatar_url"`
			Language  string `json:"language"`
			Username  string `json:"username"`
		} `json:"user"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"comment"`
	Repository struct {
		ID    int `json:"id"`
		Owner struct {
			ID        int    `json:"id"`
			Login     string `json:"login"`
			FullName  string `json:"full_name"`
			Email     string `json:"email"`
			AvatarURL string `json:"avatar_url"`
			Language  string `json:"language"`
			Username  string `json:"username"`
		} `json:"owner"`
		Name        string `json:"name"`
		FullName    string `json:"full_name"`
		Description string `json:"description"`
		Empty       bool   `json:"empty"`
		Private     bool   `json:"private"`
		Fork        bool   `json:"fork"`
		Parent      struct {
			CreatedAt       time.Time `json:"created_at"`
			Empty           bool      `json:"empty"`
			Fork            bool      `json:"fork"`
			Private         bool      `json:"private"`
			Archived        bool      `json:"archived"`
			FullName        string    `json:"full_name"`
			ForksCount      int       `json:"forks_count"`
			ID              int       `json:"id"`
			OpenIssuesCount int       `json:"open_issues_count"`
			Size            int       `json:"size"`
			StarsCount      int       `json:"stars_count"`
			Name            string    `json:"name"`
			CloneURL        string    `json:"clone_url"`
			DefaultBranch   string    `json:"default_branch"`
			SSHURL          string    `json:"ssh_url"`
			WatchersCount   int       `json:"watchers_count"`
			Description     string    `json:"description"`
			UpdatedAt       time.Time `json:"updated_at"`
			Website         string    `json:"website"`
			HTMLURL         string    `json:"html_url"`
			Owner           struct {
				ID        int    `json:"id"`
				Login     string `json:"login"`
				FullName  string `json:"full_name"`
				Email     string `json:"email"`
				AvatarURL string `json:"avatar_url"`
				Language  string `json:"language"`
				Username  string `json:"username"`
			} `json:"owner"`
			Permissions struct {
				Admin bool `json:"admin"`
				Push  bool `json:"push"`
				Pull  bool `json:"pull"`
			} `json:"permissions"`
		} `json:"parent"`
		Mirror          bool      `json:"mirror"`
		Size            int       `json:"size"`
		HTMLURL         string    `json:"html_url"`
		SSHURL          string    `json:"ssh_url"`
		CloneURL        string    `json:"clone_url"`
		Website         string    `json:"website"`
		StarsCount      int       `json:"stars_count"`
		ForksCount      int       `json:"forks_count"`
		WatchersCount   int       `json:"watchers_count"`
		OpenIssuesCount int       `json:"open_issues_count"`
		DefaultBranch   string    `json:"default_branch"`
		Archived        bool      `json:"archived"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		Permissions     struct {
			Admin bool `json:"admin"`
			Push  bool `json:"push"`
			Pull  bool `json:"pull"`
		} `json:"permissions"`
	} `json:"repository"`
	Pusher struct {
		ID        int    `json:"id"`
		Login     string `json:"login"`
		FullName  string `json:"full_name"`
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
		Language  string `json:"language"`
		Username  string `json:"username"`
	} `json:"pusher"`
	Sender struct {
		ID        int    `json:"id"`
		Login     string `json:"login"`
		FullName  string `json:"full_name"`
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
		Language  string `json:"language"`
		Username  string `json:"username"`
	} `json:"sender"`
}
