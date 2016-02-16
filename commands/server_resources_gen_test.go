package commands

import (
	"io/ioutil"
	"os"
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
		Convey("simple resource", func() {
			apiDef, err := raml.ParseFile("./fixtures/server_resources/deliveries.raml")
			So(err, ShouldBeNil)

			_, err = generateServerResources(apiDef.Resources, "./tmp")
			So(err, ShouldBeNil)

			// check interface file
			s, err := testLoadFile("./tmp/deliveries_if.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/server_resources/deliveries_if.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check api implemetation file
			s, err = testLoadFile("./tmp/deliveries_api.go")
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server_resources/deliveries_api.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)
		})

		Convey("resource with request body", func() {
			apiDef, err := raml.ParseFile("./fixtures/server_resources/usergroups.raml")
			So(err, ShouldBeNil)

			_, err = generateServerResources(apiDef.Resources, "./tmp")
			So(err, ShouldBeNil)

			// check users api implementation
			s, err := testLoadFile("./tmp/users_api.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/server_resources/users_api.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check user interface
			s, err = testLoadFile("./tmp/users_if.go")
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server_resources/users_if.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

		})

		Reset(func() {
			os.RemoveAll("./tmp")
		})
	})
}
