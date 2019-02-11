package version

import (
	"fmt"
	"runtime"
)

var (
	PreetyJsonVersion = ""
	GitSHA            = ""
)

func Version() string {
	version := `: %s
git SHA: %s
go version: %s
go OS/arch: %s/%s`

	return fmt.Sprintf(version, PreetyJsonVersion, GitSHA, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
