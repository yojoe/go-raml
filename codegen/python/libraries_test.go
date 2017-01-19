package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Jumpscale/go-raml/raml"
)

func TestLibrary(t *testing.T) {
	Convey("Library usage in server", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		_, err = raml.ParseReadFile("../fixtures/libraries/api.raml", apiDef)
		So(err, ShouldBeNil)

		server := FlaskServer{
			APIDef:   apiDef,
			Title:    apiDef.Title,
			WithMain: true,
		}
		err = server.Generate(targetDir)
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

}
