package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"text/template"
	// "github.com/davecgh/go-spew/spew"
	// "os"
)

func setupListener() {

	httpURI := cfg.Section("http").Key("http_uri").String()
	httpPort := cfg.Section("http").Key("http_port").String()

	mux := http.NewServeMux()
	mux.HandleFunc(httpURI, PostHandler)

	log.Debugf("listening on port %s", httpPort)
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

		// Unmarshall data and get room to post to
		var postData GiteaPostData
		json.Unmarshal(body, &postData)
		args := strings.Split(r.URL.String(), "/")
		room := args[len(args)-1]

		// Check token
		if mygiteabot.checkToken(room, postData.Secret) {
			log.Debugf("Posting to room: %s", room)

			msgType := cfg.Section("bot").Key("message_type").String()

			//TODO check for configured message type

			messageText, err := generateTextMessage(postData, r.Header.Get("X-Gitea-Event"))
			if err != nil {
				log.Warning(err)
				log.Warning(postData)
			}
			switch msgType {
			case "html":
				messageHTML, err := generateHTMLMessage(postData, r.Header.Get("X-Gitea-Event"))
				if err != nil {
					log.Warning(err)
					log.Warning(postData)
				}
				mygiteabot.SendHTMLToRoom(room, messageHTML, messageText)
			case "plain":
				mygiteabot.SendTextToRoom(room, messageText)
			default:
				log.Fatalf("Wrong message type %s. Supported 'html' and 'plain'", msgType)
			}

		} else {
			log.Warningf("Wrong token %s for room: %s", postData.Secret, room)
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//generateTextMessage generates the message string for a given event
func generateHTMLMessage(data GiteaPostData, eventHeader string) (string, error) {

	templHeader := `<h3><a href="{{.Repository.HTMLURL}}"><b><span data-mx-color="#000000">[{{.Repository.FullName}}]</span></b></a></h3>`

	mxc, err := mygiteabot.Client.UploadLink(data.Sender.AvatarURL)

	if err != nil {
		log.Error("Cound not upload user avatar")
		log.Error(err)
		templHeader = templHeader + `<img src='mxc://matrix.org/GZtGmjBmfljRasPdjvWdKKVe' alt='[gitea]' width='32' height='32' title='gitea'/> <span data-mx-color="#609926">{{.Sender.FullName}}</span> `
	} else {
		templHeader = templHeader + "<img src='" + mxc.ContentURI + `' alt='[gitea]' width='30' height='30' title='image title'/> <span data-mx-color="#609926">{{.Sender.FullName}}</span> `
	}

	templ := template.New("notification")
	var tpl bytes.Buffer

	mesgTemplates := map[string]string{

		"push":                       "pushed " + strconv.Itoa(len(data.Commits)) + " commit(s) to {{.Repository.Name}}",
		"issues.assigned":            "assigned issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i> to  {{}}",
		"issues.closed":              "closed issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"issues.demilestoned":        "removed milestone TODO from issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"issues.edited":              "edited issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"issues.label_cleared":       "cleared labels from issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"issues.label_updated":       "updated labels of issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"issues.milestoned":          "added issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i> to milestone TODO",
		"issues.opened":              "opened issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"issues.reopened":            "re-opened issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"issues.synchronized":        "synchronized issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"issues.unassigned":          "removed assignee from issue <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"fork":                       "forked repository {{.Repository.Parent.FullName}} to {{.Repository.FullName}}",
		"pull_request.assigned":      "assigned pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i> to {{.PullRequest.Assignee.FullName}}",
		"pull_request.closed":        "closed pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request.demilestoned":  "removed milestone TODO from pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request.edited":        "edited pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request.label_cleared": "removed labels from pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request.label_updated": "updated labels from pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request.milestoned":    "added pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i> to milestone TODO",
		"pull_request.opened":        "opened pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request.reopened":      "re-opened pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request.synchronized":  "synchronized pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request.unassigned":    "removed assinee from pull-request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"issue_comment.created":      "commented on issue <b>#{{.Issue.Number}}</b> <i>{{.Issue.Title}}</i>:<br> <pre><code class='language-markdown'>{{.Comment.Body}}</code></pre>",
		"issue_comment.deleted":      "deleted commented on <b>#{{Issue.Number}}</b> <i>{{.Issue.Title}}</i>",
		"repository.created":         "created repository {{}}",
		"repository.deleted":         "deleted repository {{}}",
		"release.published":          "published release {{}}",
		"release.updated":            "updated release {{}}",
		"release.deleted":            "deleted release {{}}",
		"pull_request_aproved":       "approved pull request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request_rejected":      "rejected pull request  <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>",
		"pull_request_comment":       "commented on pull request <b>#{{PullRequest.Number}}</b> <i>#{{PullRequest.Title}}</i>: <pre><code class='language-markdown'>{{.Comment.Body}}</code><pre>",
	}

	switch eventHeader {
	case "push":
		templ.Parse(templHeader + mesgTemplates["push"])
	case "issues":
		templ.Parse(templHeader + mesgTemplates["issues."+data.Action])
	case "fork":
		templ.Parse(templHeader + mesgTemplates["fork"])
	case "pull_request":
		templ.Parse(templHeader + mesgTemplates["pull_request."+data.Action])
	case "issue_comment":
		templ.Parse(templHeader + mesgTemplates["issue_comment."+data.Action])
	case "repository":
		templ.Parse(templHeader + mesgTemplates["repository."+data.Action])
	case "release":
		templ.Parse(templHeader + mesgTemplates["release."+data.Action])
	case "pull_request_approved":
		templ.Parse(templHeader + mesgTemplates["pull_request_aproved"])
	case "pull_request_rejected":
		templ.Parse(templHeader + mesgTemplates["pull_request_rejected"])
	case "pull_request_comment":
		templ.Parse(templHeader + mesgTemplates["pull_request_comment"])
	default:
		log.Warningf("Unknown action: %s for eventHeader %s", data.Action, eventHeader)
		templ.Parse("Gitea did something unexpected, seriously wtf was that?! Event: " + eventHeader + " Action: " + data.Action)
	}

	if err := templ.Execute(&tpl, data); err != nil {
		log.Warningf("Failed to generate text-message for eventHeader %v", eventHeader)
		return "", err
	}

	return tpl.String(), nil
}

//generateTextMessage generates the message string for a given event
func generateTextMessage(data GiteaPostData, eventHeader string) (string, error) {

	templ := template.New("notification")
	var tpl bytes.Buffer

	mesgTemplates := map[string]string{
		"push":                       "[{{.Repository.Fullname}}] {{.Pusher.FullName}} pushed " + strconv.Itoa(len(data.Commits)) + " commit(s) to {{.Repository.Name}}",
		"issues.assigned":            "[{{.Repository.FullName}}]: {{.Sender.FullName}} assigned issue #{{.Issue.Number}} {{.Issue.Title}} to  {{}}",
		"issues.closed":              "[{{.Repository.FullName}}]: {{.Sender.FullName}} closed issue #{{.Issue.Number}} {{.Issue.Title}}",
		"issues.demilestoned":        "[{{.Repository.FullName}}]: {{.Sender.FullName}} removed milestone TODO from issue #{{.Issue.Number}} {{.Issue.Title}}",
		"issues.edited":              "[{{.Repository.FullName}}]: {{.Sender.FullName}} edited issue #{{.Issue.Number}} {{.Issue.Title}}",
		"issues.label_cleared":       "[{{.Repository.FullName}}]: {{.Sender.FullName}} cleared labels from issue #{{.Issue.Number}} {{.Issue.Title}}",
		"issues.label_updated":       "[{{.Repository.FullName}}]: {{.Sender.FullName}} updated labels of issue #{{.Issue.Number}} {{.Issue.Title}}",
		"issues.milestoned":          "[{{.Repository.FullName}}]: {{.Sender.FullName}} added issue #{{.Issue.Number}} {{.Issue.Title}} to milestone TODO",
		"issues.opened":              "[{{.Repository.FullName}}]: {{.Issue.User.FullName}} opened issue #{{.Issue.Number}} {{.Issue.Title}}",
		"issues.reopened":            "[{{.Repository.FullName}}]: {{.Sender.FullName}} re-opened issue #{{.Issue.Number}} {{.Issue.Title}}",
		"issues.synchronized":        "[{{.Repository.FullName}}]: {{.Sender.FullName}} synchronized issue #{{.Issue.Number}} {{.Issue.Title}}",
		"issues.unassigned":          "[{{.Repository.FullName}}]: {{.Sender.FullName}} removed assignee from issue #{{.Issue.Number}} {{.Issue.Title}}",
		"fork":                       "[{{.Repository.FullName}}]: {{.Sender.FullName}} forked repository {{.Repository.Parent.FullName}} to {{.Repository.FullName}}",
		"pull_request.assigned":      "[{{.Repository.FullName}}]: {{.Sender.FullName}} assigned pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\" to {{.PullRequest.Assignee.FullName}}",
		"pull_request.closed":        "[{{.Repository.FullName}}]: {{.Sender.FullName}} closed pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request.demilestoned":  "[{{.Repository.FullName}}]: {{.Sender.FullName}} removed milestone TODO from pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request.edited":        "[{{.Repository.FullName}}]: {{.Sender.FullName}} edited pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request.label_cleared": "[{{.Repository.FullName}}]: {{.Sender.FullName}} removed labels from pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request.label_updated": "[{{.Repository.FullName}}]: {{.Sender.FullName}} updated labels from pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request.milestoned":    "[{{.Repository.FullName}}]: {{.Sender.FullName}} added pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\" to milestone TODO",
		"pull_request.opened":        "[{{.Repository.FullName}}]: {{.Sender.FullName}} opened pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request.reopened":      "[{{.Repository.FullName}}]: {{.Sender.FullName}} re-opened pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request.synchronized":  "[{{.Repository.FullName}}]: {{.Sender.FullName}} synchronized pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request.unassigned":    "[{{.Repository.FullName}}]: {{.Sender.FullName}} removed assinee from pull-request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"issue_comment.created":      "{{.Sender.FullName}} commented on #{{.Issue.Number}} {{.Issue.Title}}: {{.Comment.Body}}",
		"issue_comment.deleted":      "{{.Sender.FullName}} deleted commented on #{{.Issue.Number}} {{.Issue.Title}}",
		"repository.created":         "{{.Sender.FullName}} created repository {{}}",
		"repository.deleted":         "{{.Sender.FullName}} deleted repository {{}}",
		"release.published":          "[{{.Repository.FullName}}]: {{.Sender.FullName}} published release {{}}",
		"release.updated":            "[{{.Repository.FullName}}]: {{.Sender.FullName}} updated release {{}}",
		"release.deleted":            "[{{.Repository.FullName}}]: {{.Sender.FullName}} deleted release {{}}",
		"pull_request_aproved":       "[{{.Repository.FullName}}]: {{.Sender.FullName}} approved pull request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request_rejected":      "[{{.Repository.FullName}}]: {{.Sender.FullName}} rejected pull request  #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\"",
		"pull_request_comment":       "[{{.Repository.FullName}}]: {{.Sender.FullName}} commented on pull request #{{.PullRequest.Number}} \"{{.PullRequest.Title}}\": {{.Comment.Body}}",
	}

	switch eventHeader {
	case "push":
		templ.Parse(mesgTemplates["push"])
	case "issues":
		templ.Parse(mesgTemplates["issues."+data.Action])
	case "fork":
		templ.Parse(mesgTemplates["fork"])
	case "pull_request":
		templ.Parse(mesgTemplates["pull_request."+data.Action])
	case "issue_comment":
		templ.Parse(mesgTemplates["issue_comment."+data.Action])
	case "repository":
		templ.Parse(mesgTemplates["repository."+data.Action])
	case "release":
		templ.Parse(mesgTemplates["release."+data.Action])
	case "pull_request_approved":
		templ.Parse(mesgTemplates["pull_request_aproved"])
	case "pull_request_rejected":
		templ.Parse(mesgTemplates["pull_request_rejected"])
	case "pull_request_comment":
		templ.Parse(mesgTemplates["pull_request_comment"])
	default:
		log.Warningf("Unknown action: %s for eventHeader %s", data.Action, eventHeader)
		templ.Parse("Gitea did something unexpected, seriously wtf was that?! Event: " + eventHeader + " Action: " + data.Action)
	}

	if err := templ.Execute(&tpl, data); err != nil {
		log.Warningf("Failed to generate text-message for eventHeader %v", eventHeader)
		return "", err
	}

	return tpl.String(), nil
}
