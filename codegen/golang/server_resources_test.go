package golang

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
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)

		Convey("simple resource", func() {
			err := raml.ParseFile("../fixtures/server_resources/deliveries.raml", apiDef)
			So(err, ShouldBeNil)

			gs := NewServer(apiDef, "main", "apidocs", "examples.com/libro", true, false, targetDir, nil)
			_, err = gs.generateServerResources(targetDir)
			So(err, ShouldBeNil)

			// check interface file
			s, err := testLoadFile(filepath.Join(targetDir, "deliveries_if.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/server_resources/simple/deliveries_if.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check api implemetation file
			s, err = testLoadFile(filepath.Join(targetDir, "deliveries_api.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server_resources/simple/deliveries_api.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			rootFixture := "./fixtures/server_resources/simple"
			files := []string{
				"deliveries_api",
				"deliveries_if",
			}
			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f+".go"))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f+".txt"))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("simple resource with one api file per method", func() {
			err := raml.ParseFile("../fixtures/server_resources/deliveries.raml", apiDef)
			So(err, ShouldBeNil)

			gs := NewServer(apiDef, "main", "apidocs", "examples.com/libro", true, true, targetDir, nil)
			_, err = gs.generateServerResources(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/server_resources/one_file_per_method_simple"
			files := []string{
				"deliveries_api",
				"deliveries_api_Get",
				"deliveries_api_Post",
				"deliveries_api_deliveryIdDelete",
				"deliveries_api_deliveryIdPatch",
				"deliveries_api_getDeliveriesByDeliveryID",
				"deliveries_if",
			}
			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f+".go"))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f+".txt"))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("big raml to check interface consistency", func() {
			err := raml.ParseFile("../fixtures/server_resources/grid/api.raml", apiDef)
			So(err, ShouldBeNil)

			gs := NewServer(apiDef, "main", "apidocs", "examples.com/libro", true, true, targetDir, nil)
			_, err = gs.generateServerResources(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/server_resources/grid/"
			files := []string{
				"nodes_if",
			}
			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f+".go"))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f+".txt"))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Convey("resource with request body", func() {
			err := raml.ParseFile("../fixtures/server_resources/usergroups.raml", apiDef)
			So(err, ShouldBeNil)

			gs := NewServer(apiDef, "main", "apidocs", "examples.com/libro", true, false, targetDir, nil)
			_, err = gs.generateServerResources(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/server_resources/with_req_body"
			files := []string{
				"users_if",
				"users_api",
			}
			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f+".go"))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f+".txt"))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
