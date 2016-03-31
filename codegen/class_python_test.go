package codegen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGeneratePythonClass(t *testing.T) {
	Convey("generate python class from raml", t, func() {
		apiDef, err := raml.ParseFile("./fixtures/struct/struct.raml")
		So(err, ShouldBeNil)
		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("python class from raml Types", func() {
			err = generatePythonClasses(apiDef, targetdir)
			So(err, ShouldBeNil)

			// strings validator
			s, err := testLoadFile(filepath.Join(targetdir, "ValidationString.py"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/struct/ValidationString.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

		})
		Reset(func() {
			os.RemoveAll(targetdir)
		})

	})
}
