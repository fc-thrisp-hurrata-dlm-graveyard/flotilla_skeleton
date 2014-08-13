package fleet_skeleton

import "github.com/thrisp/fleet"

//import "lcl/fleet"

func FSkeletonTest(f *fleet.Context) {
	f.Data(200, "", []byte("fleetskeleton return value"))
}

func FskelCreate() (e *fleet.Engine) {
	asst := &fleet.EngineAsset{Asset, AssetNames, AssetDir}
	e = fleet.Basic()
	e.IEngineAsset = asst
	e.GET("/fleet/skeleton/:test", FSkeletonTest)
	return e
}
