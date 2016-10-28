package models

// Workflow is the container for issues and keeps track of available transitions
type Workflow struct{}

// WorkflowTransition contains information about what hooks to perform when
// performing this transition
type WorkflowTransition struct{}

// Hook contains information about what webhooks to fire when a given
// transition is run.
type Hook struct{}
