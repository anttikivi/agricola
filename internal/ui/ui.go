// Package ui implements utilities for interacting with the user. It should be
// used for interacting with the user in contrast to the logger that is used to
// record messages related to the execution of the program. For example, the
// output printed by the utilities in the ui package should be printed visibly
// to the command line by default.
package ui

// A UserInterface is used to interact with the user.
type UserInterface interface {
	// Error prints an error message to the defined standard error output.
	Error(msg string)

	// Output prints a message to the defined standard output.
	Output(msg string)
}
