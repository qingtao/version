// Package version 通过导入此模块并使用编译脚本或者Makefile可以在编译的可执行程序中增加版本号和git仓库的分支/标签/提交记录ID
package version

import "github.com/coreos/go-semver/semver"

var (
	// Version 版本号
	Version = "v0.0.0"
	// GitSHA git提交记录, 使用"git rev-parse HEAD"
	GitSHA = "unknown"
	// GitTag git的标签, 使用"git describe --tags --dirty --always"
	GitTag = "unknown"
	// GitBranch git分支
	GitBranch = "unknown"
)

func init() {
	v, err := semver.NewVersion(Version)
	if err == nil {
		Version = v.String()
	}
}
