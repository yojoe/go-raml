package jsonschema

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Jumpscale/go-raml/raml"
)

func TestGenerateStructFromRaml(t *testing.T) {
	Convey("JSON Schema test", t, func() {
		apiDef := new(raml.APIDefinition)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("array", func() {
			err := raml.ParseFile("../../fixtures/raml-examples/typesystem/array-type.raml", apiDef)
			So(err, ShouldBeNil)

			for name, t := range apiDef.Types {
				js := NewJSONSchema(t, name)
				err := js.Generate(targetDir)
				So(err, ShouldBeNil)
			}
			rootFixture := "./fixtures/array_type"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"Email_schema.json", "Email_schema.json"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("simple", func() {
			err := raml.ParseFile("../../fixtures/raml-examples/typesystem/simple.raml", apiDef)
			So(err, ShouldBeNil)

			for name, t := range apiDef.Types {
				js := NewJSONSchema(t, name)
				err := js.Generate(targetDir)
				So(err, ShouldBeNil)
			}
			rootFixture := "./fixtures/simple"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"User_schema.json", "User_schema.json"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("struct", func() {
			err := raml.ParseFile("../../fixtures/struct/struct.raml", apiDef)
			So(err, ShouldBeNil)

			for name, t := range apiDef.Types {
				js := NewJSONSchema(t, name)
				err := js.Generate(targetDir)
				So(err, ShouldBeNil)
			}
			rootFixture := "./fixtures/struct"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"EnumCity_schema.json", "EnumCity_schema.json"},
				{"EnumString_schema.json", "EnumString_schema.json"},
				{"animal_schema.json", "animal_schema.json"},
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

func testLoadFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	return string(b), err
}
