package fleet_skeleton

import (
	"github.com/thrisp/fleet"
)

func FSkeletonTest(f *fleet.Context) {
	f.Data(200, "", "fleetskeleton return value")
}

func FskelCreate() (e *fleet.Engine) {
	e := fleet.Basic()
	e.GET("/fleet/skeleton/:test", FSkeletonTest)
	return e
}
