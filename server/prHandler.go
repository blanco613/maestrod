package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/cpg1111/maestrod/lifecycle"
)

type PRSubHandler struct {
	SubHandler
	Error ErrorHandler
	Queue *lifecycle.Queue
}

func NewPRHandler(queue *lifecycle.Queue) *RouteHandler {
	sub := PRSubHandler{
		Error: ErrorHandler{},
		Queue: queue,
	}
	return &RouteHandler{
		Route: "/pull_request",
		sub:   sub,
	}

}

//type authorPayload struct {	##example json
//	Name     string `json:"name"`
//	Email    string `json:"email"`
//	UserName string `json:"username"`


type pull_request struct {
        ID	uint	`json:"id"`
	State	string	`json:"state"`
	Locked	bool	`json:"locked"`	    //is this correct?   <--------- 
	Title	string  `json:"title"`

	type user struct {				//is this how this should be nested?					
		Login	string	`json:"login"`
		ID	uint	`json:"login"`
		Type	string	`json:"user"`
		Site_Admin bool	`json:"site_admin"`		 
			 }

}
/* still to be added

  "pull_request": {

    "id": 34778301,
    "state": "open",
    "locked": false,
    "title": "Update the README with new information",
    

   
 "body": "This is a pretty simple change that we need to pull into master.",
    "created_at": "2015-05-05T23:40:27Z",
    "updated_at": "2015-05-05T23:40:27Z",
    "closed_at": null,
    "merged_at": null,
    "merge_commit_sha": null,
    "assignee": null,
    "milestone": null,
   
   "head": {
      "label": "baxterthehacker:changes",
      "ref": "changes",
      "sha": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
      "user": {
        "login": "baxterthehacker",
        "id": 6752317,
        "gravatar_id": "",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 35129377,
        "name": "public-repo",
        "full_name": "baxterthehacker/public-repo",
        "owner": {
          "login": "baxterthehacker",
          "id": 6752317,
          "gravatar_id": "",
          "type": "User",
          "site_admin": false
        },
        "private": false,
        "description": "",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": true,
        "forks_count": 0,
        "mirror_url": null,
        "open_issues_count": 1,
        "forks": 0,
        "open_issues": 1,
        "watchers": 0,
        "default_branch": "master"
      }
    },
    "base": {
      "label": "baxterthehacker:master",
      "ref": "master",
      "sha": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
      "user": {
        "login": "baxterthehacker",
        "id": 6752317,
        "gravatar_id": "",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 35129377,
        "name": "public-repo",
        "full_name": "baxterthehacker/public-repo",
        "owner": {
          "login": "baxterthehacker",
          "id": 6752317,
          "gravatar_id": "",
          "type": "User",
          "site_admin": false
        },
        "private": false,
        "description": "",
        "fork": false,
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": true,
        "forks_count": 0,
        "mirror_url": null,
        "open_issues_count": 1,
        "forks": 0,
        "open_issues": 1,
        "watchers": 0,
        "default_branch": "master"
      }
     },

      }
    },
    "merged": false,
    "mergeable": null,
    "mergeable_state": "unknown",
    "merged_by": null,
    "comments": 0,
    "review_comments": 0,
    "commits": 1,
    "additions": 1,
    "deletions": 1,
    "changed_files": 1
  },
  "repository": {
    "id": 35129377,
    "name": "public-repo",
    "full_name": "baxterthehacker/public-repo",
    "owner": {
      "login": "baxterthehacker",
      "id": 6752317,
      "gravatar_id": "",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "description": "",
    "fork": false,
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": true,
    "forks_count": 0,
    "mirror_url": null,
    "open_issues_count": 1,
    "forks": 0,
    "open_issues": 1,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "baxterthehacker",
    "id": 6752317,
    "gravatar_id": "",
    "type": "User",
    "site_admin": false
  }
}
*/
