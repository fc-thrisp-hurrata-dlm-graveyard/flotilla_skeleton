package flotilla_skeleton

import (
	"fmt"

	"github.com/thrisp/flotilla"
)

func HelloWorld(ctx *flotilla.R) error {
	fmt.Printf("\n\nHELLO WORLD context function extension passed: %s\n\n", ctx)
	return nil
}

func FSkeletonTestHandle1(ctx *flotilla.R) {
	ctx.ServeData(200, []byte("flotilla_skeleton return value"))
}

func FSkeletonTestHandle2(ctx *flotilla.R) {
	ctx.RenderTemplate("flotilla-skeleton-1.html", ctx)
}

func Fskel0() (a *flotilla.App) {
	a = flotilla.New("flotilla_skeleton_0")
	a.GET("/flotilla/skeleton0/:test0", FSkeletonTestHandle1)
	return a
}

func Fskel1() (a *flotilla.App) {
	a = flotilla.New("flotilla_skeleton_1")
	dd := &flotilla.AssetFS{Asset, AssetDir, AssetNames, ""}
	a.Assets = append(a.Assets, dd)
	a.STATIC("arbitrary/static/path")
	a.GET("/flotilla/skeleton1/:test0", FSkeletonTestHandle1)
	a.GET("/flotilla/skeleton1/:test0/template", FSkeletonTestHandle2)
	cnf, _ := a.Assets.GetByte("resources/flotilla-skeleton-1.conf")
	a.LoadConfByte(cnf, "flotilla-skeleton-1.conf")
	m := make(map[string]interface{})
	m["helloworld"] = HelloWorld
	a.AddCtxFuncs(m)
	return a
}

func Fskel2() (a *flotilla.App) {
	a = flotilla.New("flotilla_skeleton_2")
	dd := &flotilla.AssetFS{Asset, AssetDir, AssetNames, ""}
	a.Assets = append(a.Assets, dd)
	rg := a.New("/flotilla/skeleton2")
	rg.GET("/:test0/template", FSkeletonTestHandle2)
	return a
}
