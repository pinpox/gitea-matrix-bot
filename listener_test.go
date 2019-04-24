package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_setupListener(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupListener()
		})
	}
}

func TestPostHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_generateHTMLMessage(t *testing.T) {
	tests := []struct {
		name        string
		eventHeader string
		dataPathIn  string
		dataPathOut string
		wantErr     bool
	}{
		{"Push", "push", "./testdata/gitea_postdata/html/push", "./testdata/messages/html/push", false},
		{"Issues_assigned", "issues", "./testdata/gitea_postdata/html/issues_assigned", "./testdata/messages/html/issues_assigned", false},
		{"Issues_closed", "issues", "./testdata/gitea_postdata/html/issues_closed", "./testdata/messages/html/issues_closed", false},
		{"Issues_demilestoned", "issues", "./testdata/gitea_postdata/html/issues_demilestoned", "./testdata/messages/html/issues_demilestoned", false},
		{"Issues_edited", "issues", "./testdata/gitea_postdata/html/issues_edited", "./testdata/messages/html/issues_edited", false},
		{"Issues_label_cleared", "issues", "./testdata/gitea_postdata/html/issues_label_cleared", "./testdata/messages/html/issues_label_cleared", false},
		{"Issues_label_updated", "issues", "./testdata/gitea_postdata/html/issues_label_updated", "./testdata/messages/html/issues_label_updated", false},
		{"Issues_milestoned", "issues", "./testdata/gitea_postdata/html/issues_milestoned", "./testdata/messages/html/issues_milestoned", false},
		{"Issues_opened", "issues", "./testdata/gitea_postdata/html/issues_opened", "./testdata/messages/html/issues_opened", false},
		{"Issues_reopened", "issues", "./testdata/gitea_postdata/html/issues_reopened", "./testdata/messages/html/issues_reopened", false},
		{"Issues_synchronized", "issues", "./testdata/gitea_postdata/html/issues_synchronized", "./testdata/messages/html/issues_synchronized", false},
		{"Issues_unassigned", "issues", "./testdata/gitea_postdata/html/issues_unassigned", "./testdata/messages/html/issues_unassigned", false},
		{"Fork", "fork", "./testdata/gitea_postdata/html/fork", "./testdata/messages/html/fork", false},
		{"Pull_request_assigned", "pull_request", "./testdata/gitea_postdata/html/pull_request_assigned", "./testdata/messages/html/pull_request_assigned", false},
		{"Pull_request_closed", "pull_request", "./testdata/gitea_postdata/html/pull_request_closed", "./testdata/messages/html/pull_request_closed", false},
		{"Pull_request_demilestoned", "pull_request", "./testdata/gitea_postdata/html/pull_request_demilestoned", "./testdata/messages/html/pull_request_demilestoned", false},
		{"Pull_request_edited", "pull_request", "./testdata/gitea_postdata/html/pull_request_edited", "./testdata/messages/html/pull_request_edited", false},
		{"Pull_request_label_cleared", "pull_request", "./testdata/gitea_postdata/html/pull_request_label_cleared", "./testdata/messages/html/pull_request_label_cleared", false},
		{"Pull_request_label_updated", "pull_request", "./testdata/gitea_postdata/html/pull_request_label_updated", "./testdata/messages/html/pull_request_label_updated", false},
		{"Pull_request_milestoned", "pull_request", "./testdata/gitea_postdata/html/pull_request_milestoned", "./testdata/messages/html/pull_request_milestoned", false},
		{"Pull_request_opened", "pull_request", "./testdata/gitea_postdata/html/pull_request_opened", "./testdata/messages/html/pull_request_opened", false},
		{"Pull_request_reopened", "pull_request", "./testdata/gitea_postdata/html/pull_request_reopened", "./testdata/messages/html/pull_request_reopened", false},
		{"Pull_request_synchronized", "pull_request", "./testdata/gitea_postdata/html/pull_request_synchronized", "./testdata/messages/html/pull_request_synchronized", false},
		{"Pull_request_unassigned", "pull_request", "./testdata/gitea_postdata/html/pull_request_unassigned", "./testdata/messages/html/pull_request_unassigned", false},
		{"Issue_comment_created", "issue_comment", "./testdata/gitea_postdata/html/issue_comment_created", "./testdata/messages/html/issue_comment_created", false},
		{"Issue_comment_deleted", "issue_comment", "./testdata/gitea_postdata/html/issue_comment_deleted", "./testdata/messages/html/issue_comment_deleted", false},
		{"Repository_created", "repository", "./testdata/gitea_postdata/html/repository_created", "./testdata/messages/html/repository_created", false},
		{"Repository_deleted", "repository", "./testdata/gitea_postdata/html/repository_deleted", "./testdata/messages/html/repository_deleted", false},
		{"Release_published", "release", "./testdata/gitea_postdata/html/release_published", "./testdata/messages/html/release_published", false},
		{"Release_updated", "release", "./testdata/gitea_postdata/html/release_updated", "./testdata/messages/html/release_updated", false},
		{"Release_deleted", "release", "./testdata/gitea_postdata/html/release_deleted", "./testdata/messages/html/release_deleted", false},
		{"Pull_request_approved", "pull_request", "./testdata/gitea_postdata/html/pull_request_approved", "./testdata/messages/html/pull_request_approved", false},
		{"Pull_request_rejected", "pull_request", "./testdata/gitea_postdata/html/pull_request_rejected", "./testdata/messages/html/pull_request_rejected", false},
		{"Pull_request_comment", "pull_request", "./testdata/gitea_postdata/html/pull_request_comment", "./testdata/messages/html/pull_request_comment", false},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			//Read POST data from file
			input, err := ioutil.ReadFile(tt.dataPathIn)
			if err != nil {
				panic(err)
			}

			//Read expected result from file
			output, err := ioutil.ReadFile(tt.dataPathOut)
			if err != nil {
				panic(err)
			}

			//Create a data structs
			var postData GiteaPostData
			json.Unmarshal(input, &postData)

			got, err := generateHTMLMessage(postData, tt.eventHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateHTMLMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != string(output) {
				t.Errorf("generateHTMLMessage() = %v, want %v", got, string(output))
			}
		})
	}
}

func Test_generateTextMessage(t *testing.T) {
	type args struct {
		data        GiteaPostData
		eventHeader string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateTextMessage(tt.args.data, tt.args.eventHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateTextMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateTextMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
