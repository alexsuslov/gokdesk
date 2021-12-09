package gokdesk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexsuslov/gokdesk/model"
	"io/ioutil"
	"os"
)

var OKDESK_ISSUE_COMMENTS = "%s/api/v1/issues/%v/comments"

func GetCommentsByID(ctx context.Context, issueID string) (comments []model.CommentType, err error) {
	Url := fmt.Sprintf(OKDESK_ISSUE_COMMENTS, os.Getenv("OKDESK_URL"), issueID)
	if err != nil {
		return
	}
	body, err := Request(ctx, "GET", Url, nil, nil)
	if err != nil {
		return
	}
	defer body.Close()

	comments = []model.CommentType{}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	return comments, json.Unmarshal(data, &comments)
}
