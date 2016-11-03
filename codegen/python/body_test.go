package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateClassFromBody(t *testing.T) {
	Convey("generate struct body from raml", t, func() {
		apiDef := new(raml.APIDefinition)
		err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("python class from request/response bodies", func() {
			rs := getAllResources(apiDef, true)
			err = generateClassFromBodies(rs, targetDir)
			So(err, ShouldBeNil)

			// req body
			s, err := testLoadFile(filepath.Join(targetDir, "UsersPostReqBody.py"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("../fixtures/struct/UsersPostReqBody.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
