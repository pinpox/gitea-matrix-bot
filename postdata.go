package main

import (
	"time"
)

type GiteaUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	FullName  string `json:"full_name"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Language  string `json:"language"`
	Username  string `json:"username"`
}

type GiteaRepository struct {
	ID          int       `json:"id"`
	Owner       GiteaUser `json:"owner"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	Description string    `json:"description"`
	Empty       bool      `json:"empty"`
	Private     bool      `json:"private"`
	Fork        bool      `json:"fork"`
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
		Owner           GiteaUser `json:"owner"`
		Permissions     struct {
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
}

type GiteaRelease struct {
	ID              int           `json:"id"`
	TagName         string        `json:"tag_name"`
	TargetCommitish string        `json:"target_commitish"`
	Name            string        `json:"name"`
	Body            string        `json:"body"`
	URL             string        `json:"url"`
	TarballURL      string        `json:"tarball_url"`
	ZipballURL      string        `json:"zipball_url"`
	Draft           bool          `json:"draft"`
	Prerelease      bool          `json:"prerelease"`
	CreatedAt       time.Time     `json:"created_at"`
	PublishedAt     time.Time     `json:"published_at"`
	Author          GiteaUser     `json:"author"`
	Assets          []interface{} `json:"assets"`
}

type GiteaPullRequest struct {
	ID             int            `json:"id"`
	URL            string         `json:"url"`
	Number         int            `json:"number"`
	User           GiteaUser      `json:"user"`
	Title          string         `json:"title"`
	Body           string         `json:"body"`
	Labels         []interface{}  `json:"labels"`
	Milestone      GiteaMilestone `json:"milestone"`
	Assignee       GiteaUser      `json:"assignee"`
	Assignees      []GiteaUser    `json:"assignees"`
	State          string         `json:"state"`
	Comments       int            `json:"comments"`
	HTMLURL        string         `json:"html_url"`
	DiffURL        string         `json:"diff_url"`
	PatchURL       string         `json:"patch_url"`
	Mergeable      bool           `json:"mergeable"`
	Merged         bool           `json:"merged"`
	MergedAt       time.Time      `json:"merged_at"`
	MergeCommitSha interface{}    `json:"merge_commit_sha"`
	MergedBy       interface{}    `json:"merged_by"`
	Base           struct {
		Label  string          `json:"label"`
		Ref    string          `json:"ref"`
		Sha    string          `json:"sha"`
		RepoID int             `json:"repo_id"`
		Repo   GiteaRepository `json:"repo"`
	} `json:"base"`
	Head struct {
		Label  string          `json:"label"`
		Ref    string          `json:"ref"`
		Sha    string          `json:"sha"`
		RepoID int             `json:"repo_id"`
		Repo   GiteaRepository `json:"repo"`
	} `json:"head"`
	MergeBase string    `json:"merge_base"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedAt  time.Time `json:"closed_at"`
}

type GiteaMilestone struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	State        string    `json:"state"`
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	ClosedAt     time.Time `json:"closed_at"`
	DueOn        time.Time `json:"due_on"`
}

type GiteaIssue struct {
	ID          int            `json:"id"`
	URL         string         `json:"url"`
	Number      int            `json:"number"`
	User        GiteaUser      `json:"user"`
	Title       string         `json:"title"`
	Body        string         `json:"body"`
	Labels      []interface{}  `json:"labels"`
	Milestone   GiteaMilestone `json:"milestone"`
	Assignee    GiteaUser      `json:"assignee"`
	Assignees   []GiteaUser    `json:"assignees"`
	State       string         `json:"state"`
	Comments    int            `json:"comments"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	ClosedAt    time.Time      `json:"closed_at"`
	DueDate     time.Time      `json:"due_date"`
	PullRequest time.Time      `json:"pull_request"`
}

type GiteaCommit struct {
	ID           string      `json:"id"`
	Message      string      `json:"message"`
	URL          string      `json:"url"`
	Author       GiteaUser   `json:"author"`
	Committer    GiteaUser   `json:"committer"`
	Verification interface{} `json:"verification"`
	Timestamp    time.Time   `json:"timestamp"`
}

type GiteaComment struct {
	ID             int       `json:"id"`
	HTMLURL        string    `json:"html_url"`
	PullRequestURL string    `json:"pull_request_url"`
	IssueURL       string    `json:"issue_url"`
	User           GiteaUser `json:"user"`
	Body           string    `json:"body"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
type GiteaPostData struct {
	Secret      string           `json:"secret"`
	Action      string           `json:"action"`
	Number      int              `json:"number"`
	Ref         string           `json:"ref"`
	Before      string           `json:"before"`
	After       string           `json:"after"`
	CompareURL  string           `json:"compare_url"`
	Release     GiteaRelease     `json:"release"`
	PullRequest GiteaPullRequest `json:"pull_request"`
	Issue       GiteaIssue       `json:"issue"`
	Commits     []GiteaCommit    `json:"commits"`
	Comment     GiteaComment     `json:"comment"`
	Repository  GiteaRepository  `json:"repository"`
	Pusher      GiteaUser        `json:"pusher"`
	Sender      GiteaUser        `json:"sender"`
}
