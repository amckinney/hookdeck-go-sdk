// This file was auto-generated by Fern from our API Definition.

package api

type SeekPagination struct {
	OrderBy *SeekPaginationOrderBy `json:"order_by,omitempty"`
	Dir     *SeekPaginationDir     `json:"dir,omitempty"`
	Limit   *int                   `json:"limit,omitempty"`
	Prev    *string                `json:"prev,omitempty"`
	Next    *string                `json:"next,omitempty"`
}
