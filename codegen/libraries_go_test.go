package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoLibrary(t *testing.T) {
	Convey("Library usage in server", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		err = GenerateServer("./fixtures/libraries/api.raml", targetDir, "main", "go", "apidocs", "examples.com/ramlcode", true)
		So(err, ShouldBeNil)

		rootFixture := "./fixtures/libraries/go_server"
		checks := []struct {
			Result   string
			Expected string
		}{
			{"Place.go", "Place.txt"},
			{"dirs_api.go", "dirs_api.txt"},
			{"configs_api.go", "configs_api.txt"},
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

		err = GenerateClient(apiDef, targetDir, "theclient", langGo, "examples.com/theclient")
		So(err, ShouldBeNil)

		rootFixture := "./fixtures/libraries/go_client"
		checks := []struct {
			Result   string
			Expected string
		}{
			{"Place.go", "Place.txt"},
			{"client_exampleapi.go", "client_exampleapi.txt"},
			{"client_utils.go", "client_utils.txt"},
			{"dirs_service.go", "dirs_service.txt"},
			{"configs_service.go", "configs_service.txt"},
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
