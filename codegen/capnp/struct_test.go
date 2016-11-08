package capnp

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateCapnpSchema(t *testing.T) {
	Convey("generate capnp schema from raml", t, func() {
		var apiDef raml.APIDefinition
		err := raml.ParseFile("./fixtures/struct.raml", &apiDef)
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Schema for Python & Nim", func() {
			err = GenerateCapnp(&apiDef, targetDir, "nim", "")
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/vanilla"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"Animal.capnp", "Animal.capnp"},
				{"Cage.capnp", "Cage.capnp"},
				{"Admin.capnp", "Admin.capnp"},
				{"EnumClearanceLevel.capnp", "EnumClearanceLevel.capnp"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("Schema for Go", func() {
			err = GenerateCapnp(&apiDef, targetDir, "go", "main")
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/golang"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"Animal.capnp", "Animal.capnp"},
				{"Cage.capnp", "Cage.capnp"},
				{"Admin.capnp", "Admin.capnp"},
				{"EnumClearanceLevel.capnp", "EnumClearanceLevel.capnp"},
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
	return removeID(string(b)), err
}

// remove capnp ID from a file, we need it for the test.
// because capnp ID will always produce different value
// this func is not elegant
func removeID(s string) string {
	splt := strings.Split(s, "\n")
	clean := []string{}
	for i, v := range splt {
		if strings.HasPrefix(v, "@0x") {
			clean = append(clean, splt[i+1:]...)
			break
		}
		clean = append(clean, v)
	}
	return strings.Join(clean, "\n")
}
