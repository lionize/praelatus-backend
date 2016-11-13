package models

// Workflow is the container for issues and keeps track of available transitions
type Workflow struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (w *Workflow) String() string {
	return jsonString(w)
}

// Transition contains information about what hooks to perform when performing
// a transition
type Transition struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	ToStatus Status   `json:"to_status"`
	Workflow Workflow `json:"workflow,omitempty"`
}

func (t *Transition) String() string {
	return jsonString(t)
}

// Hook contains information about what webhooks to fire when a given
// transition is run.
type Hook struct{}
