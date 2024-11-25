package command

// A Meta is a struct that represents the meta-options that are available on all or most commands.
//
// The Command interface is based on OpenTofu.
// See: https://github.com/opentofu/opentofu/blob/main/internal/command/meta.go
type Meta struct {
	// IsColored tells whether or not the printed output should be colored.
	IsColored bool
}
