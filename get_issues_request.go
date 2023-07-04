// This file was auto-generated by Fern from our API Definition.

package api

type GetIssuesRequest struct {
	Id             *string                                         `json:"-"`
	IssueTriggerId *string                                         `json:"-"`
	Type           *GetIssuesRequestTypeGetIssuesRequestType       `json:"-"`
	Status         *GetIssuesRequestStatusGetIssuesRequestStatus   `json:"-"`
	MergedWith     *string                                         `json:"-"`
	CreatedAt      *string                                         `json:"-"`
	FirstSeenAt    *string                                         `json:"-"`
	LastSeenAt     *string                                         `json:"-"`
	DismissedAt    *string                                         `json:"-"`
	OrderBy        *GetIssuesRequestOrderByGetIssuesRequestOrderBy `json:"-"`
	Dir            *GetIssuesRequestDirGetIssuesRequestDir         `json:"-"`
	Limit          *int                                            `json:"-"`
	Next           *string                                         `json:"-"`
	Prev           *string                                         `json:"-"`
}
