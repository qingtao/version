package version

import "github.com/coreos/go-semver/semver"

var (
	// Version 版本号
	Version = "v0.0.0"
	// GitSHA git提交记录, 使用"git rev-parse HEAD"
	GitSHA string
	// GitTag git的标签, 使用"git describe --tags --dirty --always"
	GitTag string
)

func init() {
	v, err := semver.NewVersion(Version)
	if err == nil {
		Version = v.String()
	}
}
