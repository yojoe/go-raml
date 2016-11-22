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
			checks := []struct {
				Result   string
				Expected string
			}{
				{"SingleInheritance.go", "singleinheritance.txt"},
				{"MultipleInheritance.go", "multipleinheritance.txt"},
				{"ArrayOfCats.go", "arrayofcats.txt"},
				{"BidimensionalArrayOfCats.go", "bidimensionalarrayofcats.txt"},
				{"petshop.go", "petshop.txt"},                   // using map type & testing case sensitive type name
				{"Pet.go", "Pet.txt"},                           // Union
				{"ArrayOfPets.go", "ArrayOfPets.txt"},           // Array of union
				{"Specialization.go", "Specialization.txt"},     // Specialization
				{"EnumCity.go", "enumcity.txt"},                 // Enum Field
				{"animal.go", "animal.txt"},                     // using enum
				{"EnumString.go", "enumstring.txt"},             // Enum type
				{"ValidationString.go", "ValidationString.txt"}, // validation
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("With included & inline JSON ", func() {
			err := raml.ParseFile("../fixtures/struct/json/api.raml", apiDef)
			So(err, ShouldBeNil)

			err = generateStructs(apiDef.Types, targetDir, "main")
			So(err, ShouldBeNil)

			err = generateBodyStructs(apiDef, targetDir, "main")
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/json"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"PersonInclude.go", "PersonInclude.txt"},
				{"PersonPostReqBody.go", "PersonPostReqBody.txt"},
				{"PersonGetRespBody.go", "PersonGetRespBody.txt"},
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
