// This file was auto-generated by Fern from our API Definition.

package api

type BatchOperation struct {
	// ID of the bulk retry
	Id string `json:"id,omitempty"`
	// ID of the workspace
	TeamId string `json:"team_id,omitempty"`
	// Query object to filter records
	Query *BatchOperationQuery `json:"query,omitempty"`
	// Date the bulk retry was created
	CreatedAt string `json:"created_at,omitempty"`
	// Last time the bulk retry was updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// Date the bulk retry was cancelled
	CancelledAt *string `json:"cancelled_at,omitempty"`
	// Date the bulk retry was completed
	CompletedAt *string `json:"completed_at,omitempty"`
	// Number of batches required to complete the bulk retry
	EstimatedBatch *int `json:"estimated_batch,omitempty"`
	// Number of estimated events to be retried
	EstimatedCount *int `json:"estimated_count,omitempty"`
	// Number of batches currently processed
	ProcessedBatch *int `json:"processed_batch,omitempty"`
	// Number of events that were successfully delivered
	CompletedCount *int `json:"completed_count,omitempty"`
	// Indicates if the bulk retry is currently in progress
	InProgress bool `json:"in_progress,omitempty"`
	// Progression of the batch operations, values 0 - 1
	Progress *float64 `json:"progress,omitempty"`
	// Number of events that failed to be delivered
	FailedCount *int     `json:"failed_count,omitempty"`
	Number      *float64 `json:"number,omitempty"`
}
