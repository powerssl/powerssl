package version

import "fmt"

var BuildDate, GitBranch, GitCommit, GitState, GitSummary, Version string

func String() string {
	if GitState == "clean" {
		return fmt.Sprintf("%s\nbuild date %s", Version, BuildDate)
	}
	return fmt.Sprintf("%s (%s)\nbuild date %s", Version, GitSummary, BuildDate)
}
