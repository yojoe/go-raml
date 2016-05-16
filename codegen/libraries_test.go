package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLibrary(t *testing.T) {
	Convey("simple library", t, func() {
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
}
