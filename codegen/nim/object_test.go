package nim

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateStructFromRaml(t *testing.T) {
	Convey("generate struct from raml", t, func() {
		var apiDef raml.APIDefinition
		err := raml.ParseFile("../fixtures/struct/struct.raml", &apiDef)
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Simple struct from raml", func() {
			err = GenerateObjects(apiDef.Types, targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/object/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"EnumCity.nim", "EnumCity.nim"},
				{"animal.nim", "animal.nim"},
				{"Cage.nim", "Cage.nim"},
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
