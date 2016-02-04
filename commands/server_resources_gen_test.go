package commands

import (
	"io/ioutil"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func testLoadFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	return string(b), err
}
func TestResource(t *testing.T) {
	Convey("test resource generator", t, func() {
		Convey("Empty resource", func() {
			rs := map[string]raml.Resource{
				"/gists": raml.Resource{
					DisplayName: "Gists",
					Nested: map[string]*raml.Resource{
						"/public": &raml.Resource{
							DisplayName: "Public Gists",
						},
					},
				},
			}
			err := ServerResourcesGen(rs, "./tmp")
			So(err, ShouldBeNil)

			// check interface file
			s, err := testLoadFile("./tmp/gists_if.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/gists_if.go")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)
		})

		Convey("simple resource", func() {
			apiDef, err := raml.ParseFile("./fixtures/resource_deliveries.raml")
			So(err, ShouldBeNil)

			err = ServerResourcesGen(apiDef.Resources, "./tmp")
			So(err, ShouldBeNil)

			// check interface file
			s, err := testLoadFile("./tmp/deliveries_if.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/deliveries_if.go")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check api implemetation file
			s, err = testLoadFile("./tmp/deliveries_api.go")
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/deliveries_api.go")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)
		})

	})
}
