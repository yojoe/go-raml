package golang

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
		apiDef := new(raml.APIDefinition)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Struct from raml", func() {
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			So(err, ShouldBeNil)

			err = generateStructs(apiDef.Types, targetDir, "main")
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct"
			files := []string{
				"SingleInheritance",
				"MultipleInheritance",
				"ArrayOfCats",
				"BidimensionalArrayOfCats",
				"petshop",          // using map type & testing case sensitive type name
				"Pet",              // Union
				"ArrayOfPets",      // Array of union
				"Specialization",   // Specialization
				"EnumCity",         // Enum Field
				"animal",           // using enum
				"EnumString",       // Enum type
				"ValidationString", // validation
				"Dashed",           // field with dash
				"PlainObject",
				"NumberFormat",
			}

			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f+".go"))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f+".txt"))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("With included & inline JSON ", func() {
			err := raml.ParseFile("../fixtures/struct/json/api.raml", apiDef)
			So(err, ShouldBeNil)

			err = generateStructs(apiDef.Types, targetDir, "main")
			So(err, ShouldBeNil)

			err = generateAllStructs(apiDef, targetDir, "main")
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/json"
			files := []string{
				"PersonInclude",
				"PersonPostReqBody",
				"PersonGetRespBody",
				"PersonInType",
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
