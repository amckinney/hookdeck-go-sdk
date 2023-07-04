// This file was auto-generated by Fern from our API Definition.

package api

type BatchOperationPlan struct {
	// Number of batches required to complete the bulk retry
	EstimatedBatch *int `json:"estimated_batch,omitempty"`
	// Number of estimated events to be retried
	EstimatedCount *int `json:"estimated_count,omitempty"`
	// Progression of the batch operations, values 0 - 1
	Progress *float64 `json:"progress,omitempty"`
}
