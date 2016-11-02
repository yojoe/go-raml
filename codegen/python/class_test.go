package python

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
		apiDef := new(raml.APIDefinition)
		err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
		So(err, ShouldBeNil)
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("python class from raml Types", func() {
			err = generatePythonClasses(apiDef.Types, targetDir)
			So(err, ShouldBeNil)

			// strings validator
			s, err := testLoadFile(filepath.Join(targetDir, "ValidationString.py"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("../fixtures/struct/ValidationString.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// with form field
			s, err = testLoadFile(filepath.Join(targetDir, "Cage.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("../fixtures/struct/Cage.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// FieldList of FormField
			s, err = testLoadFile(filepath.Join(targetDir, "animal.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("../fixtures/struct/animal.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// another test could be seen at body_test.go
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})

	})
}
