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
			apiDef, err := raml.ParseFile("./fixtures/resource_deliveries.raml")
			So(err, ShouldBeNil)

			err = ServerResourcesGen(apiDef.Resources, "./tmp")
			So(err, ShouldBeNil)

			// check interface file
			s, err := testLoadFile("./tmp/deliveries_if.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/deliveries_if.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check api implemetation file
			s, err = testLoadFile("./tmp/deliveries_api.go")
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/deliveries_api.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)
		})

		Reset(func() {
			os.RemoveAll("./tmp")
		})
	})
}
