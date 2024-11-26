package version

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
)

//go:embed VERSION
var rawVersion string

// version is the parsed semantic version of Agricola.
var version *semver.Version

func Create() {
	trimmed := strings.TrimSpace(rawVersion)
	v, err := semver.NewVersion(trimmed)
	if err != nil {
		panic(fmt.Sprintf("Error parsing the semantic version string \"%s\": %v", trimmed, err.Error()))
	}
	version = v
}

func Version() *semver.Version {
	if version == nil {
		Create()
	}
	return version
}
