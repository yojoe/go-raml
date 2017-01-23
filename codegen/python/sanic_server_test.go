package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Jumpscale/go-raml/raml"
)

func TestSanicServer(t *testing.T) {
	Convey("sanic server generator", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Hello world server", func() {
			apiDef := new(raml.APIDefinition)
			_, err = raml.ParseReadFile("../fixtures/raml-examples/helloworld/helloworld.raml", apiDef)
			So(err, ShouldBeNil)

			server := NewSanicServer(apiDef, "apidocs", true)
			err = server.Generate(targetDir)
			So(err, ShouldBeNil)

			// check drones API implementation
			rootFixture := "./fixtures/sanic/raml-examples/helloworld"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"app.py", "app.py"},
				{"helloworld_api.py", "helloworld_api.py"},
				{"helloworld_if.py", "helloworld_if.py"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("Congo", func() {
			apiDef := new(raml.APIDefinition)
			_, err = raml.ParseReadFile("../fixtures/congo/api.raml", apiDef)
			So(err, ShouldBeNil)

			server := NewSanicServer(apiDef, "apidocs", true)
			err = server.Generate(targetDir)
			So(err, ShouldBeNil)

			// check drones API implementation
			rootFixture := "./fixtures/sanic/server/congo"
			files := []string{
				"app.py",
				"deliveries_api.py",
				"deliveries_if.py",
				"drones_api.py",
				"drones_if.py",
				"schema/User_schema.json",
			}

			for _, filename := range files {
				s, err := testLoadFile(filepath.Join(targetDir, filename))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, filename))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
