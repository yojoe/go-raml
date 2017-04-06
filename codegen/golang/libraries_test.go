package golang

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
		var apiDef raml.APIDefinition

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		err = raml.ParseFile("../fixtures/libraries/api.raml", &apiDef)
		So(err, ShouldBeNil)

		server := NewServer(&apiDef, "main", "apidocs", "examples.com/ramlcode", true, false, targetDir)
		err = server.Generate()
		So(err, ShouldBeNil)

		rootFixture := "../fixtures/libraries/go_server"
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
		err = raml.ParseFile("../fixtures/libraries/api.raml", apiDef)
		So(err, ShouldBeNil)

		client, err := NewClient(apiDef, "theclient", "examples.com/theclient", targetDir)
		So(err, ShouldBeNil)

		err = client.Generate()
		So(err, ShouldBeNil)

		rootFixture := "../fixtures/libraries/go_client"
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

	Convey("raml-examples", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("server", func() {
			var apiDef raml.APIDefinition

			err = raml.ParseFile("../fixtures/raml-examples/libraries/api.raml", &apiDef)
			So(err, ShouldBeNil)

			server := NewServer(&apiDef, "main", "apidocs", "examples.com/libro", true, false, targetDir)
			err = server.Generate()
			So(err, ShouldBeNil)

			rootFixture := "../fixtures/libraries/raml-examples/go_server"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"person_api.go", "person_api.txt"},
				{"types_lib/Person.go", "types_lib/Person.txt"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Convey("client", func() {
			var apiDef raml.APIDefinition

			err = raml.ParseFile("../fixtures/raml-examples/libraries/api.raml", &apiDef)
			So(err, ShouldBeNil)

			client, err := NewClient(&apiDef, "client", "examples.com/libro", targetDir)
			So(err, ShouldBeNil)

			err = client.Generate()
			So(err, ShouldBeNil)

			rootFixture := "../fixtures/libraries/raml-examples/go_client"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"person_service.go", "person_service.txt"},
				{"types_lib/Person.go", "types_lib/Person.txt"},
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
