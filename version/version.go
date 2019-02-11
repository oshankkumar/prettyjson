package version

import (
	"fmt"
	"runtime"
)

var (
	PrettyJsonVersion = ""
	GitSHA            = ""
)

func Version() string {
	version := `: %s
git SHA: %s
go version: %s
go OS/arch: %s/%s`

	return fmt.Sprintf(version, PrettyJsonVersion, GitSHA, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
