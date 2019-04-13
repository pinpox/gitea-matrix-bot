package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	// "os"
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

		//TODO check secret!
		//TODO check repo

		json.Unmarshal(body, &postData)

		// fmt.Println(string(body))

		fmt.Println("=================================================")
		fmt.Println(generateMessage(postData, r.Header.Get("X-Gitea-Event")))
		fmt.Println("=================================================")

		fmt.Printf("%+v\n", postData)
		fmt.Println("=================================================")
		spew.Dump(postData)
		fmt.Println("=================================================")

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//generateMessage generates the message string for a given event
func generateMessage(data GiteaPostData, eventHeader string) string {

	templ := template.New("notification")
	var tpl bytes.Buffer

	switch eventHeader {

	case "push":
		templ.Parse("{{.Pusher.FullName}} pushed " + strconv.Itoa(len(data.Commits)) + " commit(s) to {{.Repository.Name}}")

	case "issues":
		switch data.Action {
		case "assigned":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} assigned issue {{}} {{}} to {{}}")
		case "closed":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} closed issue {{}} {{}}")
		case "demilestoned":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} removed milestone {{}} from issue {{}} {{}}")
		case "edited":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} edited issue {{}} {{}}")
		case "label_cleared":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} cleared labels from issue {{}} {{}}")
		case "label_updated":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} updated labels of issue {{}} {{}}")
		case "milestoned":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} added issue {{}} {{}} to milestone {{}}")
		case "opened":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} opened issue {{}} {{}}")
		case "reopened":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} re-opened issue {{}} {{}}")
		case "synchronized":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} synchronized issue {{}} {{}}")
		case "unassigned":
			templ.Parse("{{.Repository.FullName}}: {{.Sender.Fullname}} removed assignee from issue {{}} {{}}")
		default:
			fmt.Println("Unknown action: " + data.Action)
		}

	case "fork":
		fmt.Println("Event: Fork")
		switch data.Action {
		case "created":
			templ.Parse("")
		case "edited":
			templ.Parse("")
		case "deleted":
			templ.Parse("")
		default:
			fmt.Println("Unknown action: " + data.Action)

		}

	case "pull_request":
		fmt.Println("Event: Pull Request")
		switch data.Action {
		case "assigned":
			templ.Parse("")
		case "closed":
			templ.Parse("")
		case "demilestoned":
			templ.Parse("")
		case "edited":
			templ.Parse("")
		case "label_cleared":
			templ.Parse("")
		case "label_updated":
			templ.Parse("")
		case "milestoned":
			templ.Parse("")
		case "opened":
			templ.Parse("")
		case "reopened":
			templ.Parse("")
		case "synchronized":
			templ.Parse("")
		case "unassigned":
			templ.Parse("")
		default:
			fmt.Println("Unknown action: " + data.Action)

		}

	case "issue_comment":
		fmt.Println("Event: Issues comment")
		templ.Parse("")

	case "repository":
		fmt.Println("Event: Repository")
		switch data.Action {
		case "created":
			templ.Parse("")
		case "deleted":
			templ.Parse("")
		default:
			fmt.Println("Unknown action: " + data.Action)
		}

	case "release":
		fmt.Println("Event: Release")
		switch data.Action {
		case "published":
			templ.Parse("")
		case "updated":
			templ.Parse("")
		case "deleted":
			templ.Parse("")

		default:
			fmt.Println("Unknown action: " + data.Action)

		}

	case "pull_request_approved":
		fmt.Println("Event: PR approoved")
		fmt.Println("Unknown action: " + data.Action)
		templ.Parse("")
	case "pull_request_rejected":
		fmt.Println("Event: PR reject")
		fmt.Println("Unknown action: " + data.Action)
		templ.Parse("")
	case "pull_request_comment":
		fmt.Println("Event: PR comment")
		fmt.Println("Unknown action: " + data.Action)
		templ.Parse("")
	default:
		fmt.Println("Event: Unknown")
		fmt.Println(eventHeader)
		fmt.Println("Unknown action: " + data.Action)
	}

	if err := templ.Execute(&tpl, data); err != nil {
		panic(err)
	}

	return tpl.String()

	// tmplIssue, err := template.New("test").Parse("{{.User}} {{.Action}} Issue {{.IssueID}} in repository {{.Repo}}")

}
