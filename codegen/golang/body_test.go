package golang

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateStructBodyFromRaml(t *testing.T) {
	Convey("generate struct body from raml", t, func() {
		apiDef := new(raml.APIDefinition)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("simple body", func() {
			err = raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			So(err, ShouldBeNil)

			s := NewServer(apiDef, "main", "", "examples.com", false, false)
			err := s.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct"
			files := []string{
				"UsersIdGetRespBody",
				"UsersPostReqBody",
				"Catanimal",
				"users_api",
				"UnionCatanimal",
			}

			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f+".go"))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f+".txt"))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Convey("builtin type doesn't need validation code", func() {
			err = raml.ParseFile("../fixtures/struct/validation.raml", apiDef)
			So(err, ShouldBeNil)

			s := NewServer(apiDef, "main", "", "examples.com", false, false)
			err := s.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/validation"
			files := []string{
				"builtin_api",
			}

			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f+".go"))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f+".txt"))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
