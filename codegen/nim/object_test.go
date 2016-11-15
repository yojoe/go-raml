package nim

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateObjectFromRaml(t *testing.T) {
	Convey("generate object from raml", t, func() {
		var apiDef raml.APIDefinition
		err := raml.ParseFile("../fixtures/struct/struct.raml", &apiDef)
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Simple struct from raml", func() {
			err = generateObjects(apiDef.Types, targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/object/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"EnumCity.nim", "EnumCity.nim"},
				{"EnumEnumCityEnum_homeNum.nim", "EnumEnumCityEnum_homeNum.nim"},
				{"EnumEnumCityEnum_parks.nim", "EnumEnumCityEnum_parks.nim"},
				{"animal.nim", "animal.nim"},
				{"Cage.nim", "Cage.nim"},
				{"Cat.nim", "Cat.nim"},
				{"ArrayOfCats.nim", "ArrayOfCats.nim"},
				{"BidimensionalArrayOfCats.nim", "BidimensionalArrayOfCats.nim"},
				{"EnumString.nim", "EnumString.nim"}, // object is enum type
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
func TestGenerateObjectMethodBody(t *testing.T) {
	Convey("generate object from method body", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("generate request body", func() {
			var body raml.Bodies
			properties := map[string]interface{}{
				"age": map[interface{}]interface{}{
					"type": "integer",
				},
				"ID": map[interface{}]interface{}{
					"type": "string",
				},
				"item": map[interface{}]interface{}{},
				"grades": map[interface{}]interface{}{
					"type": "integer[]",
				},
			}
			body.ApplicationJSON = &raml.BodiesProperty{
				Properties: properties,
			}

			_, err := generateObjectFromBody("usersPost", &body, true, targetDir)
			So(err, ShouldBeNil)

			s, err := testLoadFile(filepath.Join(targetDir, "usersPostReqBody.nim"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/object/usersPostReqBody.nim")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

		})

		Convey("Simple object from raml", func() {
			var apiDef raml.APIDefinition
			err := raml.ParseFile("../fixtures/struct/struct.raml", &apiDef)
			So(err, ShouldBeNil)

			_, err = generateObjectsFromBodies(getAllResources(&apiDef, true), targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/object/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"usersPostReqBody.nim", "usersPostReqBody.nim"},
				{"usersByIdGetRespBody.nim", "usersByIdGetRespBody.nim"},
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
