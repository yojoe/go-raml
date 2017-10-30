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
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Congo python server", func() {
			apiDef := new(raml.APIDefinition)
			err = raml.ParseFile("../fixtures/congo/api.raml", apiDef)
			So(err, ShouldBeNil)

			server := NewFlaskServer(apiDef, "apidocs", true, nil)
			err = server.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "../fixtures/congo/python_server"
			files := []string{
				"drones_api.py",
				"deliveries_api.py",
				"app.py",
			}

			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
