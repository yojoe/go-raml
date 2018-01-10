package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"
)

func TestServer(t *testing.T) {
	Convey("server generator", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/congo/api.raml", apiDef)
		So(err, ShouldBeNil)
		rootFixture := "../fixtures/congo/python_server"
		files := []string{
			"drones_api.py",
			"deliveries_api.py",
			"app.py",
		}

		Convey("Congo python server", func() {
			server := NewFlaskServer(apiDef, "apidocs", true, nil, false)
			err = server.Generate(targetDir)
			So(err, ShouldBeNil)
			validateFiles(files, targetDir, rootFixture)
			// test that this file exist
			filesExist := []string{
				"types/User.py",
				"types/client_support.py",
			}
			for _, f := range filesExist {
				_, err := os.Stat(filepath.Join(targetDir, f))
				So(err, ShouldBeNil)
			}
		})

		Convey("Congo gevent python server", func() {
			files = append(files, "server.py")
			server := NewFlaskServer(apiDef, "apidocs", true, nil, true)
			err = server.Generate(targetDir)
			So(err, ShouldBeNil)
			validateFiles(files, targetDir, rootFixture)
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}

func validateFiles(files []string, targetDir string, rootFixture string) {
	for _, f := range files {
		s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
		So(err, ShouldBeNil)

		tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
		So(err, ShouldBeNil)

		So(s, ShouldEqual, tmpl)
	}
}