package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLibrary(t *testing.T) {
	Convey("Library usage in server", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		err = GenerateServer("./fixtures/libraries/api.raml", targetDir, "main", langPython, "apidocs", "examples.com/ramlcode", true)
		So(err, ShouldBeNil)

		rootFixture := "./fixtures/libraries/python_server"
		checks := []struct {
			Result   string
			Expected string
		}{
			{"Place.py", "Place.py"},
			{"configs.py", "configs.py"},
			{"libraries/security/oauth2_Dropbox.py", "libraries/security/oauth2_Dropbox.py"},
			{"libraries/files/Directory.py", "libraries/files/Directory.py"},
		}

		for _, check := range checks {
			s, err := testLoadFile(filepath.Join(targetDir, check.Result))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		}

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
	Convey("Library usage in client", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("./fixtures/libraries/api.raml", apiDef)
		So(err, ShouldBeNil)

		err = GenerateClient(apiDef, targetDir, langGo, "examples.com/client")
		So(err, ShouldBeNil)

		rootFixture := "./fixtures/libraries/go_client"
		checks := []struct {
			Result   string
			Expected string
		}{
			{"Place.go", "Place.txt"},
			{"client_exampleapi.go", "client_exampleapi.txt"},
			{"client_utils.go", "client_utils.txt"},
		}

		for _, check := range checks {
			s, err := testLoadFile(filepath.Join(targetDir, check.Result))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		}

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}
