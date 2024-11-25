package cmd

import "github.com/anttikivi/agricola/internal/logging"

func Execute() int {
	defer logging.PanicHandler()

	return 0
}
