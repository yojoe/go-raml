package commands

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
		apiDef, err := raml.ParseFile("./fixtures/client_resources/client.raml")
		So(err, ShouldBeNil)

		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Simple client from raml", func() {
			err = generateClient(apiDef, targetdir, "go")
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
