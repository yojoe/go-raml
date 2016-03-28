package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServer(t *testing.T) {
	Convey("server generator", t, func() {
		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)
		Convey("simple Go server", func() {
			apiDef, err := raml.ParseFile("./fixtures/server/user_api/api.raml")
			So(err, ShouldBeNil)

			err = GenerateServer(apiDef, targetdir, "main", "go", true)
			So(err, ShouldBeNil)

			// check users api implementation
			s, err := testLoadFile(filepath.Join(targetdir, "users_api.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/server/user_api/users_api.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check user interface
			s, err = testLoadFile(filepath.Join(targetdir, "users_if.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server/user_api/users_if.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check main file
			s, err = testLoadFile(filepath.Join(targetdir, "main.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server/user_api/main.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

		})

		Convey("Congo python server", func() {
			apiDef, err := raml.ParseFile("./fixtures/congo/api.raml")
			So(err, ShouldBeNil)

			err = GenerateServer(apiDef, targetdir, "main", "python", true)
			So(err, ShouldBeNil)

			// check drones API implementation
			s, err := testLoadFile(filepath.Join(targetdir, "drones.py"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/congo/python_server/drones.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check deliveries API implementation
			s, err = testLoadFile(filepath.Join(targetdir, "deliveries.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/congo/python_server/deliveries.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check main file
			s, err = testLoadFile(filepath.Join(targetdir, "app.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/congo/python_server/app.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
