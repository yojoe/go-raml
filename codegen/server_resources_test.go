package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func testLoadFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	return string(b), err
}
func TestResource(t *testing.T) {
	Convey("resource generator", t, func() {
		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)

		Convey("simple resource", func() {
			err := raml.ParseFile("./fixtures/server_resources/deliveries.raml", apiDef)
			So(err, ShouldBeNil)

			_, err = generateServerResources(apiDef, targetdir, "main", "go")
			So(err, ShouldBeNil)

			// check interface file
			s, err := testLoadFile(filepath.Join(targetdir, "deliveries_if.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/server_resources/deliveries_if.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check api implemetation file
			s, err = testLoadFile(filepath.Join(targetdir, "deliveries_api.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server_resources/deliveries_api.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)
		})

		Convey("resource with request body", func() {
			err := raml.ParseFile("./fixtures/server_resources/usergroups.raml", apiDef)
			So(err, ShouldBeNil)

			_, err = generateServerResources(apiDef, targetdir, "main", "go")
			So(err, ShouldBeNil)

			// check users api implementation
			s, err := testLoadFile(filepath.Join(targetdir, "users_api.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/server_resources/users_api.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check user interface
			s, err = testLoadFile(filepath.Join(targetdir, "users_if.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server_resources/users_if.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
