package version

import "fmt"

// See http://semver.org/ for more information on Semantic Versioning
var (
	Major = 1
	Minor = 0
	Patch = 0
)

var Version = fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
