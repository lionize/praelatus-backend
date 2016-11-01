package models

// Workflow is the container for issues and keeps track of available transitions
type Workflow struct{}

// Transition contains information about what hooks to perform when performing
// a transition
type Transition struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	WorkflowID int64 `json:"-" db:"workflow_id"`
	StatusID   int64 `json:"-" db:"status_id"`
}

// TransitionJSON contains additional information for when we are turning a
// Transition into JSON
type TransitionJSON struct {
	Transition

	ToStatus Status `json:"to_status"`
}

// Hook contains information about what webhooks to fire when a given
// transition is run.
type Hook struct{}
