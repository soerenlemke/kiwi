package log

// Action represents the type of action that can be performed in the logging system.
type Action int

const (
	// ActionUnknown unknown action placeholder.
	ActionUnknown = iota
	// ActionSet represents the action of setting a value.
	ActionSet
	// ActionDelete represents the action of deleting a value.
	ActionDelete
)
