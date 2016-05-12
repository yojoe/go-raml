package raml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLibraries(t *testing.T) {
	Convey("Libraries", t, func() {
		apiDef := new(APIDefinition)

		Convey("simple library", func() {
			err := ParseFile("./samples/simple_with_lib.raml", apiDef)
			So(err, ShouldBeNil)

			// check Uses
			So(apiDef.Uses, ShouldContainKey, "files")
			So(apiDef.Uses["files"], ShouldEqual, "libraries/files.raml")

			// Check Libraries property
			So(apiDef.Libraries, ShouldContainKey, "files")

			files := apiDef.Libraries["files"]
			So(files.Usage, ShouldEqual, "Use to define some basic file-related constructs.")
			So(files.Traits, ShouldContainKey, "drm")
			So(files.Uses, ShouldContainKey, "file-type")
		})

	})
}
