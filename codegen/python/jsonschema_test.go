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
