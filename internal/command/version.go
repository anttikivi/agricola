package command

// A Version is the version subcommand.
type Version struct{}

func (v *Version) Execute(_ []string) int {
	return 0
}

func (v *Version) Summary() string {
	return ""
}
