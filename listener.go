package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

func setupListener() {

	httpURI := cfg.Section("http").Key("http_uri").String()
	httpPort := cfg.Section("http").Key("http_port").String()

	mux := http.NewServeMux()
	mux.HandleFunc(httpURI, PostHandler)

	log.Printf("listening on port %s", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, mux))

}

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		var postData GiteaPostData

		json.Unmarshal(body, &postData)

		// fmt.Println(postData)
		generateMessage(postData)

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

type GiteaEvent int

const (
	ISSUE_OPEN   = 0
	ISSUE_CLOSE  = 1
	ISSUE_REOPEN = 2

	COMMENT_ADD  = 3
	COMMENT_EDIT = 4
	COMMENT_DEL  = 5

	PUSH = 6
)

//determineAction figures out what the hell the user did
func determineAction(data GiteaPostData) GiteaEvent {
	return 7
}

//generateMessage generates the message string for a given event
func generateMessage(data GiteaPostData) {

	switch determineAction(data) {
	case ISSUE_OPEN:
		fmt.Println("Issue Opened")
	case ISSUE_CLOSE:
		fmt.Println("Issue Closed")
	case ISSUE_REOPEN:
		fmt.Println("Issue Reopened")
	case COMMENT_ADD:
		fmt.Println("Comment added")
	case COMMENT_DEL:
		fmt.Println("Comment deleted")
	case COMMENT_EDIT:
		fmt.Println("Comment edited")
	default:
		fmt.Println("Unknown action")
		fmt.Println(data)
	}

	//TODO Events:
	//Pushed x commits
	// tmplPush, err := template.New("test").Parse("{{.User}} pushed {{.NumCommits}} to {{.Repo}}")

	//Actions: Opened, Closed, Repoened, Commented
	// tmplIssue, err := template.New("test").Parse("{{.User}} {{.Action}} Issue {{.IssueID}} in repository {{.Repo}}")
	//Issue Closed
	//Issue Opened
	//Issue Commented

	//PR Stuff

}
