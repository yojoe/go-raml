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

		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Simple client from raml", func() {
			err = GenerateClient(apiDef, targetdir, "go", "client")
			So(err, ShouldBeNil)

			s, err := testLoadFile(filepath.Join(targetdir, "client_structapitest.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/client_resources/client_structapitest.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)
		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
