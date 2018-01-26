package golang

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Jumpscale/go-raml/raml"
)

func TestSetImportPath(t *testing.T) {
	Convey("TestSetImportPath", t, func() {
		oriGoPath := os.Getenv("GOPATH")
		Convey("users api", func() {
			fakeGopath := "/gopath"
			os.Setenv("GOPATH", fakeGopath)
			So(setRootImportPath("import.com/a", "target"), ShouldEqual, "import.com/a")
			So(setRootImportPath("", "/gopath/src/johndoe.com/cool"), ShouldEqual, "johndoe.com/cool")
		})

		Reset(func() {
			os.Setenv("GOPATH", oriGoPath)
		})
	})
}

func TestCheckDuplicatedTitleTypes(t *testing.T) {
	Convey("TestCheckDuplicatedTitleTypes", t, func() {
		tests := []struct {
			apiDef *raml.APIDefinition
			err    bool
		}{
			{
				&raml.APIDefinition{
					Types: map[string]raml.Type{
						"One": raml.Type{},
						"one": raml.Type{},
					},
				}, true,
			},
			{
				&raml.APIDefinition{
					Types: map[string]raml.Type{
						"One": raml.Type{},
						"oNe": raml.Type{},
					},
				}, false,
			},
		}

		for _, test := range tests {
			err := checkDuplicatedTitleTypes(test.apiDef)
			if test.err {
				So(err, ShouldNotBeNil)
			} else {
				So(err, ShouldBeNil)
			}
		}

	})
}
