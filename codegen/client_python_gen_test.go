package commands

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
		apiDef, err := raml.ParseFile("./fixtures/python_client/client.raml")
		So(err, ShouldBeNil)

		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Simple client", func() {
			err = generateClient(apiDef, targetdir, "python")
			So(err, ShouldBeNil)

			// cek with generated with fixtures
			fixture, err := testLoadFile("./fixtures/python_client/client.py")
			So(err, ShouldBeNil)

			generated, err := testLoadFile(filepath.Join(targetdir, "client.py"))
			So(err, ShouldBeNil)

			So(generated, ShouldEqual, fixture)
		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
