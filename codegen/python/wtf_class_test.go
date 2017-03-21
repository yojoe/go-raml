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
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("python class from raml Types", func() {
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			So(err, ShouldBeNil)

			err = generateWtfClasses(apiDef.Types, targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/class/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"ValidationString.py", "ValidationString.py"}, // strings validator
				{"Cage.py", "Cage.py"},                         // with form field
				{"animal.py", "animal.py"},                     // FieldList of FormField
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("python class from raml with JSON", func() {
			err := raml.ParseFile("../fixtures/struct/json/api.raml", apiDef)
			So(err, ShouldBeNil)

			err = generateWtfClasses(apiDef.Types, targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/class/json/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"PersonInclude.py", "PersonInclude.py"},
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
