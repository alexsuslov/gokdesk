package gokdesk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexsuslov/gokdesk/model"
	"io"
	"os"
)

var OKDESK_ISSUE = "%s/api/v1/issues/%v"

func Getter(ctx context.Context, issueID string) (body io.ReadCloser, err error) {

	Url := fmt.Sprintf(OKDESK_ISSUE, os.Getenv("OKDESK_URL"), issueID)

	return Request(ctx, "GET", Url, nil, nil)
}

func GetByID(ctx context.Context, issueID string) (issue *model.Issue, err error) {

	body, err := Getter(ctx, issueID)
	if err != nil {
		err = fmt.Errorf("Request:%v", err)
		return
	}
	defer body.Close()
	issue = &model.Issue{}

	err = json.NewDecoder(body).Decode(issue)

	if err != nil {
		return
	}

	if issue.CommentsInfo.Count > 0 {
		var comments []model.CommentType
		comments, err = GetCommentsByID(ctx, issueID)
		if err != nil {
			return
		}
		issue.Comments = comments
	}

	return issue, err
}
