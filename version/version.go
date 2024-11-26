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
	v, err := semver.NewVersion(strings.TrimSpace(rawVersion))
	if err != nil {
		// TODO: Add a better way of handling error in parsing the version.
		panic(fmt.Sprintf("Error parsing the semantic version string \"%s\": %v", strings.TrimSpace(rawVersion), err.Error()))
	}

	version = v
}

func Version() *semver.Version {
	if version == nil {
		Create()
	}

	return version
}
