package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPythonResource(t *testing.T) {
	Convey("resource generator", t, func() {
		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("resource with request body", func() {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/server_resources/deliveries.raml", apiDef)
			So(err, ShouldBeNil)

			fs := NewFlaskServer(apiDef, "apidocs", true)

			err = fs.generateResources(targetdir)
			So(err, ShouldBeNil)

			// check  api implementation
			s, err := testLoadFile(filepath.Join(targetdir, "deliveries.py"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("../fixtures/server_resources/deliveries.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)
		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
