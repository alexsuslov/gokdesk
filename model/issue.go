package model

type Issue struct {
	ID                      int          `json:"id"`
	Title                   string       `json:"title"`
	CreatedAt               string       `json:"created_at"`
	CompletedAt             string       `json:"completed_at"`
	DeadlineAt              string       `json:"deadline_at"`
	StartExecutionUntil     string       `json:"start_execution_until"`
	PlannedExecutionInHours float32      `json:"planned_execution_in_hours"`
	DelayedTo               *string      `json:"delayed_to"`
	CompanyID               int          `json:"company_id"`
	GroupID                 int          `json:"group_id"`
	Status                  Status       `json:"status"`
	Assignee                User         `json:"assignee"`
	Author                  User         `json:"author"`
	Agreement               *string      `json:"agreement_id,omitempty"`
	Contact                 User         `json:"contact"`
	Priority                Priority     `json:"priority"`
	Type                    Type         `json:"type"`
	Attachments             []Attachment `json:"attachments"`
	Observers               []User       `json:"observers"`
	ObserverGroups          []Group      `json:"observer_groups"`
	Parameters              []Parameter  `json:"parameters"`
	CommentsInfo            Comments     `json:"comments"`
	Errors                  *string      `json:"errors"`
	Comments                []CommentType
}
