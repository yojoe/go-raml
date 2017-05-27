package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateClientFromRaml(t *testing.T) {
	Convey("generate client from raml", t, func() {
		apiDef := new(raml.APIDefinition)
		err := raml.ParseFile("./fixtures/client_resources/client.raml", apiDef)
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		err = GenerateClient(apiDef, targetDir, "theclient", "go", "client", "", nil)
		So(err, ShouldBeNil)
		rootFixture := "./golang/fixtures/client_resources"
		checks := []struct {
			Result   string
			Expected string
		}{
			{"client_structapitest.go", "client_structapitest.txt"},
			{"users_service.go", "users_service.txt"},
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

func testLoadFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	return string(b), err
}
