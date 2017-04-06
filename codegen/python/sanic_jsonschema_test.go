package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Jumpscale/go-raml/raml"
)

func TestJSONSchema(t *testing.T) {
	Convey("JSON Schema test", t, func() {
		apiDef := new(raml.APIDefinition)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("array", func() {
			err := raml.ParseFile("../fixtures/raml-examples/typesystem/array-type.raml", apiDef)
			So(err, ShouldBeNil)

			server := NewSanicServer(apiDef, "apidocs", true)
			err = server.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/json_schema/array_type"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"Email_schema.json", "Email_schema.json"},
				{"EmailsLong_schema.json", "EmailsLong_schema.json"},
				{"EmailsShort_schema.json", "EmailsShort_schema.json"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, server.schemaDir(), check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("simple", func() {
			err := raml.ParseFile("../fixtures/raml-examples/typesystem/simple.raml", apiDef)
			So(err, ShouldBeNil)

			server := NewSanicServer(apiDef, "apidocs", true)
			err = server.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/json_schema/simple"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"User_schema.json", "User_schema.json"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, server.schemaDir(), check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("struct", func() {
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			So(err, ShouldBeNil)

			server := NewSanicServer(apiDef, "apidocs", true)
			err = server.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/json_schema/struct"
			files := []string{
				"EnumCity_schema.json",
				"EnumString_schema.json",
				"animal_schema.json",
				"Cage_schema.json",
				"MultipleInheritance_schema.json",
				"Catanimal_schema.json",
				"UsersPostReqBody_schema.json",
			}

			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, server.schemaDir(), f))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
