package constants

// Command is the actual command to be used
type Command int

const (
	// Execute sql command with no return values
	Execute Command = iota
	// Query multiple rows
	Query Command = iota
	// QueryRow queries single row
	QueryRow Command = iota
)
