package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Jumpscale/go-raml/raml"
)

func TestServer(t *testing.T) {
	Convey("server generator", t, func() {
		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Congo python server", func() {
			apiDef := new(raml.APIDefinition)
			_, err = raml.ParseReadFile("../fixtures/congo/api.raml", apiDef)
			So(err, ShouldBeNil)

			server := NewFlaskServer(apiDef, "apidocs", true)
			err = server.Generate(targetdir)
			So(err, ShouldBeNil)

			// check drones API implementation
			s, err := testLoadFile(filepath.Join(targetdir, "drones.py"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("../fixtures/congo/python_server/drones.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check deliveries API implementation
			s, err = testLoadFile(filepath.Join(targetdir, "deliveries.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("../fixtures/congo/python_server/deliveries.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check main file
			s, err = testLoadFile(filepath.Join(targetdir, "app.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("../fixtures/congo/python_server/app.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
