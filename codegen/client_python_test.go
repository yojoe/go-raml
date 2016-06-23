package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGeneratePythonClientFromRaml(t *testing.T) {
	Convey("generate python client", t, func() {
		apiDef := new(raml.APIDefinition)
		err := raml.ParseFile("./fixtures/python_client/client.raml", apiDef)
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Simple client", func() {
			err = GenerateClient(apiDef, targetDir, "", "python", "")
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/python_client"
			// cek with generated with fixtures
			checks := []struct {
				Result   string
				Expected string
			}{
				{"client.py", "client.py"},
				{"__init__.py", "__init__.py"},
				{"client_utils.py", "client_utils.py"},
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
