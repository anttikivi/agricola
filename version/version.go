package version

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
)

// rawVersion is the raw version value read from the VERSION file. It is used
// if buildVersion is not set.
//
//go:embed VERSION
var rawVersion string

// buildVersion is the version set using linker flags build time. It is used to
// over the value embedded from the VERSION file if set.
var buildVersion string //nolint:gochecknoglobals

// version is the parsed semantic version of Agricola.
var version *semver.Version

func Create() {
	str := buildVersion
	if str == "" {
		str = rawVersion
	}

	v, err := semver.NewVersion(strings.TrimSpace(str))
	if err != nil {
		// TODO: Add a better way of handling error in parsing the version.
		panic(fmt.Sprintf("Error parsing the semantic version string \"%s\": %v", strings.TrimSpace(str), err.Error()))
	}

	version = v
}

func Version() *semver.Version {
	if version == nil {
		Create()
	}

	return version
}
