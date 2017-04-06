package golang

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtil(t *testing.T) {
	Convey("TestUtil", t, func() {
		oriGoPath := os.Getenv("GOPATH")
		Convey("users api", func() {
			fakeGopath := "/gopath"
			os.Setenv("GOPATH", fakeGopath)
			So(setRootImportPath("", "target"), ShouldEqual, "")
			So(setRootImportPath("import.com/a", "target"), ShouldEqual, "import.com/a")
			So(setRootImportPath("", "/gopath/src/johndoe.com/cool"), ShouldEqual, "johndoe.com/cool")
			So(setRootImportPath("", "/gopath/johndoe.com/cool"), ShouldEqual, "")
		})

		Reset(func() {
			os.Setenv("GOPATH", oriGoPath)
		})
	})
}
