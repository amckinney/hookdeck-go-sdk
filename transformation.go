// This file was auto-generated by Fern from our API Definition.

package api

type Transformation struct {
	// ID of the transformation
	Id string `json:"id"`
	// ID of the workspace
	TeamId string `json:"team_id"`
	// A unique, human-friendly name for the transformation
	Name string `json:"name"`
	// JavaScript code to be executed
	Code         string  `json:"code"`
	EncryptedEnv *string `json:"encrypted_env,omitempty"`
	Iv           *string `json:"iv,omitempty"`
	// Key-value environment variables to be passed to the transformation
	Env *map[string]*string `json:"env,omitempty"`
	// Date the transformation was last updated
	UpdatedAt string `json:"updated_at"`
	// Date the transformation was created
	CreatedAt string `json:"created_at"`
}
