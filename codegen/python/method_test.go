package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMethod(t *testing.T) {
	Convey("server method with display name", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("resource with request body", func() {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/server_resources/display_name/api.raml", apiDef)
			So(err, ShouldBeNil)

			fs := NewFlaskServer(apiDef, "apidocs", true)

			err = fs.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/method/flask/display_name"
			files := []string{
				"CreateSomethingCoolReqBody.py",
				"coolness.py",
			}

			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
