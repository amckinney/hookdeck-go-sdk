// This file was auto-generated by Fern from our API Definition.

package api

type ConnectionPaginatedResult struct {
	Pagination *SeekPagination `json:"pagination,omitempty"`
	Count      *int            `json:"count,omitempty"`
	Models     *[]*Connection  `json:"models,omitempty"`
}
