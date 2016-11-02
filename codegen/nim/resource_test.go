package nim

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateResourceAPI(t *testing.T) {
	Convey("generate object from raml", t, func() {
		var apiDef raml.APIDefinition
		err := raml.ParseFile("../fixtures/server_resources/deliveries.raml", &apiDef)
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("deliveries API", func() {
			err = generateResourceAPIs(getAllResources(&apiDef, true), targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/resource/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"deliveries_api.nim", "deliveries_api.nim"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
