// The Ui package is based on `mitchellh/cli` by Mitchell Hashimoto.
// See: https://github.com/mitchellh/cli
package ui

// Ui is an interface for interacting with the terminal, or "interface" of the CLI.
type Ui interface {
	// Output prints to the set standard output.
	Output(string)

	// Error prints to the set standard error.
	Error(string)
}
