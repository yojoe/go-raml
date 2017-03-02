package golang

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClient(t *testing.T) {
	Convey("Test client", t, func() {
		var apiDef raml.APIDefinition
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("users api", func() {
			err = raml.ParseFile("../fixtures/client_resources/client.raml", &apiDef)
			So(err, ShouldBeNil)

			client, err := NewClient(&apiDef, "theclient", "examples.com/libro")
			So(err, ShouldBeNil)

			err = client.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "../fixtures/client_resources"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"users_service.go", "users_service.txt"},
				{"client_structapitest.go", "client_structapitest.txt"},
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
