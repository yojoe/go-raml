package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateStructBodyFromRaml(t *testing.T) {
	Convey("generate struct body from raml", t, func() {
		apiDef, err := raml.ParseFile("./fixtures/struct/struct.raml")
		So(err, ShouldBeNil)

		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("simple body", func() {
			err := generateBodyStructs(apiDef, targetdir, "main")
			So(err, ShouldBeNil)

			//load and compare UsersIdGetRespBody
			s, err := testLoadFile(filepath.Join(targetdir, "UsersIdGetRespBody.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/struct/UsersIdGetRespBody.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			//load and compare usersgetreqbody
			s, err = testLoadFile(filepath.Join(targetdir, "UsersPostReqBody.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/UsersPostReqBody.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
